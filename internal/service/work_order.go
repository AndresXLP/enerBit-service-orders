package service

import (
	"context"
	"fmt"

	"enerBit-service-orders/api"
	"enerBit-service-orders/internal/entity"
	"enerBit-service-orders/internal/repos/postgres/models"
	pgRepo "enerBit-service-orders/internal/repos/postgres/repo"
	redisRepo "enerBit-service-orders/internal/repos/redis/repo"
	"enerBit-service-orders/pkg/utils"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkOrderService interface {
	CreateWorkOrder(ctx context.Context, req entity.CreateWorkOrderRequest) (*api.WorkOrder, error)
	FinishWorkOrder(ctx context.Context, req entity.FinishWorkOrder) (*api.WorkOrder, error)
	CancelWorkOrder(ctx context.Context, id uuid.UUID) error
	GetWorkOrders(ctx context.Context, req entity.GetWorkOrders) (*api.GetWorkOrdersResp, error)
	GetWorkOrderByCustomerID(ctx context.Context, customerID uuid.UUID) (*api.GetWorkOrdersResp, error)
	GetWorkOrderByIDWithCustomer(ctx context.Context, ID uuid.UUID) (*api.WorkOrderWithCustomerResp, error)
}

type workOrderSrv struct {
	pgRepo      pgRepo.Repository
	redisRepo   redisRepo.Repository
	customerSrv CustomerService
}

func NewWorkOrderService(pgRepo pgRepo.Repository, redisRepo redisRepo.Repository, customerSrv CustomerService) WorkOrderService {
	return &workOrderSrv{
		pgRepo,
		redisRepo,
		customerSrv,
	}
}

func (srv *workOrderSrv) CreateWorkOrder(ctx context.Context, req entity.CreateWorkOrderRequest) (*api.WorkOrder, error) {
	if err := utils.ValidatePlannedDates(req.PlannedDateBegin, req.PlannedDateEnd); err != nil {
		return &api.WorkOrder{}, err
	}

	customer, err := srv.customerSrv.GetCustomer(ctx, uuid.MustParse(req.CustomerID))
	workOrderModel := models.WorkOrder{}
	workOrderModel.BuildModelCreate(req)
	workOrder, err := srv.pgRepo.CreateWorkOrder(ctx, workOrderModel, &customer.IsActive)
	if err != nil {
		return &api.WorkOrder{}, err
	}

	return workOrder.ToDomainGRPC(), nil
}

func (srv *workOrderSrv) GetWorkOrderByID(ctx context.Context, ID uuid.UUID) (entity.WorkOrderResponse, error) {
	workOrder, err := srv.pgRepo.GetWorkOrderByID(ctx, ID)
	if err != nil {
		return entity.WorkOrderResponse{}, nil
	}

	if workOrder.ID == uuid.Nil {
		return entity.WorkOrderResponse{}, status.Errorf(codes.NotFound, "Orden de trabajo no encontrada")
	}

	return workOrder, err
}

func (srv *workOrderSrv) FinishWorkOrder(ctx context.Context, req entity.FinishWorkOrder) (*api.WorkOrder, error) {
	workOrder, err := srv.GetWorkOrderByID(ctx, req.WorkOrderID)
	if err != nil {
		return &api.WorkOrder{}, err
	}

	if workOrder.Status != "new" {
		return &api.WorkOrder{}, status.Errorf(codes.InvalidArgument, fmt.Sprintf("La orden de trabajo ya fue finalizada con estatus '%s'", workOrder.Status))
	}

	workOrderModel := models.WorkOrder{}
	workOrderModel.BuildModelFinish(workOrder)

	if workOrder, err = srv.pgRepo.FinishWorkOrder(ctx, workOrderModel); err != nil {
		return &api.WorkOrder{}, err
	}

	go srv.redisRepo.SendStreamLog(workOrder.ToDomainGRPC())

	return workOrder.ToDomainGRPC(), nil
}

func (srv *workOrderSrv) CancelWorkOrder(ctx context.Context, ID uuid.UUID) error {
	workOrder, err := srv.GetWorkOrderByID(ctx, ID)
	if err != nil {
		return err
	}

	if workOrder.Status != "new" {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("La orden de trabajo no se puede cancelar, su status actual es '%s'", workOrder.Status))
	}

	if err = srv.pgRepo.CancelWorkOrder(ctx, ID); err != nil {
		return err
	}

	return nil
}

func (srv *workOrderSrv) GetWorkOrders(ctx context.Context, req entity.GetWorkOrders) (*api.GetWorkOrdersResp, error) {
	workOrders, err := srv.pgRepo.GetWorkOrders(ctx, req)
	if err != nil {
		return &api.GetWorkOrdersResp{}, err
	}

	return workOrders.ToDomainGRPC(), nil
}

func (srv *workOrderSrv) GetWorkOrderByCustomerID(ctx context.Context, customerID uuid.UUID) (*api.GetWorkOrdersResp, error) {
	workOrders, err := srv.pgRepo.GetWorkOrderByCustomerID(ctx, customerID)
	if err != nil {
		return &api.GetWorkOrdersResp{}, err
	}

	return workOrders.ToDomainGRPC(), nil
}

func (srv *workOrderSrv) GetWorkOrderByIDWithCustomer(ctx context.Context, ID uuid.UUID) (*api.WorkOrderWithCustomerResp, error) {
	workOrder, err := srv.pgRepo.GetWorkOrderByID(ctx, ID)
	if err != nil {
		return &api.WorkOrderWithCustomerResp{}, err
	}

	customer, err := srv.pgRepo.GetCustomerByID(ctx, workOrder.CustomerID)
	if err != nil {
		return &api.WorkOrderWithCustomerResp{}, err
	}

	return &api.WorkOrderWithCustomerResp{
		WorkOrder: workOrder.ToDomainGRPC(),
		Customer:  customer.ToDomainGRPC(),
	}, nil
}

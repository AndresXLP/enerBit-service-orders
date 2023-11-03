package handler

import (
	"context"

	"enerBit-service-orders/api"
	"enerBit-service-orders/internal/entity"
	"enerBit-service-orders/internal/service"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcWorkOrder struct {
	api.UnimplementedWorkOrderServiceServer
	workOrderSrv service.WorkOrderService
}

func NewWorkOrder(workOrderSrv service.WorkOrderService) api.WorkOrderServiceServer {
	return &grpcWorkOrder{
		workOrderSrv: workOrderSrv,
	}
}

func (hand *grpcWorkOrder) Create(ctx context.Context, req *api.CreateWorkOrderReq) (*api.WorkOrder, error) {
	requestWorkOrder := entity.CreateWorkOrderRequest{}
	if err := requestWorkOrder.Bind(req); err != nil {
		return nil, err
	}

	workOrder, err := hand.workOrderSrv.CreateWorkOrder(ctx, requestWorkOrder)
	if err != nil {
		return &api.WorkOrder{}, err
	}

	return workOrder, nil
}

func (hand *grpcWorkOrder) Finish(ctx context.Context, req *api.WorkOrderReq) (*api.WorkOrder, error) {
	requestFinishWorkOrder := entity.FinishWorkOrder{}
	if err := requestFinishWorkOrder.Bind(req); err != nil {
		return &api.WorkOrder{}, err
	}

	workOrder, err := hand.workOrderSrv.FinishWorkOrder(ctx, requestFinishWorkOrder)
	if err != nil {
		return &api.WorkOrder{}, err
	}

	return workOrder, nil
}

func (hand *grpcWorkOrder) Cancel(ctx context.Context, req *api.WorkOrderReq) (*api.WorkOrderMsg, error) {
	id, err := uuid.Parse(req.GetWorkOrderId())
	if err != nil {
		return &api.WorkOrderMsg{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err = hand.workOrderSrv.CancelWorkOrder(ctx, id); err != nil {
		return &api.WorkOrderMsg{}, err
	}

	return &api.WorkOrderMsg{Message: "La orden de servicio fue cancelada con exito"}, nil
}

func (hand *grpcWorkOrder) GetWorkOrders(ctx context.Context, req *api.GetWorkOrdersReq) (*api.GetWorkOrdersResp, error) {
	requestWorkOrder := entity.GetWorkOrders{}
	requestWorkOrder.Bind(req)

	workOrders, err := hand.workOrderSrv.GetWorkOrders(ctx, requestWorkOrder)
	if err != nil {
		return &api.GetWorkOrdersResp{}, err
	}

	return workOrders, nil
}

func (hand *grpcWorkOrder) GetWorkOrderByCustomerID(ctx context.Context, req *api.GetWorkOrderByCustomerIDReq) (*api.GetWorkOrdersResp, error) {
	customerID, err := uuid.Parse(req.GetCustomerId())
	if err != nil {
		return &api.GetWorkOrdersResp{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	workOrders, err := hand.workOrderSrv.GetWorkOrderByCustomerID(ctx, customerID)
	if err != nil {
		return &api.GetWorkOrdersResp{}, err
	}

	return workOrders, nil
}

func (hand *grpcWorkOrder) GetWorkOrderByID(ctx context.Context, req *api.GetWorkOrderByIDReq) (*api.WorkOrderWithCustomerResp, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return &api.WorkOrderWithCustomerResp{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	workOrderWithCustomer, err := hand.workOrderSrv.GetWorkOrderByIDWithCustomer(ctx, id)
	if err != nil {
		return &api.WorkOrderWithCustomerResp{}, err
	}

	return workOrderWithCustomer, nil
}

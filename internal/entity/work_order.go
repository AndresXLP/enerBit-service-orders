package entity

import (
	"time"

	"enerBit-service-orders/api"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateWorkOrderRequest struct {
	CustomerID       string    `validate:"required"`
	Title            string    `validate:"required"`
	PlannedDateBegin time.Time `validate:"required"`
	PlannedDateEnd   time.Time `validate:"required"`
}

func (w *CreateWorkOrderRequest) Bind(req *api.CreateWorkOrderReq) error {
	w.CustomerID = req.GetCustomerId()
	w.Title = req.GetTitle()
	w.PlannedDateBegin = req.PlannedDateBegin.AsTime()
	w.PlannedDateEnd = req.PlannedDateEnd.AsTime()

	return validate.Struct(w)
}

type WorkOrderResponse struct {
	ID               uuid.UUID
	CustomerID       uuid.UUID
	Title            string
	PlannedDateBegin time.Time
	PlannedDateEnd   time.Time
	Status           string
	CreatedAt        time.Time
}

func (w *WorkOrderResponse) ToDomainGRPC() *api.WorkOrder {
	workOrder := &api.WorkOrder{
		Id:               w.ID.String(),
		CustomerId:       w.CustomerID.String(),
		Title:            w.Title,
		PlannedDateBegin: timestamppb.New(w.PlannedDateBegin),
		PlannedDateEnd:   timestamppb.New(w.PlannedDateEnd),
		CreatedAt:        timestamppb.New(w.CreatedAt),
	}

	switch w.Status {
	case "new":
		workOrder.Status = 0
		break
	case "done":
		workOrder.Status = 1
		break
	case "cancelled":
		workOrder.Status = 2
	default:
		workOrder.Status = -1
	}

	return workOrder
}

type FinishWorkOrder struct {
	WorkOrderID uuid.UUID `validate:"required"`
}

func (w *FinishWorkOrder) Bind(req *api.WorkOrderReq) error {
	id, err := uuid.Parse(req.WorkOrderId)
	if err != nil {
		return err
	}

	w.WorkOrderID = id

	return validate.Struct(w)
}

type GetWorkOrders struct {
	Since  time.Time
	Until  time.Time
	Status int32
}

func (g *GetWorkOrders) Bind(req *api.GetWorkOrdersReq) {
	status := req.Status.Number()

	if req.Since != nil && req.Until != nil {
		g.Since = req.Since.AsTime()
		g.Until = req.Until.AsTime()
	}

	g.Status = int32(status)
}

type GetWorkOrdersResponse []WorkOrderResponse

func (g *GetWorkOrdersResponse) Add(worksOrders ...WorkOrderResponse) {
	*g = append(*g, worksOrders...)
}

func (g *GetWorkOrdersResponse) ToDomainGRPC() *api.GetWorkOrdersResp {
	workOrders := make([]*api.WorkOrder, 0)

	for _, workOrder := range *g {
		workOrders = append(workOrders, workOrder.ToDomainGRPC())
	}

	return &api.GetWorkOrdersResp{WorkOrders: workOrders}
}

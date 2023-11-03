package models

import (
	"time"

	"enerBit-service-orders/internal/entity"
	"github.com/google/uuid"
)

type WorkOrder struct {
	ID               uuid.UUID
	CustomerID       uuid.UUID
	Title            string
	PlannedDateBegin time.Time
	PlannedDateEnd   time.Time
	Status           string
	CreatedAt        time.Time
}

func (w *WorkOrder) BuildModelCreate(req entity.CreateWorkOrderRequest) {
	w.ID = uuid.New()
	w.CustomerID = uuid.MustParse(req.CustomerID)
	w.Title = req.Title
	w.PlannedDateBegin = req.PlannedDateBegin
	w.PlannedDateEnd = req.PlannedDateEnd
	w.Status = "new"
	w.CreatedAt = time.Now()
}

func (w *WorkOrder) ToDomainEntity() entity.WorkOrderResponse {
	return entity.WorkOrderResponse{
		ID:               w.ID,
		CustomerID:       w.CustomerID,
		Title:            w.Title,
		PlannedDateBegin: w.PlannedDateBegin,
		PlannedDateEnd:   w.PlannedDateEnd,
		Status:           w.Status,
		CreatedAt:        w.CreatedAt,
	}
}

func (w *WorkOrder) BuildModelFinish(req entity.WorkOrderResponse) {
	w.ID = req.ID
	w.CustomerID = req.CustomerID
	w.Title = req.Title
	w.PlannedDateBegin = req.PlannedDateBegin
	w.PlannedDateEnd = req.PlannedDateEnd
	w.Status = req.Status
	w.CreatedAt = req.CreatedAt
}

type GetWorkOrders []WorkOrder

func (g *GetWorkOrders) ToDomainEntity() entity.GetWorkOrdersResponse {
	var worksOrders entity.GetWorkOrdersResponse

	for _, workOrder := range *g {
		worksOrders.Add(workOrder.ToDomainEntity())
	}

	return worksOrders
}

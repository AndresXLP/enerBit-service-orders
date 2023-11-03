package repo

import (
	"context"
	"time"

	"enerBit-service-orders/internal/entity"
	"enerBit-service-orders/internal/repos/postgres/models"
	"enerBit-service-orders/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCustomer(ctx context.Context, req models.Customer) (entity.CustomerResponse, error)
	GetActiveCustomers(ctx context.Context) (entity.ActiveCustomers, error)
	GetCustomer(ctx context.Context, ID uuid.UUID) (entity.CustomerResponse, error)
	CreateWorkOrder(ctx context.Context, req models.WorkOrder, firstOrder *bool) (entity.WorkOrderResponse, error)
	GetWorkOrderByID(ctx context.Context, ID uuid.UUID) (entity.WorkOrderResponse, error)
	FinishWorkOrder(ctx context.Context, req models.WorkOrder) (entity.WorkOrderResponse, error)
	CancelWorkOrder(ctx context.Context, ID uuid.UUID) error
	GetWorkOrders(ctx context.Context, orders entity.GetWorkOrders) (entity.GetWorkOrdersResponse, error)
	GetWorkOrderByCustomerID(ctx context.Context, customerID uuid.UUID) (entity.GetWorkOrdersResponse, error)
	GetCustomerByID(ctx context.Context, ID uuid.UUID) (entity.CustomerResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (repo repository) CreateCustomer(ctx context.Context, req models.Customer) (entity.CustomerResponse, error) {
	if err := repo.db.WithContext(ctx).
		Table("customers").
		Create(&req).
		Scan(&req).Error; err != nil {
		return entity.CustomerResponse{}, err
	}

	return req.ToDomainEntity(), nil
}

func (repo repository) GetActiveCustomers(ctx context.Context) (entity.ActiveCustomers, error) {
	var activeCustomers models.ActiveCustomers

	if err := repo.db.WithContext(ctx).
		Table("customers").
		Where("is_active = ?", true).
		Scan(&activeCustomers).
		Error; err != nil {
		return entity.ActiveCustomers{}, err
	}

	return activeCustomers.ToDomainEntity(), nil
}

func (repo repository) GetCustomer(ctx context.Context, ID uuid.UUID) (entity.CustomerResponse, error) {
	customer := models.Customer{}
	if err := repo.db.WithContext(ctx).
		Table("customers").
		Where("id = ?", ID).
		Scan(&customer).
		Error; err != nil {
		return entity.CustomerResponse{}, err
	}

	return customer.ToDomainEntity(), nil
}

func (repo repository) CreateWorkOrder(ctx context.Context, req models.WorkOrder, isActiveUser *bool) (entity.WorkOrderResponse, error) {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).
			Table("work_orders").
			Create(&req).
			Scan(&req).Error; err != nil {
			return err
		}

		tx = tx.WithContext(ctx).
			Table("customers").
			Where("id = ? ", req.CustomerID).
			Update("is_active", false)

		if *isActiveUser {
			if err := tx.Update("end_date", time.Now()).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return entity.WorkOrderResponse{}, err
	}

	return req.ToDomainEntity(), nil
}

func (repo repository) GetWorkOrderByID(ctx context.Context, ID uuid.UUID) (entity.WorkOrderResponse, error) {
	workOrder := models.WorkOrder{}

	if err := repo.db.WithContext(ctx).
		Table("work_orders").
		Where("id = ?", ID).
		Scan(&workOrder).
		Error; err != nil {
		return entity.WorkOrderResponse{}, err
	}

	return workOrder.ToDomainEntity(), nil
}

func (repo repository) FinishWorkOrder(ctx context.Context, req models.WorkOrder) (entity.WorkOrderResponse, error) {
	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).
			Table("work_orders").
			Where("id = ?", req.ID).
			Update("status", "done").
			Scan(&req).
			Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).
			Table("customers").
			Where("id = ?", req.CustomerID).
			Update("start_date", time.Now()).
			Update("is_active", true).
			Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return entity.WorkOrderResponse{}, err
	}

	return req.ToDomainEntity(), nil
}

func (repo repository) GetWorkOrders(ctx context.Context, req entity.GetWorkOrders) (entity.GetWorkOrdersResponse, error) {
	var workOrders models.GetWorkOrders

	query := ""

	utils.BuildQueryByDate(&query, req.Since, req.Until, "OR")
	utils.BuildQueryByStatus(&query, req.Status, "OR")

	if err := repo.db.WithContext(ctx).
		Table("work_orders").
		Where(query).
		Scan(&workOrders).
		Error; err != nil {
		return entity.GetWorkOrdersResponse{}, err
	}

	return workOrders.ToDomainEntity(), nil
}

func (repo repository) GetWorkOrderByCustomerID(ctx context.Context, customerID uuid.UUID) (entity.GetWorkOrdersResponse, error) {
	workOrders := models.GetWorkOrders{}
	if err := repo.db.WithContext(ctx).
		Table("work_orders").
		Where("customer_id = ?", customerID).
		Scan(&workOrders).
		Error; err != nil {
		return entity.GetWorkOrdersResponse{}, err
	}

	return workOrders.ToDomainEntity(), nil
}

func (repo repository) GetCustomerByID(ctx context.Context, ID uuid.UUID) (entity.CustomerResponse, error) {
	customer := models.Customer{}
	if err := repo.db.WithContext(ctx).
		Table("customers").
		Where("id = ?", ID).
		Scan(&customer).
		Error; err != nil {
		return entity.CustomerResponse{}, err
	}
	return customer.ToDomainEntity(), nil
}

func (repo repository) CancelWorkOrder(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.WithContext(ctx).
		Table("work_orders").
		Where("id = ?", ID).
		Update("status", "cancelled").
		Error; err != nil {
		return err
	}

	return nil
}

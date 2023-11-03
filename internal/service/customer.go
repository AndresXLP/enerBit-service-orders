package service

import (
	"context"

	"enerBit-service-orders/api"
	"enerBit-service-orders/internal/entity"
	"enerBit-service-orders/internal/repos/postgres/models"
	"enerBit-service-orders/internal/repos/postgres/repo"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerService interface {
	CreateCustomer(ctx context.Context, req entity.CreateCustomerRequest) (*api.Customer, error)
	GetActiveCustomers(ctx context.Context) (*api.ActiveCustomers, error)
	GetCustomer(ctx context.Context, ID uuid.UUID) (*api.Customer, error)
}

type customerSrv struct {
	repo repo.Repository
}

func NewCustomerService(repo repo.Repository) CustomerService {
	return &customerSrv{repo}
}

func (srv *customerSrv) CreateCustomer(ctx context.Context, req entity.CreateCustomerRequest) (*api.Customer, error) {
	customerModel := models.Customer{}
	customerModel.BuildModel(req)

	customer, err := srv.repo.CreateCustomer(ctx, customerModel)
	if err != nil {
		return &api.Customer{}, err
	}

	return customer.ToDomainGRPC(), err
}

func (srv *customerSrv) GetActiveCustomers(ctx context.Context) (*api.ActiveCustomers, error) {
	activeCustomers, err := srv.repo.GetActiveCustomers(ctx)
	if err != nil {
		return &api.ActiveCustomers{}, err
	}

	return activeCustomers.ToDomainGRPC(), nil
}

func (srv *customerSrv) GetCustomer(ctx context.Context, ID uuid.UUID) (*api.Customer, error) {
	customer, err := srv.repo.GetCustomer(ctx, ID)
	if err != nil {
		return &api.Customer{}, err
	}

	if customer.ID == uuid.Nil {
		return &api.Customer{}, status.Errorf(codes.NotFound, "Usuario relacionado a la orden de no encontrado")
	}

	return customer.ToDomainGRPC(), nil
}

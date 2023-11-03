package handler

import (
	"context"

	customers "enerBit-service-orders/api"
	"enerBit-service-orders/internal/entity"
	"enerBit-service-orders/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcCustomer struct {
	customers.UnimplementedCustomerServiceServer
	customerSrv service.CustomerService
}

func NewCustomer(customerSrv service.CustomerService) customers.CustomerServiceServer {
	return &grpcCustomer{
		customerSrv: customerSrv,
	}
}

func (hand *grpcCustomer) Create(ctx context.Context, req *customers.CreateCustomerReq) (*customers.Customer, error) {
	request := entity.CreateCustomerRequest{}
	if err := request.Bind(req); err != nil {
		return &customers.Customer{}, status.Errorf(codes.InvalidArgument, err.Error())
	}

	customer, err := hand.customerSrv.CreateCustomer(ctx, request)
	if err != nil {
		return &customers.Customer{}, err
	}

	return customer, nil
}

func (hand *grpcCustomer) GetActiveCustomers(ctx context.Context, _ *customers.GetActiveCustomersReq) (*customers.ActiveCustomers, error) {
	activeCustomers, err := hand.customerSrv.GetActiveCustomers(ctx)
	if err != nil {
		return &customers.ActiveCustomers{}, err
	}

	return activeCustomers, nil
}

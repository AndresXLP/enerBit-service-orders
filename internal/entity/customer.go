package entity

import (
	"time"

	"enerBit-service-orders/api"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	validate = validator.New()
)

type CreateCustomerRequest struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Address   string `validate:"required"`
}

func (c *CreateCustomerRequest) Bind(req *api.CreateCustomerReq) error {
	c.FirstName = req.GetFirstName()
	c.LastName = req.GetLastName()
	c.Address = req.GetAddress()

	return validate.Struct(c)
}

type CustomerResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string
	LastName  string
	Address   string
	StartDate *time.Time
	EndDate   *time.Time
	IsActive  *bool
	CreatedAt time.Time
}

func (c *CustomerResponse) ToDomainGRPC() *api.Customer {
	customer := &api.Customer{
		Id:        c.ID.String(),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address:   c.Address,
		IsActive:  *c.IsActive,
		CreatedAt: timestamppb.New(c.CreatedAt),
	}

	if c.StartDate != nil {
		customer.StartDate = timestamppb.New(*c.StartDate)
	}

	if c.EndDate != nil {
		customer.EndDate = timestamppb.New(*c.EndDate)
	}

	return customer
}

func (c *ActiveCustomers) Add(customer ...CustomerResponse) {
	*c = append(*c, customer...)
}

type ActiveCustomers []CustomerResponse

func (c *ActiveCustomers) ToDomainGRPC() *api.ActiveCustomers {
	activeCustomers := make([]*api.Customer, 0)

	for _, customer := range *c {
		activeCustomers = append(activeCustomers, customer.ToDomainGRPC())
	}
	return &api.ActiveCustomers{Customers: activeCustomers}
}

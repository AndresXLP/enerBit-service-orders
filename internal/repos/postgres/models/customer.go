package models

import (
	"time"

	"enerBit-service-orders/internal/entity"
	"github.com/google/uuid"
)

var (
	isTrue  = true
	isFalse = false
)

type Customer struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Address   string
	StartDate *time.Time
	EndDate   *time.Time
	IsActive  *bool
	CreatedAt time.Time
}

func (c *Customer) BuildModel(req entity.CreateCustomerRequest) {
	c.ID = uuid.New()
	c.FirstName = req.FirstName
	c.LastName = req.LastName
	c.Address = req.Address
	c.IsActive = &isFalse
	c.StartDate = nil
	c.EndDate = nil
	c.CreatedAt = time.Now()
}

func (c *Customer) ToDomainEntity() entity.CustomerResponse {
	return entity.CustomerResponse{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address:   c.Address,
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
		IsActive:  c.IsActive,
		CreatedAt: c.CreatedAt,
	}

}

type ActiveCustomers []Customer

func (a *ActiveCustomers) ToDomainEntity() entity.ActiveCustomers {
	var activeCustomers entity.ActiveCustomers

	for _, customer := range *a {
		activeCustomers.Add(customer.ToDomainEntity())
	}

	return activeCustomers
}

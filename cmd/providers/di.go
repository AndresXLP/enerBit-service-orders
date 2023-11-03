package providers

import (
	"enerBit-service-orders/internal/handler"
	"enerBit-service-orders/internal/repos/postgres"
	pgRepo "enerBit-service-orders/internal/repos/postgres/repo"
	"enerBit-service-orders/internal/repos/redis"
	redisRepo "enerBit-service-orders/internal/repos/redis/repo"
	"enerBit-service-orders/internal/server"
	"enerBit-service-orders/internal/service"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(server.NewGRPCServer)

	_ = Container.Provide(handler.NewCustomer)
	_ = Container.Provide(handler.NewWorkOrder)

	_ = Container.Provide(service.NewCustomerService)
	_ = Container.Provide(service.NewWorkOrderService)

	_ = Container.Provide(postgres.NewConnection)
	_ = Container.Provide(redis.NewConnection)

	_ = Container.Provide(pgRepo.NewPostgresRepository)
	_ = Container.Provide(redisRepo.NewRedisRepository)

	return Container
}

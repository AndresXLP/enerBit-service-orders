package server

import (
	"context"
	"fmt"
	"net"

	"enerBit-service-orders/api"
	"enerBit-service-orders/config"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

type Server interface {
	Serve()
}

type connection struct {
	host             string
	port             int
	protocol         string
	customerService  api.CustomerServiceServer
	workOrderService api.WorkOrderServiceServer
}

func NewGRPCServer(customerService api.CustomerServiceServer, workOrderService api.WorkOrderServiceServer) Server {
	return &connection{
		config.Environments().GRPC.Host,
		config.Environments().GRPC.Port,
		config.Environments().GRPC.Protocol,
		customerService,
		workOrderService,
	}
}

func (c *connection) Serve() {
	addr := fmt.Sprintf("%s:%d", c.host, c.port)

	listener, err := net.Listen(c.protocol, addr)
	if err != nil {
		log.Panicf("Error listen api Server %v\n", err)
	}

	gRPC := grpc.NewServer(
		grpc.UnaryInterceptor(recoveryInterceptor()))
	api.RegisterCustomerServiceServer(gRPC, c.customerService)
	api.RegisterWorkOrderServiceServer(gRPC, c.workOrderService)

	fmt.Printf(color.Magenta("⇋ api Server Running on %s %s \n"), color.Red(c.protocol), color.Green(listener.Addr()))
	log.Fatal(gRPC.Serve(listener))
}

func recoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Error(r)
			}
		}()

		return handler(ctx, req)
	}
}

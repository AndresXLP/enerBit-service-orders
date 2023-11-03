package main

import (
	"enerBit-service-orders/cmd/providers"
	"enerBit-service-orders/internal/server"
	"github.com/labstack/gommon/log"
)

func main() {
	container := providers.BuildContainer()

	if err := container.Invoke(func(grpcServer server.Server) {
		grpcServer.Serve()
	}); err != nil {
		log.Panic(err)
	}
}

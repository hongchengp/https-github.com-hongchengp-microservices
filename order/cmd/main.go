package main

import (
	"fmt"
	"log"

	"github.com/hongchengp/microservices/order/config"
	"github.com/hongchengp/microservices/order/internal/adapters/db"
	"github.com/hongchengp/microservices/order/internal/adapters/grpc"
	"github.com/hongchengp/microservices/order/internal/application/core/api"
)

// injecting & running service
func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	fmt.Println(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicaionPort())
	grpcAdapter.Run()
}
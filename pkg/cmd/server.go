package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/adomokos/go-grpc-http-rest-microservice-tutorial/pkg/protocol/grpc"
	"github.com/adomokos/go-grpc-http-rest-microservice-tutorial/pkg/service/v1"
)

type Config struct {
	GRPCPort        string
	DatastoreDBFile string
}

func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBFile, "db-file", "", "Database file path")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Error("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	db, err = &toDoServiceServer.Open("sqlite3", cfg.DatastoreDBFile)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	v1API := v1.NewToDoServiceServer(db)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}

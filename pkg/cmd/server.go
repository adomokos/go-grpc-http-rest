package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/adomokos/go-grpc-http-rest/pkg/protocol/grpc"
	"github.com/adomokos/go-grpc-http-rest/pkg/service/v1"
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
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	db, err := gorm.Open("sqlite3", cfg.DatastoreDBFile)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()
	db.LogMode(true)

	v1API := v1.NewToDoServiceServer(db)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}

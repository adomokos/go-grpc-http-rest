package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/adomokos/go-grpc-http-rest/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

const (
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create
	id := callCreate(ctx, &c)

	// Read
	callRead(ctx, c, id)
}

func callCreate(ctx context.Context, c *v1.ToDoServiceClient) int64 {
	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	pfx := t.Format(time.RFC3339Nano)

	// Call Create
	req1 := v1.CreateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Title:       "title (" + pfx + ")",
			Description: "description (" + pfx + ")",
			Reminder:    reminder,
		},
	}

	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	return res1.Id
}

func callRead(ctx context.Context, c v1.ToDoServiceClient, id int64) *v1.ToDo {
	req2 := v1.ReadRequest{
		Api: apiVersion,
		Id:  id,
	}

	res2, err := c.Read(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}

	log.Printf("Read result: <%+v>\n\n", res2)

	return res2.ToDo
}

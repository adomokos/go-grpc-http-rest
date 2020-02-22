package v1

import (
	"context"
	"errors"
	// "fmt"
	// "time"

	// "github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/adomokos/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
)

const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
	db *gorm.DB
}

func NewToDoServiceServer(db *gorm.DB) v1.ToDoServiceServer {
	return &toDoServiceServer{db: db}
}

func (s *toDoServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// func (s *toDoServiceServer) connect(ctx context.Context) (*gorm.DB, error) {
// db, err = &toDoServiceServer.Open("sqlite3", "db/todo-db.sqlt")
// if err != nil {
// return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
// }

// return db, nil
// }

func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *toDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *toDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *toDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	return nil, errors.New("Not implemented")
}

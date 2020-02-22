package v1

import (
	"context"
	"errors"
	"time"

	// "github.com/golang/protobuf/ptypes"
	"github.com/adomokos/go-grpc-http-rest/pkg/api/v1"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ToDoEntity struct {
	ID          uint
	Title       string
	Description string
	Reminder    time.Time
}

func (ToDoEntity) TableName() string {
	return "todos"
}

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
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	todoEntity := convertTodo(req.ToDo)
	s.db.Create(&todoEntity)

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  int64(todoEntity.ID),
	}, nil
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

func convertTodo(todo *v1.ToDo) *ToDoEntity {
	unixTimeUTC := time.Unix(todo.Reminder.GetSeconds(), 0)

	todoEntity := ToDoEntity{
		Title:       todo.Title,
		Description: todo.Description,
		Reminder:    unixTimeUTC,
	}

	return &todoEntity
}

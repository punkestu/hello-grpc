package service

import (
	"context"
	"encoding/json"
	"errors"
	userProto "github.com/punkestu/hello-grpc/proto"
	"log"
	"os"
	"sync"
)

// DataUserService implement proto server
type DataUserService struct {
	userProto.UnimplementedUserServiceServer
	mu    sync.Mutex
	users []*userProto.User
}

// GetByEmail follow the protobuf file
func (dus *DataUserService) GetByEmail(ctx context.Context, user *userProto.User) (*userProto.User, error) {
	for _, User := range dus.users {
		if User.Email == user.Email {
			return User, nil
		}
	}
	return nil, errors.New("user not found")
}

// load mock data
func (dus *DataUserService) loadDummy() {
	data, err := os.ReadFile("mock/users.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := json.Unmarshal(data, &dus.users); err != nil {
		log.Fatalln(err.Error())
	}
}

func NewUserService() *DataUserService {
	service := DataUserService{}
	service.loadDummy()
	return &service
}

package handler

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/service"
	pb "github.com/mellgit/someuser/internal/transport/grpc/proto"
	"golang.org/x/net/context"
)

type SomeUserGRPC struct {
	pb.UnimplementedSomeUserServer
	service service.Service
}

func NewSomeUserGRPC(service service.Service) *SomeUserGRPC {
	return &SomeUserGRPC{
		service: service,
	}
}

func (s *SomeUserGRPC) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	payload := model.CreateUserRequest{
		Username: in.GetName(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}
	user, err := s.service.CreateUser(payload)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateUserResponse{
		Id:       user.ID.(string),
		Name:     user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	return resp, nil
}

func (s *SomeUserGRPC) GetAllUsers(_ context.Context, _ *empty.Empty) (*pb.GetAllUsersResponse, error) {
	users, err := s.service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	resp := &pb.GetAllUsersResponse{
		SchemaSomeUsers: make([]*pb.CreateUserResponse, 0),
	}
	for _, val := range *users {
		resp.SchemaSomeUsers = append(resp.SchemaSomeUsers, &pb.CreateUserResponse{
			Name:     val.Username,
			Email:    val.Email,
			Password: val.Password,
		})
	}
	return resp, nil
}

func (s *SomeUserGRPC) GetUserByID(ctx context.Context, in *pb.UserIDRequest) (*pb.CreateUserResponse, error) {
	user, err := s.service.GetUserByID(in.GetId())
	if err != nil {
		return nil, err
	}
	resp := &pb.CreateUserResponse{
		Id:       user.ID.(string),
		Name:     user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	return resp, nil
}

func (s *SomeUserGRPC) DeleteUserByID(ctx context.Context, in *pb.UserIDRequest) (*pb.MessageResponse, error) {
	err := s.service.DeleteUserByID(in.GetId())
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("User deleted with id: %s", in.GetId())
	resp := &pb.MessageResponse{
		Message: msg,
	}
	return resp, nil
}

func (s *SomeUserGRPC) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.CreateUserResponse, error) {
	payload := model.UpdateUserRequest{
		Username: in.GetUpdateBody().GetName(),
		Email:    in.GetUpdateBody().GetEmail(),
		Password: in.GetUpdateBody().GetPassword(),
	}

	updateUser, err := s.service.UpdateUser(in.GetId(), payload)
	if err != nil {
		return nil, err
	}
	resp := &pb.CreateUserResponse{
		Id:       updateUser.ID.(string),
		Name:     updateUser.Username,
		Email:    updateUser.Email,
		Password: updateUser.Password,
	}
	return resp, nil
}

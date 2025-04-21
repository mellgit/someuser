package handler

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/service"
	pb "github.com/mellgit/someuser/internal/transport/grpc/proto"
	"github.com/mellgit/someuser/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type SomeUserGRPC struct {
	pb.UnimplementedSomeUserServer
	service service.Service
	Logger  *log.Entry
}

func NewSomeUserGRPC(service service.Service, logger *log.Entry) *SomeUserGRPC {
	return &SomeUserGRPC{
		service: service,
		Logger:  logger,
	}
}

func (h *SomeUserGRPC) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	payload := model.CreateUserRequest{
		Username: in.GetUsername(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}
	if err := utils.ValidateStruct(payload); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "CreateUser",
		}).Errorf("%v", err)
		return nil, fmt.Errorf("validate user: %v", err)
	}

	user, err := h.service.CreateUser(payload)
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "CreateUser",
		}).Errorf("%v", err)
		return nil, fmt.Errorf("create user: %v", err)
	}

	resp := &pb.CreateUserResponse{
		Id:       user.ID.(string),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	return resp, nil
}

func (h *SomeUserGRPC) GetAllUsers(_ context.Context, _ *empty.Empty) (*pb.GetAllUsersResponse, error) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "GetAllUsers",
		}).Errorf("%v", err)
		return nil, fmt.Errorf("get all users: %v", err)
	}
	resp := &pb.GetAllUsersResponse{
		SchemaSomeUsers: make([]*pb.CreateUserResponse, 0),
	}
	for _, val := range *users {
		resp.SchemaSomeUsers = append(resp.SchemaSomeUsers, &pb.CreateUserResponse{
			Id:       val.ID.(string),
			Username: val.Username,
			Email:    val.Email,
			Password: val.Password,
		})
	}
	return resp, nil
}

func (h *SomeUserGRPC) GetUserByID(ctx context.Context, in *pb.UserIDRequest) (*pb.CreateUserResponse, error) {
	user, err := h.service.GetUserByID(in.GetId())
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "GetUserByID",
		}).Errorf("%v", err)
		return nil, fmt.Errorf("get user by id: %v", err)
	}
	resp := &pb.CreateUserResponse{
		Id:       user.ID.(string),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	return resp, nil
}

func (h *SomeUserGRPC) DeleteUserByID(ctx context.Context, in *pb.UserIDRequest) (*pb.MessageResponse, error) {
	err := h.service.DeleteUserByID(in.GetId())
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "DeleteUserByID",
		}).Errorf("%v", err)
		return nil, fmt.Errorf("delete user by id: %v", err)
	}
	msg := fmt.Sprintf("User deleted with id: %s", in.GetId())
	resp := &pb.MessageResponse{
		Message: msg,
	}
	return resp, nil
}

func (h *SomeUserGRPC) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.CreateUserResponse, error) {
	payload := model.UpdateUserRequest{
		Username: in.GetUpdateBody().GetUsername(),
		Email:    in.GetUpdateBody().GetEmail(),
		Password: in.GetUpdateBody().GetPassword(),
	}

	if err := utils.ValidateStruct(payload); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "UpdateUser",
		}).Errorf("%v", err)
		return nil, fmt.Errorf("validate user: %v", err)
	}

	updateUser, err := h.service.UpdateUser(in.GetId(), payload)
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "UpdateUser",
		}).Errorf("%v", err)
		return nil, fmt.Errorf("update user: %v", err)
	}
	resp := &pb.CreateUserResponse{
		Id:       updateUser.ID.(string),
		Username: updateUser.Username,
		Email:    updateUser.Email,
		Password: updateUser.Password,
	}
	return resp, nil
}

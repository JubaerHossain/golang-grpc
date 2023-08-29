package handlers

import (
	"context"
	"errors"

	users "github.com/JubaerHossain/golang-grpc/api/proto/users"
)

type UsersServiceHandler struct {
	users map[string]*users.User
}

func NewUsersServiceHandler() *UsersServiceHandler {
	return &UsersServiceHandler{
		users: make(map[string]*users.User),
	}
}

func (h *UsersServiceHandler) CreateUser(ctx context.Context, req *users.User) (*users.UserResponse, error) {
	if _, exists := h.users[req.Id]; exists {
		return nil, errors.New("user already exists")
	}
	h.users[req.Id] = req
	return &users.UserResponse{User: req}, nil
}

func (h *UsersServiceHandler) ReadUser(ctx context.Context, req *users.UserId) (*users.User, error) {
	user, exists := h.users[req.Id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (h *UsersServiceHandler) UpdateUser(ctx context.Context, req *users.User) (*users.UserResponse, error) {
	if _, exists := h.users[req.Id]; !exists {
		return nil, errors.New("user not found")
	}
	h.users[req.Id] = req
	return &users.UserResponse{User: req}, nil
}

func (h *UsersServiceHandler) DeleteUser(ctx context.Context, req *users.UserId) (*users.User, error) {
	user, exists := h.users[req.Id]
	if !exists {
		return nil, errors.New("user not found")
	}
	delete(h.users, req.Id)
	return user, nil
}

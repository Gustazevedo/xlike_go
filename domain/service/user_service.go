package service

import (
    "github.com/Gustazevedo/xlike_go/tree/develop/domain/model"
    "github.com/Gustazevedo/xlike_go/tree/develop/ports"
    "errors"
)

type UserService struct {
    repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
    return &UserService{
        repo: repo,
    }
}

func (s *UserService) RegisterUser(username, email string) (*model.User, error) {
    existingUser, _ := s.repo.GetUserByUsername(username)
    if existingUser != nil {
        return nil, errors.New("username already exists")
    }

    user := model.User{
        Username: username,
        Email:    email,
    }

    err := s.repo.SaveUser(user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

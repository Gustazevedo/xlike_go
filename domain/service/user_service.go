package service

import "github.com/Gustazevedo/xlike_go/domain/model"

type UserService struct {}

func (s *UserService) RegisterUser(username, email string) model.User {
    return model.User{
        ID:       "random-generated-id",
        Username: username,
        Email:    email,
    }
}

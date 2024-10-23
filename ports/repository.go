package ports

import "github.com/Gustazevedo/xlike_go/tree/develop/domain/model"

type UserRepository interface {
    SaveUser(user model.User) error
    GetUserByUsername(username string) (*model.User, error)
}

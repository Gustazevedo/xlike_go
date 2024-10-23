package database

import (
    "context"
    "github.com/Gustazevedo/xlike_go/tree/develop/domain/model"
    "github.com/Gustazevedo/xlike_go/tree/develop/ports"
    "github.com/jackc/pgx/v4/pgxpool"
)

type PostgresUserRepository struct {
    db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) ports.UserRepository {
    return &PostgresUserRepository{
        db: db,
    }
}

func (r *PostgresUserRepository) SaveUser(user model.User) error {
    _, err := r.db.Exec(context.Background(), "INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
    return err
}

func (r *PostgresUserRepository) GetUserByUsername(username string) (*model.User, error) {
    var user model.User
    err := r.db.QueryRow(context.Background(), "SELECT id, username, email FROM users WHERE username=$1", username).Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

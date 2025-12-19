package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/varun053101/GO-TASK/internal/repository/sqlc"
)

// user repository wraps sqlc queries
type UserRepository struct {
	q *sqlc.Queries
}

// create a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		q: sqlc.New(DB),
	}
}

// create a new user record
func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	dob pgtype.Date,
) (sqlc.User, error) {
	return r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

// get user by id
func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int32,
) (sqlc.User, error) {
	return r.q.GetUserByID(ctx, id)
}

// update user by id
func (r *UserRepository) UpdateUser(
	ctx context.Context,
	id int32,
	name string,
	dob pgtype.Date,
) (sqlc.User, error) {
	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}
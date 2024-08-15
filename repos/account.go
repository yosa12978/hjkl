package repos

import (
	"context"
	"database/sql"

	"github.com/yosa12978/hjkl/types"
)

type AccountRepo interface {
	GetById(ctx context.Context, id string) (*types.Account, error)
	GetByCredentials(ctx context.Context, username, password string) (*types.Account, error)
	GetByUsername(ctx context.Context, username string) (*types.Account, error)

	Create(ctx context.Context, acc types.Account) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, acc types.Account) error
}

type accountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) AccountRepo {
	return &accountRepo{
		db: db,
	}
}

// Create implements AccountRepo.
func (a *accountRepo) Create(ctx context.Context, acc types.Account) error {
	panic("unimplemented")
}

// Delete implements AccountRepo.
func (a *accountRepo) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetByCredentials implements AccountRepo.
func (a *accountRepo) GetByCredentials(ctx context.Context, username string, password string) (*types.Account, error) {
	panic("unimplemented")
}

// GetById implements AccountRepo.
func (a *accountRepo) GetById(ctx context.Context, id string) (*types.Account, error) {
	panic("unimplemented")
}

// GetByUsername implements AccountRepo.
func (a *accountRepo) GetByUsername(ctx context.Context, username string) (*types.Account, error) {
	panic("unimplemented")
}

// Update implements AccountRepo.
func (a *accountRepo) Update(ctx context.Context, id string, acc types.Account) error {
	panic("unimplemented")
}

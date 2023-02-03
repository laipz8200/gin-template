package uow

import (
	"_template_/domain/repository"
	"_template_/model"
	"context"

	"gorm.io/gorm"
)

type UnitOfWorkBlock func(repository.Repository) error

type UnitOfWork interface {
	Atomic(context.Context, UnitOfWorkBlock) error
}

var _ UnitOfWork = (*unitOfWork)(nil)

type unitOfWork struct {
	conn *gorm.DB
}

// Atomic implements UnitOfWork
func (u *unitOfWork) Atomic(ctx context.Context, fn UnitOfWorkBlock) error {
	return u.conn.Transaction(func(tx *gorm.DB) error {
		// Init new repository with tx
		newRepository := model.NewRepository(
			model.NewUserRepository(tx),
		)
		return fn(newRepository)
	})
}

func New(conn *gorm.DB) UnitOfWork {
	return &unitOfWork{
		conn: conn,
	}
}

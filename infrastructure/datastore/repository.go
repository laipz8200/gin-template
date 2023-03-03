package datastore

import (
	"_template_/domain"
	"context"

	"gorm.io/gorm"
)

var _ domain.IRepository = (*repository)(nil)

type repository struct {
	db *gorm.DB
}

// Atomic implements domain.IRepository
func (r *repository) Atomic(ctx context.Context, fn func(domain.IRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		repo := NewRepository(tx)
		return fn(repo)
	})
}

func NewRepository(db *gorm.DB) domain.IRepository {
	return &repository{
		db: db,
	}
}

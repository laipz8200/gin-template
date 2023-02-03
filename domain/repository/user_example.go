package repository

import (
	"_template_/domain/entity"
	"context"
)

type IUsers interface {
	FindOne(context.Context, string) (entity.User, error)
	Save(context.Context, *entity.User) error
}

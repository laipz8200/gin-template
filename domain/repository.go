package domain

import "context"

type IRepository interface {
	Atomic(context.Context, func(IRepository) error) error
}

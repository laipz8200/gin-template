package domain

var registry IRegistry

type IRegistry interface {
	Repository() IRepository
}

func Registry() IRegistry {
	return registry
}

func SetRegistry(r IRegistry) {
	registry = r
}

func InitRegistry() IRegistry {
	return nil
}

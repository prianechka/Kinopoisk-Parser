package domain

import "context"

type Person struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Height   uint64 `json:"height,omitempty"`
	Age      uint64 `json:"age,omitempty"`
}

type PersonRepository interface {
	GetByID(ctx context.Context, id uint64) (Person, error)
	GetByFullName(ctx context.Context, title string) (Person, error)
	GetPersons(ctx context.Context, limit, offset uint64) ([]Person, error)
	Add(ctx context.Context, m *Person) error
	Delete(ctx context.Context, id uint64) error
}

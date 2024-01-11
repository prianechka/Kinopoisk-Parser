package domain

import "context"

const (
	DirectorRole uint64 = 1
	ProducerRole uint64 = 2
	WriterRole   uint64 = 3
	ActorsRole   uint64 = 4
)

type Profession struct {
	ID       uint64 `json:"id"`
	MovieID  uint64 `json:"movie_id"`
	PersonID uint64 `json:"person_id"`
	Role     uint64 `json:"role"`
}

type ProfessionRepository interface {
	GetIDByParams(ctx context.Context, movieID, personID, role uint64) (uint64, error)
	GetByID(ctx context.Context, id uint64) (Profession, error)
	GetDirectorsByMovie(ctx context.Context, movieID uint64) ([]Profession, error)
	GetProducersByMovie(ctx context.Context, movieID uint64) ([]Profession, error)
	GetWritersByMovie(ctx context.Context, movieID uint64) ([]Profession, error)
	GetActorsByMovie(ctx context.Context, movieID uint64) ([]Profession, error)
	Add(ctx context.Context, movieID, personID, role uint64) error
	Delete(ctx context.Context, id uint64) error
}

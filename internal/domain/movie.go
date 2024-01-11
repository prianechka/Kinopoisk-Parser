package domain

import "context"

type MovieBaseInfo struct {
	ID       uint64  `json:"id"`
	Title    string  `json:"title"`
	Year     uint64  `json:"year"`
	Tagline  string  `json:"tagline"`
	Duration uint64  `json:"duration"`
	Rating   float64 `json:"rating"`
	Budget   uint64  `json:"budget"`
	Gross    uint64  `json:"gross"`
}

type Movie struct {
	BaseInfo  MovieBaseInfo `json:"info"`
	Producers []Person      `json:"producer"`
	Directors []Person      `json:"directors"`
	Actors    []Person      `json:"actors"`
	Writers   []Person      `json:"writers"`
}

type MovieRepoDTO struct {
	Movie     Movie
	Country   string
	Directors string
	Producers string
	Writers   string
	Actors    string
}

type MovieUsecase interface {
	GetByID(ctx context.Context, id uint64) (Movie, error)
	GetByTitle(ctx context.Context, title string) (Movie, error)
	GetMovies(ctx context.Context, limit, offset uint64) ([]Movie, error)
	Add(ctx context.Context, m *Movie) error
	Delete(ctx context.Context, id uint64) error
}

type MovieRepository interface {
	GetByID(ctx context.Context, id uint64) (MovieBaseInfo, error)
	GetByTitle(ctx context.Context, title string) (MovieBaseInfo, error)
	GetMovies(ctx context.Context, limit, offset uint64) ([]MovieBaseInfo, error)
	Add(ctx context.Context, m *MovieBaseInfo) error
	Delete(ctx context.Context, id uint64) error
}

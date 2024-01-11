package postgresql

import (
	"Kinopoisk-Parser/internal/domain"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type pgMovieRepo struct {
	Conn *sql.DB
}

func New(conn *sql.DB) domain.MovieRepository {
	return &pgMovieRepo{Conn: conn}
}

func (r pgMovieRepo) GetByID(ctx context.Context, id uint64) (domain.MovieBaseInfo, error) {
	query := `SELECT id, title, movie_year, tagline, duration, rating, budget, gross FROM movie WHERE id = $1;`

	rows, err := r.Conn.QueryContext(ctx, query, id)
	if err != nil {
		logrus.Errorf("Repo error: %v", err)
		return domain.MovieBaseInfo{}, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("Repo closing error: %v", err)
		}
	}()

	movie := domain.MovieBaseInfo{}
	err = rows.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Year,
		&movie.Tagline,
		&movie.Duration,
		&movie.Rating,
		&movie.Budget,
		&movie.Gross)

	if err != nil {
		logrus.Errorf("Repo error: %v", err)
		return domain.MovieBaseInfo{}, err
	}

	return movie, nil
}

func (r pgMovieRepo) GetByTitle(ctx context.Context, title string) (domain.MovieBaseInfo, error) {
	query := `SELECT id, title, movie_year, tagline, duration, rating, budget, gross FROM movie WHERE title = $1;`

	rows, err := r.Conn.QueryContext(ctx, query, title)
	if err != nil {
		logrus.Errorf("Repo error: %v", err)
		return domain.MovieBaseInfo{}, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("Repo closing error: %v", err)
		}
	}()

	movie := domain.MovieBaseInfo{}
	if rows.Next() {
		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Year,
			&movie.Tagline,
			&movie.Duration,
			&movie.Rating,
			&movie.Budget,
			&movie.Gross)
	} else {
		err = domain.MovieNotFound
	}

	if err != nil {
		logrus.Errorf("Repo error: %v", err)
		return domain.MovieBaseInfo{}, err
	}

	return movie, nil
}

func (r pgMovieRepo) GetMovies(ctx context.Context, limit, offset uint64) ([]domain.MovieBaseInfo, error) {
	query := `SELECT id, title, movie_year, tagline, duration, rating, budget, gross FROM movie LIMIT $1 OFFSET $2;`

	rows, err := r.Conn.QueryContext(ctx, query, limit, offset)

	if err != nil {
		logrus.Errorf("Repo error: %v", err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("Repo closing error: %v", err)
		}
	}()

	result := make([]domain.MovieBaseInfo, 0)
	for rows.Next() {
		tmpMovie := domain.MovieBaseInfo{}
		err = rows.Scan(
			&tmpMovie.ID,
			&tmpMovie.Title,
			&tmpMovie.Year,
			&tmpMovie.Tagline,
			&tmpMovie.Duration,
			&tmpMovie.Rating,
			&tmpMovie.Budget,
			&tmpMovie.Gross)

		if err != nil {
			logrus.Errorf("Repo error: %v", err)
			return nil, err
		}

		result = append(result, tmpMovie)
	}

	return result, nil
}

func (r pgMovieRepo) Add(ctx context.Context, m *domain.MovieBaseInfo) error {
	query := `INSERT into movie(id, title, movie_year, tagline, duration, rating, budget, gross) VALUES 
			 ($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err := r.Conn.ExecContext(ctx, query, m.ID, m.Title, m.Year, m.Tagline, m.Duration,
		m.Rating, m.Budget, m.Gross)

	return err
}

func (r pgMovieRepo) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM movie WHERE id = $1;`

	_, err := r.Conn.ExecContext(ctx, query, id)
	return err
}

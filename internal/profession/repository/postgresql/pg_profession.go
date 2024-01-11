package postgresql

import (
	"Kinopoisk-Parser/internal/domain"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type pgProfessionalRepo struct {
	Conn *sql.DB
}

func New(conn *sql.DB) domain.ProfessionRepository {
	return &pgProfessionalRepo{Conn: conn}
}

func (p pgProfessionalRepo) GetIDByParams(ctx context.Context, movieID, personID, role uint64) (uint64, error) {
	query := `SELECT id, movie_id, person_id, movie_role FROM professions WHERE movie_id = $1 AND person_id = $2 AND movie_role = $3;`

	rows, err := p.Conn.QueryContext(ctx, query, movieID, personID, role)
	if err != nil {
		logrus.Errorf("Repo error: %v", err)
		return 0, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("Repo closing error: %v", err)
		}
	}()

	tmpProfession := domain.Profession{}
	if rows.Next() {
		err = rows.Scan(
			&tmpProfession.ID,
			&tmpProfession.MovieID,
			&tmpProfession.PersonID,
			&tmpProfession.Role)

		if err != nil {
			logrus.Errorf("Repo error: %v", err)
			return 0, err
		}
	} else {
		err = domain.ProfessionNotFound
	}

	if err != nil {
		return 0, err
	}

	return tmpProfession.ID, err
}

func (p pgProfessionalRepo) GetByID(ctx context.Context, id uint64) (domain.Profession, error) {
	query := `SELECT id, movie_id, person_id, movie_role FROM professions WHERE id = $1;`

	rows, err := p.Conn.QueryContext(ctx, query, id)
	if err != nil {
		logrus.Errorf("Repo error: %v", err)
		return domain.Profession{}, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("Repo closing error: %v", err)
		}
	}()

	tmpProfession := domain.Profession{}
	if rows.Next() {
		err = rows.Scan(
			&tmpProfession.ID,
			&tmpProfession.MovieID,
			&tmpProfession.PersonID,
			&tmpProfession.Role)

		if err != nil {
			logrus.Errorf("Repo error: %v", err)
			return domain.Profession{}, err
		}
	} else {
		err = domain.ProfessionNotFound
	}

	if err != nil {
		return domain.Profession{}, err
	}

	return tmpProfession, err
}

func (p pgProfessionalRepo) GetDirectorsByMovie(ctx context.Context, movieID uint64) ([]domain.Profession, error) {
	query := `SELECT id, movie_id, person_id, movie_role FROM professions WHERE movie_id = $1 AND movie_role = $2;`

	rows, err := p.Conn.QueryContext(ctx, query, movieID, domain.DirectorRole)
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

	result := make([]domain.Profession, 0)
	for rows.Next() {
		tmpProfession := domain.Profession{}
		err = rows.Scan(
			&tmpProfession.ID,
			&tmpProfession.MovieID,
			&tmpProfession.PersonID,
			&tmpProfession.Role)

		if err != nil {
			logrus.Errorf("Repo error: %v", err)
			return nil, err
		}

		result = append(result, tmpProfession)
	}

	return result, err
}

func (p pgProfessionalRepo) GetProducersByMovie(ctx context.Context, movieID uint64) ([]domain.Profession, error) {
	query := `SELECT id, movie_id, person_id, movie_role FROM professions WHERE movie_id = $1 AND movie_role = $2;`

	rows, err := p.Conn.QueryContext(ctx, query, movieID, domain.ProducerRole)
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

	result := make([]domain.Profession, 0)
	for rows.Next() {
		tmpProfession := domain.Profession{}
		err = rows.Scan(
			&tmpProfession.ID,
			&tmpProfession.MovieID,
			&tmpProfession.PersonID,
			&tmpProfession.Role)

		if err != nil {
			logrus.Errorf("Repo error: %v", err)
			return nil, err
		}

		result = append(result, tmpProfession)
	}

	return result, err
}

func (p pgProfessionalRepo) GetWritersByMovie(ctx context.Context, movieID uint64) ([]domain.Profession, error) {
	query := `SELECT id, movie_id, person_id, movie_role FROM professions WHERE movie_id = $1 AND movie_role = $2;`

	rows, err := p.Conn.QueryContext(ctx, query, movieID, domain.WriterRole)
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

	result := make([]domain.Profession, 0)
	for rows.Next() {
		tmpProfession := domain.Profession{}
		err = rows.Scan(
			&tmpProfession.ID,
			&tmpProfession.MovieID,
			&tmpProfession.PersonID,
			&tmpProfession.Role)

		if err != nil {
			logrus.Errorf("Repo error: %v", err)
			return nil, err
		}

		result = append(result, tmpProfession)
	}

	return result, err
}

func (p pgProfessionalRepo) GetActorsByMovie(ctx context.Context, movieID uint64) ([]domain.Profession, error) {
	query := `SELECT id, movie_id, person_id, movie_role FROM professions WHERE movie_id = $1 AND movie_role = $2;`

	rows, err := p.Conn.QueryContext(ctx, query, movieID, domain.ActorsRole)
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

	result := make([]domain.Profession, 0)
	for rows.Next() {
		tmpProfession := domain.Profession{}
		err = rows.Scan(
			&tmpProfession.ID,
			&tmpProfession.MovieID,
			&tmpProfession.PersonID,
			&tmpProfession.Role)

		if err != nil {
			logrus.Errorf("Repo error: %v", err)
			return nil, err
		}

		result = append(result, tmpProfession)
	}

	return result, err
}

func (p pgProfessionalRepo) Add(ctx context.Context, movieID, personID, role uint64) error {
	query := `INSERT into professions(movie_id, person_id, movie_role) VALUES ($1, $2, $3);`

	_, err := p.Conn.ExecContext(ctx, query, movieID, personID, role)

	return err
}

func (p pgProfessionalRepo) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM professions WHERE id = $1;`

	_, err := p.Conn.ExecContext(ctx, query, id)

	return err
}

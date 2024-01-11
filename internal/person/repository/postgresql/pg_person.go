package postgresql

import (
	"Kinopoisk-Parser/internal/domain"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type pgPersonRepo struct {
	Conn *sql.DB
}

func New(conn *sql.DB) domain.PersonRepository {
	return &pgPersonRepo{Conn: conn}
}

func (p pgPersonRepo) GetByID(ctx context.Context, id uint64) (domain.Person, error) {
	query := `SELECT id, full_name, age, height FROM person WHERE id = $1;`

	rows, err := p.Conn.QueryContext(ctx, query, id)

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("Repo closing error: %v", err)
		}
	}()

	person := domain.Person{}
	if rows.Next() {
		err = rows.Scan(
			&person.ID,
			&person.FullName,
			&person.Age,
			&person.Height)
	} else {
		err = domain.PersonNotFound
	}

	if err != nil && err != domain.PersonNotFound {
		logrus.Errorf("Repo error: %v", err)
		return domain.Person{}, err
	}

	return person, err

}

func (p pgPersonRepo) GetByFullName(ctx context.Context, title string) (domain.Person, error) {
	query := `SELECT id, full_name, age, height FROM person WHERE full_name = $1;`

	rows, err := p.Conn.QueryContext(ctx, query, title)

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Errorf("Repo closing error: %v", err)
		}
	}()

	person := domain.Person{}
	if rows.Next() {
		err = rows.Scan(
			&person.ID,
			&person.FullName,
			&person.Age,
			&person.Height)
	} else {
		err = domain.PersonNotFound
	}

	if err != nil && err != domain.PersonNotFound {
		logrus.Errorf("Repo error: %v", err)
		return domain.Person{}, err
	}

	return person, err

}

func (p pgPersonRepo) GetPersons(ctx context.Context, limit, offset uint64) ([]domain.Person, error) {
	query := `SELECT id, full_name, age, height FROM person LIMIT $1 OFFSET $2;`

	rows, err := p.Conn.QueryContext(ctx, query, limit, offset)

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

	result := make([]domain.Person, 0)
	for rows.Next() {
		tmpPerson := domain.Person{}
		err = rows.Scan(
			&tmpPerson.ID,
			&tmpPerson.FullName,
			&tmpPerson.Age,
			&tmpPerson.Height)

		if err != nil {
			logrus.Errorf("Repo error: %v", err)
			return nil, err
		}

		result = append(result, tmpPerson)
	}

	return result, err
}

func (p pgPersonRepo) Add(ctx context.Context, m *domain.Person) error {
	query := `INSERT into person(id, full_name, age, height) VALUES ($1, $2, $3, $4);`

	_, err := p.Conn.ExecContext(ctx, query, m.ID, m.FullName, m.Age, m.Height)

	return err
}

func (p pgPersonRepo) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM person WHERE id = $1;`

	_, err := p.Conn.ExecContext(ctx, query, id)

	return err
}

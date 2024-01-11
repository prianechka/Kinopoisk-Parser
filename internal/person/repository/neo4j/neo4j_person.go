package neo4jPersonRepo

import (
	"Kinopoisk-Parser/internal/domain"
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sirupsen/logrus"
)

type Neo4jPersonRepo struct {
	Driver neo4j.DriverWithContext
}

func New(driver neo4j.DriverWithContext) domain.PersonRepository {
	return &Neo4jPersonRepo{Driver: driver}
}

func (n Neo4jPersonRepo) GetByID(ctx context.Context, id uint64) (domain.Person, error) {
	result, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (p:Person {ID: $id}) return p",
		map[string]any{
			"ID": id,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		logrus.Error(err)
		return domain.Person{}, err
	}

	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "p")
	if err != nil {
		return domain.Person{}, domain.PersonNotFound
	}

	person := domain.Person{ID: id}
	person.FullName, _ = neo4j.GetProperty[string](itemNode, "Full_name")

	age, _ := neo4j.GetProperty[int64](itemNode, "Age")
	person.Age = uint64(age)

	height, _ := neo4j.GetProperty[int64](itemNode, "Height")
	person.Age = uint64(height)

	return person, nil
}

func (n Neo4jPersonRepo) GetByFullName(ctx context.Context, title string) (domain.Person, error) {
	result, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (p:Person {Full_name: $full_name}) return p",
		map[string]any{
			"full_name": title,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		logrus.Error(err)
		return domain.Person{}, err
	}

	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "p")
	if err != nil {
		return domain.Person{}, domain.PersonNotFound
	}

	person := domain.Person{FullName: title}

	id, _ := neo4j.GetProperty[int64](itemNode, "ID")
	person.ID = uint64(id)

	person.FullName, _ = neo4j.GetProperty[string](itemNode, "Full_name")

	age, _ := neo4j.GetProperty[int64](itemNode, "Age")
	person.Age = uint64(age)

	height, _ := neo4j.GetProperty[int64](itemNode, "Height")
	person.Height = uint64(height)

	return person, nil
}

func (n Neo4jPersonRepo) GetPersons(ctx context.Context, limit, offset uint64) ([]domain.Person, error) {
	result, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (p:Person) return p",
		map[string]any{}, neo4j.EagerResultTransformer)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	resultPersons := make([]domain.Person, 0)

	for _, record := range result.Records[offset : offset+limit] {
		itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")

		if err != nil {
			return resultPersons, err
		}

		person := domain.Person{}
		id, _ := neo4j.GetProperty[int64](itemNode, "ID")
		person.ID = uint64(id)

		age, _ := neo4j.GetProperty[int64](itemNode, "Age")
		person.Age = uint64(age)

		height, _ := neo4j.GetProperty[int64](itemNode, "Height")
		person.Height = uint64(height)

		resultPersons = append(resultPersons, person)
	}

	return resultPersons, nil
}

func (n Neo4jPersonRepo) Add(ctx context.Context, m *domain.Person) error {
	_, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"CREATE (p:Person {ID: $id, Full_name: $full_name, Age: $age, Height: $height}) RETURN m",
		map[string]any{
			"ID":        m.ID,
			"full_name": m.FullName,
			"age":       m.Age,
			"height":    m.Height,
		}, neo4j.EagerResultTransformer)

	return err
}

func (n Neo4jPersonRepo) Delete(ctx context.Context, id uint64) error {
	_, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"DELETE (p:Person {ID: $id})",
		map[string]any{
			"id": id,
		}, neo4j.EagerResultTransformer)

	return err
}

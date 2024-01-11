package neo4jProfessionRepo

import (
	"Kinopoisk-Parser/internal/domain"
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sirupsen/logrus"
)

type Neo4jProfessionRepo struct {
	Driver neo4j.DriverWithContext
}

func New(driver neo4j.DriverWithContext) domain.ProfessionRepository {
	return &Neo4jProfessionRepo{Driver: driver}
}

func (n Neo4jProfessionRepo) GetDirectorsByMovie(ctx context.Context, movieID uint64) ([]domain.Profession, error) {
	result, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (p:Person)-[:DIRECTED]->(m:Movie {ID: $id) return p, m",
		map[string]any{
			"ID": movieID,
		}, neo4j.EagerResultTransformer)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	resultProfessions := make([]domain.Profession, 0)

	for _, record := range result.Records {
		personNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
		if err != nil {
			return nil, err
		}

		movieNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "m")
		if err != nil {
			return nil, err
		}

		profession := domain.Profession{}

		personID, _ := neo4j.GetProperty[int64](personNode, "ID")
		profession.PersonID = uint64(personID)

		movieID, _ := neo4j.GetProperty[int64](movieNode, "ID")
		profession.MovieID = uint64(movieID)

		profession.Role = domain.DirectorRole

		resultProfessions = append(resultProfessions, profession)
	}

	return resultProfessions, nil
}

func (n Neo4jProfessionRepo) GetProducersByMovie(ctx context.Context, movieID uint64) ([]domain.Profession, error) {
	result, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (p:Person)-[:PRODUCED]->(m:Movie {ID: $id) return p, m",
		map[string]any{
			"ID": movieID,
		}, neo4j.EagerResultTransformer)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	resultProfessions := make([]domain.Profession, 0)

	for _, record := range result.Records {
		personNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
		if err != nil {
			return nil, err
		}

		movieNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "m")
		if err != nil {
			return nil, err
		}

		profession := domain.Profession{}

		personID, _ := neo4j.GetProperty[int64](personNode, "ID")
		profession.PersonID = uint64(personID)

		movieID, _ := neo4j.GetProperty[int64](movieNode, "ID")
		profession.MovieID = uint64(movieID)

		profession.Role = domain.ProducerRole

		resultProfessions = append(resultProfessions, profession)
	}

	return resultProfessions, nil
}

func (n Neo4jProfessionRepo) GetWritersByMovie(ctx context.Context, movieID uint64) ([]domain.Profession, error) {
	result, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (p:Person)-[:WROTE]->(m:Movie {ID: $id) return p, m",
		map[string]any{
			"ID": movieID,
		}, neo4j.EagerResultTransformer)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	resultProfessions := make([]domain.Profession, 0)

	for _, record := range result.Records {
		personNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
		if err != nil {
			return nil, err
		}

		movieNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "m")
		if err != nil {
			return nil, err
		}

		profession := domain.Profession{}

		personID, _ := neo4j.GetProperty[int64](personNode, "ID")
		profession.PersonID = uint64(personID)

		movieID, _ := neo4j.GetProperty[int64](movieNode, "ID")
		profession.MovieID = uint64(movieID)

		profession.Role = domain.WriterRole

		resultProfessions = append(resultProfessions, profession)
	}

	return resultProfessions, nil
}

func (n Neo4jProfessionRepo) GetActorsByMovie(ctx context.Context, movieID uint64) ([]domain.Profession, error) {
	result, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (p:Person)-[:ACTED_IN]->(m:Movie {ID: $id) return p, m",
		map[string]any{
			"ID": movieID,
		}, neo4j.EagerResultTransformer)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	resultProfessions := make([]domain.Profession, 0)

	for _, record := range result.Records {
		personNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
		if err != nil {
			return nil, err
		}

		movieNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "m")
		if err != nil {
			return nil, err
		}

		profession := domain.Profession{}

		personID, _ := neo4j.GetProperty[int64](personNode, "ID")
		profession.PersonID = uint64(personID)

		movieID, _ := neo4j.GetProperty[int64](movieNode, "ID")
		profession.MovieID = uint64(movieID)

		profession.Role = domain.ActorsRole

		resultProfessions = append(resultProfessions, profession)
	}

	return resultProfessions, nil
}

func (n Neo4jProfessionRepo) Add(ctx context.Context, movieID, personID, role uint64) error {
	var strRole string
	switch role {
	case domain.DirectorRole:
		strRole = "DIRECTED"
	case domain.ProducerRole:
		strRole = "PRODUCED"
	case domain.WriterRole:
		strRole = "WROTE"
	default:
		strRole = "ACTED_IN"

	}

	_, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (p:Person {ID:$personID}), (m:Movie {ID:$movieID}) CREATE (p)-[:$role]->(m)",
		map[string]any{
			"personID": personID,
			"movieID":  movieID,
			"role":     strRole,
		}, neo4j.EagerResultTransformer)

	return err
}

func (n Neo4jProfessionRepo) Delete(ctx context.Context, id uint64) error {
	return nil
}

func (n Neo4jProfessionRepo) GetIDByParams(ctx context.Context, movieID, personID, role uint64) (uint64, error) {
	return 0, nil
}

func (n Neo4jProfessionRepo) GetByID(ctx context.Context, id uint64) (domain.Profession, error) {
	return domain.Profession{}, nil
}

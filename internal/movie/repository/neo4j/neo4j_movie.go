package neo4jMovieRepo

import (
	"Kinopoisk-Parser/internal/domain"
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sirupsen/logrus"
)

type Neo4jMovieRepo struct {
	Driver neo4j.DriverWithContext
}

func New(driver neo4j.DriverWithContext) domain.MovieRepository {
	return &Neo4jMovieRepo{Driver: driver}
}

func (n Neo4jMovieRepo) GetByID(ctx context.Context, id uint64) (domain.MovieBaseInfo, error) {
	result, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (m:Movie {ID: $id}) return m",
		map[string]any{
			"id": id,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		logrus.Error(err)
		return domain.MovieBaseInfo{}, err
	}

	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "m")
	if err != nil {
		return domain.MovieBaseInfo{}, domain.MovieNotFound
	}

	movie := domain.MovieBaseInfo{ID: id}

	movie.Title, _ = neo4j.GetProperty[string](itemNode, "Title")
	movie.Tagline, _ = neo4j.GetProperty[string](itemNode, "Tagline")

	year, _ := neo4j.GetProperty[int64](itemNode, "Year")
	movie.Year = uint64(year)

	budget, _ := neo4j.GetProperty[int64](itemNode, "Budget")
	movie.Budget = uint64(budget)

	gross, _ := neo4j.GetProperty[int64](itemNode, "Gross")
	movie.Gross = uint64(gross)

	duration, _ := neo4j.GetProperty[int64](itemNode, "Duration")
	movie.Duration = uint64(duration)

	movie.Rating, _ = neo4j.GetProperty[float64](itemNode, "Rating")

	return movie, nil
}

func (n Neo4jMovieRepo) GetByTitle(ctx context.Context, title string) (domain.MovieBaseInfo, error) {
	result, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (m:Movie {Title: $title}) return m",
		map[string]any{
			"title": title,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		logrus.Error(err)
		return domain.MovieBaseInfo{}, err
	}

	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "m")
	if err != nil {
		return domain.MovieBaseInfo{}, domain.MovieNotFound
	}

	movie := domain.MovieBaseInfo{Title: title}

	id, _ := neo4j.GetProperty[int64](itemNode, "ID")
	movie.ID = uint64(id)

	movie.Tagline, _ = neo4j.GetProperty[string](itemNode, "Tagline")

	year, _ := neo4j.GetProperty[int64](itemNode, "Year")
	movie.Year = uint64(year)

	budget, _ := neo4j.GetProperty[int64](itemNode, "Budget")
	movie.Budget = uint64(budget)

	gross, _ := neo4j.GetProperty[int64](itemNode, "Gross")
	movie.Gross = uint64(gross)

	duration, _ := neo4j.GetProperty[int64](itemNode, "Duration")
	movie.Duration = uint64(duration)

	movie.Rating, _ = neo4j.GetProperty[float64](itemNode, "Rating")

	return movie, nil
}

func (n Neo4jMovieRepo) GetMovies(ctx context.Context, limit, offset uint64) ([]domain.MovieBaseInfo, error) {
	result, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"MATCH (m:Movie) return m",
		map[string]any{}, neo4j.EagerResultTransformer)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	resultMovies := make([]domain.MovieBaseInfo, 0)

	for _, record := range result.Records[offset : offset+limit] {
		itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](record, "m")

		if err != nil {
			return resultMovies, err
		}

		movie := domain.MovieBaseInfo{}

		id, _ := neo4j.GetProperty[int64](itemNode, "ID")
		movie.ID = uint64(id)

		movie.Title, _ = neo4j.GetProperty[string](itemNode, "Title")
		movie.Tagline, _ = neo4j.GetProperty[string](itemNode, "Tagline")

		year, _ := neo4j.GetProperty[int64](itemNode, "Year")
		movie.Year = uint64(year)

		budget, _ := neo4j.GetProperty[int64](itemNode, "Budget")
		movie.Budget = uint64(budget)

		gross, _ := neo4j.GetProperty[int64](itemNode, "Gross")
		movie.Gross = uint64(gross)

		duration, _ := neo4j.GetProperty[int64](itemNode, "Duration")
		movie.Duration = uint64(duration)

		movie.Rating, _ = neo4j.GetProperty[float64](itemNode, "Rating")

		resultMovies = append(resultMovies, movie)
	}

	return resultMovies, nil
}

func (n Neo4jMovieRepo) Add(ctx context.Context, m *domain.MovieBaseInfo) error {
	_, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"CREATE (m:Movie {Title: $title, Tagline: $tagline, Year: $year, Budget: $budget, Gross: $gross, Rating: $rating, Duration: $duration}) RETURN m",
		map[string]any{
			"title":    m.Title,
			"tagline":  m.Tagline,
			"year":     m.Year,
			"budget":   m.Budget,
			"gross":    m.Gross,
			"duration": m.Duration,
			"rating":   m.Rating,
		}, neo4j.EagerResultTransformer)

	return err
}

func (n Neo4jMovieRepo) Delete(ctx context.Context, id uint64) error {
	_, err := neo4j.ExecuteQuery(ctx, n.Driver,
		"DELETE (m:Movie {ID: $id})",
		map[string]any{
			"id": id,
		}, neo4j.EagerResultTransformer)

	return err
}

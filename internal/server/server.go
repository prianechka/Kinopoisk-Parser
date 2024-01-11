package server

import (
	"Kinopoisk-Parser/config"
	"Kinopoisk-Parser/internal/domain"
	delivery "Kinopoisk-Parser/internal/movie/delivery/http"
	neo4jMovieRepo "Kinopoisk-Parser/internal/movie/repository/neo4j"
	postgresqlMovieRepo "Kinopoisk-Parser/internal/movie/repository/postgresql"
	"Kinopoisk-Parser/internal/movie/usecase"
	movieParser "Kinopoisk-Parser/internal/parser"
	neo4jPersonRepo "Kinopoisk-Parser/internal/person/repository/neo4j"
	postgresPersonRepo "Kinopoisk-Parser/internal/person/repository/postgresql"
	neo4jProfessionRepo "Kinopoisk-Parser/internal/profession/repository/neo4j"
	postgresqlProfessionRepo "Kinopoisk-Parser/internal/profession/repository/postgresql"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct {
	config *config.ServerConfig
}

func CreateServer(config *config.ServerConfig) *Server {
	return &Server{config: config}
}

func (s *Server) Start() error {
	r := mux.NewRouter()

	var (
		personRepo     domain.PersonRepository
		movieRepo      domain.MovieRepository
		professionRepo domain.ProfessionRepository
	)

	if s.config.DbParams.Scheme == "neo4j" {
		db := domain.InitNeo4jConnectionByParams(s.config.DbParams)
		personRepo = neo4jPersonRepo.New(db)
		movieRepo = neo4jMovieRepo.New(db)
		professionRepo = neo4jProfessionRepo.New(db)
	} else {
		db := domain.InitPgConnectionByParams(s.config.DbParams)
		personRepo = postgresPersonRepo.New(db)
		movieRepo = postgresqlMovieRepo.New(db)
		professionRepo = postgresqlProfessionRepo.New(db)
	}

	movieUsecase := usecase.NewMovieUsecase(movieRepo, personRepo, professionRepo, 5*time.Second)
	movieHandler := delivery.NewMovieHandler(movieUsecase)

	parser := movieParser.NewParser(s.config.MaxMovies, s.config.TimeForSleep, s.config.MovieURL, s.config.PersonURL, s.config.Token, movieUsecase)

	go parser.Parse()

	r.HandleFunc("/add", movieHandler.Add).Methods("POST")
	r.HandleFunc("/movies/{movies-title}", movieHandler.GetMovie).Methods("GET")
	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")

	return http.ListenAndServe(s.config.StartPort, r)
}

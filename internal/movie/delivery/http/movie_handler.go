package http

import (
	"Kinopoisk-Parser/internal/domain"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type MovieHandler struct {
	MUsecase domain.MovieUsecase
}

func NewMovieHandler(usecase domain.MovieUsecase) MovieHandler {
	return MovieHandler{MUsecase: usecase}
}

func (h *MovieHandler) Add(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error("Bad request")
		return
	}

	var movie domain.Movie

	err = json.Unmarshal(body, &movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error("Bad request while unmarshal")
		return
	}

	err = h.MUsecase.Add(context.Background(), &movie)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("Error while add: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	movieTitle, _ := mux.Vars(r)["title"]
	movieInfo, err := h.MUsecase.GetByTitle(context.Background(), movieTitle)
	switch err {
	case domain.MovieNotFound:
		w.WriteHeader(http.StatusNotFound)
		logrus.Errorf("movie not found: %v", err)
		return
	case nil:

	default:
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("movie get error: %v", err)
		return
	}

	movieRaw, err := json.Marshal(movieInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("Error while add: %v", err)
		return
	}

	w.Write(movieRaw)
	w.WriteHeader(http.StatusOK)
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movieInfo, err := h.MUsecase.GetMovies(context.Background(), 1, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	movieRaw, err := json.Marshal(movieInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("Error while add: %v", err)
		return
	}

	w.Write(movieRaw)
	w.WriteHeader(http.StatusOK)
}

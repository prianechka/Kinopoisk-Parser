package parser

import (
	"Kinopoisk-Parser/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type Parser struct {
	MaxMovies       uint64
	TimeForSleep    uint64
	MovieURL        string
	PersonURL       string
	Token           string
	Usecase         domain.MovieUsecase
	ListOfActorsIDs []int
}

func NewParser(maxMovies, TimeForSleep uint64, movieURL, personURL, token string, usecase domain.MovieUsecase) Parser {
	return Parser{
		MaxMovies:       maxMovies,
		MovieURL:        movieURL,
		PersonURL:       personURL,
		Token:           token,
		TimeForSleep:    TimeForSleep,
		Usecase:         usecase,
		ListOfActorsIDs: []int{},
	}
}

func (p *Parser) Parse() {
	index := 1000
	for {
		for len(p.ListOfActorsIDs) > 0 {
			err := p.parsePerson(p.ListOfActorsIDs[0])
			if err != nil {
				logrus.Error(err)
			}
			p.ListOfActorsIDs = p.ListOfActorsIDs[1:]
		}
		err := p.parseMovie(index)
		if err != nil {
			logrus.Error(err)
		}
		index += 1
	}
}

func (p *Parser) sendRequest(req *http.Request) (*http.Response, error) {
	time.Sleep(time.Second * time.Duration(p.TimeForSleep))
	client := &http.Client{}
	return client.Do(req)
}

func (p *Parser) parsePerson(index int) error {
	logrus.Infof("Parse person with index = %d", index)
	URL := fmt.Sprintf("%s/%d", p.PersonURL, index)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		logrus.Error(err)
		return err
	}

	req.Header.Add("X-API-KEY", p.Token)
	resp, err := p.sendRequest(req)

	if err != nil {
		logrus.Error(err)
		return err
	}

	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		logrus.Errorf("Request status code = %d", statusCode)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Parser error person read all: %v", err)
		return err
	}

	var person domain.PersonDTO

	err = json.Unmarshal(body, &person)
	if err != nil {
		logrus.Errorf("Parser error person unmarshal: %v", err)
		return err
	}

	for _, hisMovie := range person.Movies {
		err := p.parseMovie(hisMovie.Id)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}

func (p *Parser) parseMovie(index int) error {
	logrus.Infof("Parse movie with index = %d", index)

	URL := fmt.Sprintf("%s/%d", p.MovieURL, index)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		logrus.Error(err)
		return err
	}

	req.Header.Add("X-API-KEY", p.Token)
	resp, err := p.sendRequest(req)

	if err != nil {
		logrus.Error(err)
		return err
	}

	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		logrus.Errorf("Request status code = %d", statusCode)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Parser error movie read all: %v", err)
		return err
	}

	var movie domain.MovieDTO

	err = json.Unmarshal(body, &movie)
	if err != nil {
		logrus.Errorf("Parser error movie unmarshal: %v", err)
		return err
	}

	result := domain.Movie{
		BaseInfo: domain.MovieBaseInfo{
			ID:       uint64(movie.Id),
			Title:    movie.Name,
			Year:     uint64(movie.Year),
			Tagline:  movie.Slogan,
			Duration: uint64(movie.MovieLength),
			Rating:   movie.Rating.Imdb,
			Budget:   uint64(movie.Budget.Value),
			Gross:    uint64(movie.Fees.World.Value),
		},
		Producers: []domain.Person{},
		Directors: []domain.Person{},
		Actors:    []domain.Person{},
		Writers:   []domain.Person{},
	}

	for _, person := range movie.Persons {
		p.ListOfActorsIDs = append(p.ListOfActorsIDs, person.Id)
		tmpPerson := domain.Person{
			ID:       uint64(person.Id),
			FullName: person.Name,
		}

		switch person.EnProfession {
		case "producer":
			result.Producers = append(result.Producers, tmpPerson)
		case "director":
			result.Directors = append(result.Directors, tmpPerson)
		case "writer":
			result.Writers = append(result.Writers, tmpPerson)
		case "actor":
			result.Actors = append(result.Actors, tmpPerson)
		}
	}

	return p.Usecase.Add(context.Background(), &result)
}

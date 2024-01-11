package usecase

import (
	"Kinopoisk-Parser/internal/domain"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type movieUsecase struct {
	movieRepo        domain.MovieRepository
	personRepo       domain.PersonRepository
	professionalRepo domain.ProfessionRepository
	contextTimeout   time.Duration
}

func NewMovieUsecase(m domain.MovieRepository, p domain.PersonRepository, pr domain.ProfessionRepository,
	timeout time.Duration) domain.MovieUsecase {

	return &movieUsecase{
		movieRepo:        m,
		personRepo:       p,
		professionalRepo: pr,
		contextTimeout:   timeout,
	}
}

func (u *movieUsecase) GetByID(ctx context.Context, id uint64) (result domain.Movie, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	movieBaseInfo, err := u.movieRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return domain.Movie{}, err
	}

	result.BaseInfo = movieBaseInfo
	err = u.getAdditionalInfo(ctx, &result)

	return result, err
}

func (u *movieUsecase) getAdditionalInfo(ctx context.Context, movie *domain.Movie) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.getAllProfessionals(movie)
}

func (u *movieUsecase) getAllProfessionals(movie *domain.Movie) error {
	err := u.getProducers(movie)
	if err != nil {
		return err
	}

	err = u.getDirectors(movie)
	if err != nil {
		return err
	}

	err = u.getWriters(movie)
	if err != nil {
		return err
	}

	err = u.getActors(movie)
	return err
}

func (u *movieUsecase) getProducers(movie *domain.Movie) error {
	producers, err := u.professionalRepo.GetProducersByMovie(context.Background(), movie.BaseInfo.ID)
	if err != nil {
		logrus.Errorf("Usecase(getProducers) err: %v", err)
		return err
	}

	producersInfo, err := u.addPersons(producers)
	if err == nil {
		movie.Producers = producersInfo
	}

	return err
}

func (u *movieUsecase) getDirectors(movie *domain.Movie) error {
	directors, err := u.professionalRepo.GetDirectorsByMovie(context.Background(), movie.BaseInfo.ID)
	if err != nil {
		logrus.Errorf("Usecase(getDirectors) err: %v", err)
		return err
	}

	directorsInfo, err := u.addPersons(directors)
	if err == nil {
		movie.Directors = directorsInfo
	}

	return err
}

func (u *movieUsecase) getWriters(movie *domain.Movie) error {
	writers, err := u.professionalRepo.GetWritersByMovie(context.Background(), movie.BaseInfo.ID)
	if err != nil {
		logrus.Errorf("Usecase(getWriters) err: %v", err)
		return err
	}

	writersInfo, err := u.addPersons(writers)
	if err == nil {
		movie.Writers = writersInfo
	}

	return err
}

func (u *movieUsecase) getActors(movie *domain.Movie) error {
	actors, err := u.professionalRepo.GetActorsByMovie(context.Background(), movie.BaseInfo.ID)
	if err != nil {
		logrus.Errorf("Usecase(getActors) err: %v", err)
		return err
	}

	actorsInfo, err := u.addPersons(actors)
	if err == nil {
		movie.Actors = actorsInfo
	}

	return err
}

func (u *movieUsecase) addPersons(professions []domain.Profession) ([]domain.Person, error) {
	result := make([]domain.Person, 0)
	for _, person := range professions {
		personInfo, err := u.personRepo.GetByID(context.Background(), person.PersonID)
		if err != nil {
			logrus.Errorf("Usecase(getActors) err: %v", err)
			return nil, err
		}

		result = append(result, personInfo)
	}

	return result, nil
}

func (u *movieUsecase) GetByTitle(ctx context.Context, title string) (result domain.Movie, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	movieBaseInfo, err := u.movieRepo.GetByTitle(ctx, title)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return domain.Movie{}, err
	}

	result.BaseInfo = movieBaseInfo
	err = u.getAdditionalInfo(ctx, &result)

	return result, err
}

func (u *movieUsecase) GetMovies(ctx context.Context, limit, offset uint64) ([]domain.Movie, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	movies, err := u.movieRepo.GetMovies(ctx, limit, offset)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return nil, err
	}

	result := make([]domain.Movie, 0)

	for i := range movies {
		currentMovie := domain.Movie{
			BaseInfo: movies[i],
		}
		err := u.getAdditionalInfo(ctx, &currentMovie)
		if err != nil {
			logrus.Errorf("Usecase: %v", err)
			return nil, err
		}

		result = append(result, currentMovie)
	}

	return result, nil
}

func (u *movieUsecase) Add(ctx context.Context, m *domain.Movie) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.movieRepo.Add(ctx, &m.BaseInfo)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	err = u.addAllPersonsToDB(m)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	err = u.addAllProfessionsToDB(m)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	return nil
}

func (u *movieUsecase) addAllPersonsToDB(m *domain.Movie) error {
	err := u.addPersonsToDB(m.Directors)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	err = u.addPersonsToDB(m.Producers)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	err = u.addPersonsToDB(m.Writers)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	err = u.addPersonsToDB(m.Actors)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	return nil
}

func (u *movieUsecase) addPersonsToDB(persons []domain.Person) error {
	for _, el := range persons {
		_, err := u.personRepo.GetByFullName(context.Background(), el.FullName)
		switch err {
		case domain.PersonNotFound:
			err = u.personRepo.Add(context.Background(), &el)
			if err != nil {
				logrus.Errorf("Usecase: %v", err)
				return fmt.Errorf("usecase: %v", err)
			}
		case nil:
			logrus.Infof("Person with id = %d already in base", el.ID)
		default:
			logrus.Errorf("Usecase: %v", err)
			return fmt.Errorf("usecase: %v", err)
		}
	}

	return nil
}

func (u *movieUsecase) addAllProfessionsToDB(m *domain.Movie) error {
	err := u.addProfessionalToDB(m.Directors, m.BaseInfo.ID, domain.DirectorRole)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	err = u.addProfessionalToDB(m.Producers, m.BaseInfo.ID, domain.ProducerRole)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	err = u.addProfessionalToDB(m.Writers, m.BaseInfo.ID, domain.WriterRole)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	err = u.addProfessionalToDB(m.Actors, m.BaseInfo.ID, domain.ActorsRole)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	return nil
}

func (u *movieUsecase) addProfessionalToDB(persons []domain.Person, movieID, role uint64) error {
	for _, person := range persons {
		err := u.professionalRepo.Add(context.Background(), movieID, person.ID, role)
		if err != nil {
			logrus.Errorf("Usecase: %v", err)
			return fmt.Errorf("usecase: %v", err)
		}
	}

	return nil
}

func (u *movieUsecase) Delete(ctx context.Context, id uint64) error {
	err := u.movieRepo.Delete(ctx, id)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	err = u.professionalRepo.Delete(ctx, id)
	if err != nil {
		logrus.Errorf("Usecase: %v", err)
		return fmt.Errorf("usecase: %v", err)
	}

	return nil
}

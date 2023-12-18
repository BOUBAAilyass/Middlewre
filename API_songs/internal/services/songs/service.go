package songs

import (
	"Projet_Middleware/internal/models"
	repository "Projet_Middleware/internal/repositories/songs"
	

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"database/sql"
	"errors"
	"net/http"
)

func CreateSong(song models.Song) error {
	err := repository.CreateSong(song)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du song : %s", err.Error())
		return err
	}
	return nil
}

func GetAllSongs() ([]models.Song, error) {
	var err error
	// calling repository
	songs, err := repository.GetAllSongs()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}
func GetSongById(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetSongById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}
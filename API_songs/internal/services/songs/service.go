package songs

import (
	"Projet_Middleware/internal/models"
	repository "Projet_Middleware/internal/repositories/songs"
	

	
	"github.com/sirupsen/logrus"
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
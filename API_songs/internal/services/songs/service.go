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


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

func UpdateSong(songID uuid.UUID, updatedSong models.Song) error {
	song, err := repository.GetSongById(songID)
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération du song : %s", err.Error())
		return err
	}

	// Mettre à jour les champs nécessaires du commentaire récupéré avec les données du commentaire mis à jour
	song.Title = updatedSong.Title
	song.Artist = updatedSong.Artist
	song.Album = updatedSong.Album
	song.Year = updatedSong.Year
	song.Path = updatedSong.Path

	err = repository.UpdateSong(song)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du song en base de données : %s", err.Error())
		return err
	}

	return nil
}

func DeleteSong(songID uuid.UUID) error {
	err := repository.DeleteSong(songID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du commentaire : %s", err.Error())
		return err
	}

	return nil
}
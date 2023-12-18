package songs

import (
	"Projet_Middleware/internal/helpers"
	"Projet_Middleware/internal/models"

	
	"github.com/sirupsen/logrus"
)

func CreateSong(song models.Song) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO songs ( id ,title, artist, album, year, path) VALUES (?, ?, ?, ?, ?, ?)",
		song.ID, song.Title, song.Artist, song.Album, song.Year, song.Path)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion du song dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.ID, &data.Title, &data.Artist, &data.Album, &data.Year, &data.Path)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}

	_ = rows.Close()

	return songs, err
}

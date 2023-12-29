package songs

import (
	"Projet_Middleware/internal/helpers"
	"Projet_Middleware/internal/models"

	"github.com/gofrs/uuid"

	"github.com/sirupsen/logrus"
)

func CreateSong(song models.Song) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO songs ( id ,title, artist, file_name, published_date) VALUES (?, ?, ?, ?, ?)",
		song.ID, song.Title, song.Artist, song.FileName, song.PublishedDate)
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
		err = rows.Scan(&data.ID, &data.Title, &data.Artist, &data.FileName, &data.PublishedDate)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}

	_ = rows.Close()

	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {

		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id)
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.ID, &song.Title, &song.Artist, &song.FileName, &song.PublishedDate)

	if err != nil {

		return nil, err // Autres erreurs lors du scan
	}
	return &song, err
}

func UpdateSong(song *models.Song) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)
	_, err = db.Exec("UPDATE songs SET title=?, artist=?, file_name=? WHERE id=?",
		song.Title, song.Artist, song.FileName, song.ID)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du song dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

func DeleteSong(songID uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM songs WHERE id=?", songID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du song dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

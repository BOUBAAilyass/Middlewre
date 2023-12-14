package ratings

import (
	"Projet_Middleware/internal/helpers"
	"Projet_Middleware/internal/models"

	"github.com/sirupsen/logrus"
)

func CreateRating(rating models.Rating) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)
	if _, e := db.Exec("PRAGMA foreign_keys = ON;"); e != nil {
		logrus.Fatalln("Could not enable foreign keys ! Error was : " + e.Error())
	}
	_, err = db.Exec("INSERT INTO ratings ( id ,music_id, user_id, content, rating ) VALUES (?, ?, ?, ?, ?)",
		rating.ID, rating.MusicID, rating.UserID, rating.Content, rating.Rating)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion du rating dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

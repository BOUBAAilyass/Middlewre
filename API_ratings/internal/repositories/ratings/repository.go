package ratings

import (
	"Projet_Middleware/internal/helpers"
	"Projet_Middleware/internal/models"

	"github.com/gofrs/uuid"
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
	_, err = db.Exec("INSERT INTO ratings ( id ,music_id, user_id, content, rating_date, rating ) VALUES (?, ?, ?, ?, ?, ?)",
		rating.ID, rating.MusicID, rating.UserID, rating.Content, rating.Date, rating.Rating)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion du rating dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

func DeleteRating(ratingID uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM ratings WHERE id=?", ratingID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du rating dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

func GetAllRatings() ([]models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM ratings")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	ratings := []models.Rating{}
	for rows.Next() {
		var data models.Rating
		err = rows.Scan(&data.ID, &data.MusicID, &data.UserID, &data.Content, &data.Date, &data.Rating)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, data)
	}

	_ = rows.Close()

	return ratings, err
}

func GetRatingById(id uuid.UUID) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {

		return nil, err
	}
	row := db.QueryRow("SELECT * FROM ratings WHERE id=?", id)
	helpers.CloseDB(db)

	var rating models.Rating
	err = row.Scan(&rating.ID, &rating.MusicID, &rating.UserID, &rating.Content, &rating.Date, &rating.Rating)
	if err != nil {

		return nil, err // Autres erreurs lors du scan
	}
	return &rating, err
}

func UpdateRating(rating *models.Rating) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)
	_, err = db.Exec("UPDATE ratings SET music_id=?, user_id=?, content=?, rating=? WHERE id=?",
		rating.MusicID, rating.UserID, rating.Content, rating.Rating, rating.ID)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du rating dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

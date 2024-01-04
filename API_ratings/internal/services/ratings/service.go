package ratings

import (
	"Projet_Middleware/internal/models"
	repository "Projet_Middleware/internal/repositories/ratings"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func CreateRating(rating models.Rating) error {
	err := repository.CreateRating(rating)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du rating : %s", err.Error())
		return err
	}
	return nil
}

func GetAllRatings() ([]models.Rating, error) {
	var err error
	// calling repository
	ratings, err := repository.GetAllRatings()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return ratings, nil
}

func GetRatingById(id uuid.UUID) (*models.Rating, error) {
	rating, err := repository.GetRatingById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "rating not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving rating : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return rating, err
}

func UpdateRating(ratingID uuid.UUID, updatedRating models.Rating) error {
	rating, err := repository.GetRatingById(ratingID)
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération du rating : %s", err.Error())
		return err
	}

	// Mettre à jour les champs nécessaires du commentaire récupéré avec les données du commentaire mis à jour

	rating.Content = updatedRating.Content

	rating.Rating = updatedRating.Rating

	err = repository.UpdateRating(rating)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du rating en base de données : %s", err.Error())
		return err
	}

	return nil
}

func DeleteRating(ratingID uuid.UUID) error {
	err := repository.DeleteRating(ratingID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du commentaire : %s", err.Error())
		return err
	}

	return nil
}

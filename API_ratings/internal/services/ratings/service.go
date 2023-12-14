package ratings

import (
	"Projet_Middleware/internal/models"
	repository "Projet_Middleware/internal/repositories/ratings"

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

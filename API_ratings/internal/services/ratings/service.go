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

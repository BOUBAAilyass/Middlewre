package users

import (
	"Projet_Middleware/internal/models"
	repository "Projet_Middleware/internal/repositories/users"

	"github.com/sirupsen/logrus"
)

func CreateUser(user models.User) error {
	err := repository.CreateUser(user)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du user : %s", err.Error())
		return err
	}
	return nil
}

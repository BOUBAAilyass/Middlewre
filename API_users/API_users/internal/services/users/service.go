package users

import (
	"Projet_Middleware/internal/models"
	repository "Projet_Middleware/internal/repositories/users"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gofrs/uuid"
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

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving user : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}

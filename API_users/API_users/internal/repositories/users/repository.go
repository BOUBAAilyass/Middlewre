package users

import (
	"Projet_Middleware/internal/helpers"
	"Projet_Middleware/internal/models"

	"github.com/sirupsen/logrus"
)

func CreateUser(user models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO users ( id ,name, email, password) VALUES (?, ?, ?, ?)",
		user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion du user dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

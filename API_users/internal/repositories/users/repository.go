package users

import (
	"Projet_Middleware/internal/helpers"
	"Projet_Middleware/internal/models"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func CreateUser(user models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO users ( id ,name, username, inscription_date) VALUES (?, ?, ?, ?)",
		user.ID, user.Name, user.Username, user.InscriptionDate)
	if err != nil {
		logrus.Errorf("Erreur lors de l'insertion du user dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.ID, &data.Name, &data.Username, &data.InscriptionDate)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}

	_ = rows.Close()

	return users, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {

		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id)
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.InscriptionDate)

	if err != nil {

		return nil, err // Autres erreurs lors du scan
	}
	return &user, err
}

func UpdateUser(user *models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)
	_, err = db.Exec("UPDATE users SET name=?, username=?  WHERE id=?",
		user.Name, user.Username, user.ID)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du user dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

func DeleteUser(userID uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Erreur lors de l'ouverture de la base de données : %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM users WHERE id=?", userID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du user dans la base de données : %s", err.Error())
		return err
	}

	return nil
}

package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/models"
	"Projet_Middleware/internal/services/users"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		logrus.Errorf("Erreur lors de la lecture du corps de la requête : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Générer un nouvel ID UUID
	id, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("Erreur lors de la génération de l'identifiant UUID : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Affecter l'ID généré au commentaire
	newUser.ID = &id
	newUser.InscriptionDate = time.Now().Format("2006-01-02 15:04")
	err = users.CreateUser(newUser)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du User : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(newUser)
	_, _ = w.Write(response)
}

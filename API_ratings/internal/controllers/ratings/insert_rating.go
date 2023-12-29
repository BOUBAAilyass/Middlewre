package ratings

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/models"
	"Projet_Middleware/internal/services/ratings"
)

func InsertRating(w http.ResponseWriter, r *http.Request) {
	var newRating models.Rating
	err := json.NewDecoder(r.Body).Decode(&newRating)
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

	// la date est générée automatiquement par la base de données
	newRating.Date = time.Now().Format("2006-01-02 15:04")

	// Affecter l'ID généré au commentaire
	newRating.ID = &id

	err = ratings.CreateRating(newRating)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du rating : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(newRating)
	_, _ = w.Write(response)
}

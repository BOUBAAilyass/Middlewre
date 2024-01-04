package ratings

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/models"
	"Projet_Middleware/internal/services/ratings"

	"github.com/go-chi/chi/v5"
)

func UpdateRating(w http.ResponseWriter, r *http.Request) {
	ratingID, err := uuid.FromString(chi.URLParam(r, "id"))

	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'ID du rating : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedRating models.Rating
	err = json.NewDecoder(r.Body).Decode(&updatedRating)
	if err != nil {
		logrus.Errorf("Erreur lors de la lecture du corps de la requête : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(updatedRating)
	err = ratings.UpdateRating(ratingID, updatedRating)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du rating : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(updatedRating)
	_, _ = w.Write(response)
}

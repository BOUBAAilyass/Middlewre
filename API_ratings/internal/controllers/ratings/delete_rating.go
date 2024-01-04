package ratings

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/services/ratings"

	"github.com/go-chi/chi/v5"
)

func DeleteRating(w http.ResponseWriter, r *http.Request) {

	ratingID, err := uuid.FromString(chi.URLParam(r, "id"))

	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'ID du rating : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ratings.DeleteRating(ratingID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du rating : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

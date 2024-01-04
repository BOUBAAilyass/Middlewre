package songs

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/services/songs"

	"github.com/go-chi/chi/v5"
)

func DeleteSong(w http.ResponseWriter, r *http.Request) {

	songID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'ID du Song : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = songs.DeleteSong(songID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du Song : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

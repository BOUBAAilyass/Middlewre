package songs

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/models"
	"Projet_Middleware/internal/services/songs"

	"github.com/go-chi/chi/v5"
)

func UpdateSong(w http.ResponseWriter, r *http.Request) {
	songID, err := uuid.FromString(chi.URLParam(r, "id"))

	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'ID du song : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedSong models.Song
	err = json.NewDecoder(r.Body).Decode(&updatedSong)
	if err != nil {
		logrus.Errorf("Erreur lors de la lecture du corps de la requête : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = songs.UpdateSong(songID, updatedSong)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du song : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(updatedSong)
	_, _ = w.Write(response)
}

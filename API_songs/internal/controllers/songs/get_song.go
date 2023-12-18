package songs

import (
	"encoding/json"

	"Projet_Middleware/internal/models"
	"Projet_Middleware/internal/services/songs"
	"net/http"

	"github.com/gofrs/uuid"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func GetSong(w http.ResponseWriter, r *http.Request) {

	songId, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "erreur lors la recuperation de song ID", http.StatusBadRequest)
		return
	}

	song, err := songs.GetSongById(songId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
	return
}

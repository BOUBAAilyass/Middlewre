package songs

import (
	"Projet_Middleware/internal/models"
	"Projet_Middleware/internal/services/songs"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func GetSongs(w http.ResponseWriter, _ *http.Request) {
	// calling service
	songs, err := songs.GetAllSongs()
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(songs)
	_, _ = w.Write(body)
	return
}

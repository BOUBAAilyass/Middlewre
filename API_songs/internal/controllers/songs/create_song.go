package songs

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/models"
	"Projet_Middleware/internal/services/songs"
)

func InsertSong(w http.ResponseWriter, r *http.Request) {
	var newSong models.Song
	err := json.NewDecoder(r.Body).Decode(&newSong)
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
	newSong.ID = &id
	newSong.PublishedDate = time.Now().Format("2006-01-02 15:04")
	err = songs.CreateSong(newSong)
	if err != nil {
		logrus.Errorf("Erreur lors de la création du Song : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(newSong)
	_, _ = w.Write(response)
}

package users

import (
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"Projet_Middleware/internal/services/users"

	"github.com/go-chi/chi/v5"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'ID du User : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = users.DeleteUser(userID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du User : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(http.StatusNoContent)
	w.WriteHeader(http.StatusNoContent)
}

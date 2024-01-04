package ratings

import (
	"Projet_Middleware/internal/models"
	"Projet_Middleware/internal/services/ratings"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func GetRatings(w http.ResponseWriter, r *http.Request) {

	/* 	MusicID, err := uuid.FromString(chi.URLParam(r, "songId"))
	   	if err != nil {
	   		logrus.Errorf("error : %s", err.Error())
	   		w.WriteHeader(http.StatusBadRequest)
	   		return
	   	} */
	// calling service
	ratings, err := ratings.GetAllRatings()
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
	body, _ := json.Marshal(ratings)
	_, _ = w.Write(body)
	return
}

package main

import (
	"Projet_Middleware/internal/controllers/ratings"
	"Projet_Middleware/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func main() {

	router := chi.NewRouter()
	// comments------------------------------------------------------------------------------------------------------------------
	router.Post("/ratings", ratings.InsertRating)
	router.Get("/ratings", ratings.GetRatings)

	logrus.Info("[INFO] Web server started. Now listening on *:8084")
	logrus.Fatalln(http.ListenAndServe(":8084", router))

}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening users database : %s", err.Error())
	}

	// Ratings------------------------------------------------------------------------------------------------------------------

	rating_schemes := []string{
		`CREATE TABLE IF NOT EXISTS ratings (
            id UUID PRIMARY KEY,
            music_id INTEGER NOT NULL,
            user_id INTEGER NOT NULL,
            content TEXT NOT NULL,
			rating REAL CHECK (rating BETWEEN 0 AND 5)
		
            
        );`,
	}

	for _, rating_scheme := range rating_schemes {
		if _, err := db.Exec(rating_scheme); err != nil {
			logrus.Fatalln("Could not generate ratings table ! Error was : " + err.Error())
		}
	}

	helpers.CloseDB(db)

}

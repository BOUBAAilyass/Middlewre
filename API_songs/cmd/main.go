package main

import (
	"Projet_Middleware/internal/controllers/songs"
	"Projet_Middleware/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func main() {

	router := chi.NewRouter()
	// songs------------------------------------------------------------------------------------------------------------------
	router.Post("/songs", songs.InsertSong)
	router.Get("/songs", songs.GetSongs)
	router.Get("/songs/{id}", songs.GetSong)
	router.Put("/songs/{id}", songs.UpdateSong)
	logrus.Info("[INFO] Web server started. Now listening on *:8084")
	logrus.Fatalln(http.ListenAndServe(":8084", router))

}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening songs database : %s", err.Error())
	}

	// songs------------------------------------------------------------------------------------------------------------------

	song_schemes := []string{
		`CREATE TABLE IF NOT EXISTS songs (
			id UUID PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			artist VARCHAR(255) NOT NULL,
			album VARCHAR(255) NOT NULL,
			year INT NOT NULL,
			path VARCHAR(255) NOT NULL
			
		);`,
	}

	for _, song_scheme := range song_schemes {
		if _, err := db.Exec(song_scheme); err != nil {
			logrus.Fatalln("Could not generate songs table ! Error was : " + err.Error())
		}
	}

	helpers.CloseDB(db)

}

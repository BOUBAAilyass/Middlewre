package main

import (
	"Projet_Middleware/internal/controllers/users"
	"Projet_Middleware/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func main() {

	router := chi.NewRouter()
	// users------------------------------------------------------------------------------------------------------------------
	router.Post("/users", users.InsertUser)
	router.Get("/users", users.GetUsers)
	router.Get("/users/{id}", users.GetUser)
	router.Put("/users/{id}", users.UpdateUser)
	router.Delete("/users/{id}", users.DeleteUser)

	logrus.Info("[INFO] Web server started. Now listening on *:8082")
	logrus.Fatalln(http.ListenAndServe(":8082", router))

}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening users database : %s", err.Error())
	}

	// users------------------------------------------------------------------------------------------------------------------

	user_schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			username VARCHAR(255) NOT NULL,
			inscription_date DATE NOT NULL
			
		);`,
	}

	for _, user_scheme := range user_schemes {
		if _, err := db.Exec(user_scheme); err != nil {
			logrus.Fatalln("Could not generate users table ! Error was : " + err.Error())
		}
	}

	helpers.CloseDB(db)

}

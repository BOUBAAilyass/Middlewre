package models

import (
	"github.com/gofrs/uuid"
)

type Rating struct {
	ID      *uuid.UUID `json:"id"`
	MusicID *uuid.UUID `json:"music_id"`
	UserID  *uuid.UUID `json:"user_id"`
	Content string     `json:"content"`
	Date    string     `json:"date"`
	Rating  float64    `json:"rating"`
}

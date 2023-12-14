package models

import (
	"github.com/gofrs/uuid"
)

type Rating struct {
	ID      *uuid.UUID `json:"id"`
	MusicID int        `json:"music_id"`
	UserID  int        `json:"user_id"`
	Content string     `json:"content"`
	Rating  float64    `json:"rating"`
}

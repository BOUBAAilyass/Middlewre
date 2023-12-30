package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	ID            *uuid.UUID `json:"id"`
	Title         string     `json:"title"`
	Artist        string     `json:"artist"`
	FileName      string     `json:"filename"`
	PublishedDate string     `json:"published"`
}

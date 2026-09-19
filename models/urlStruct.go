package models

import "time"

type URL struct {
	ID          string    `bson:"shortenUrl" json:"shortenUrl"`
	OriginalUrl string    `bson:"originalUrl" json:"originalUrl"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
}

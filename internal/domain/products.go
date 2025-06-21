package domain

import "time"

type Product struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	Name         string    `json:"name" bson:"name"`
	CreationTime time.Time `json:"creationTime" bson:"creationTime"`
}

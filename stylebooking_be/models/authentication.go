package models

import (
	"time"
)

type Authentication struct {
	ID         string    `bson:"_id,omitempty"`
	CustomerId string    `bson:"customer_id,omitempty"`
	Password   string    `bson:"password,omitempty"`
	CreatedAt  time.Time `bson:"created_at,omitempty"`
	UpdatedAt  time.Time `bson:"updated_at,omitempty"`
}

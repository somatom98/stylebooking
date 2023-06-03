package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID          primitive.ObjectID            `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string                        `bson:"name,omitempty" json:"name,omitempty"`
	Description string                        `bson:"description,omitempty" json:"description,omitempty"`
	Location    StoreLocation                 `bson:"location,omitempty" json:"location,omitempty"`
	Hours       map[time.Weekday][]StoreHours `bson:"hours,omitempty" json:"hours,omitempty"`
	Services    []Service                     `bson:"services,omitempty" json:"services,omitempty"`
}

type StoreLocation struct {
	Address     string    `bson:"address,omitempty" json:"address,omitempty"`
	City        string    `bson:"city,omitempty" json:"city,omitempty"`
	PostalCode  string    `bson:"postal_code,omitempty" json:"postal_code,omitempty"`
	Country     string    `bson:"country,omitempty" json:"country,omitempty"`
	Coordinates []float64 `bson:"coordinates,omitempty" json:"coordinates,omitempty"`
}

type StoreHours struct {
	Open  time.Time `bson:"open,omitempty" json:"open,omitempty"`
	Close time.Time `bson:"close,omitempty" json:"close,omitempty"`
}

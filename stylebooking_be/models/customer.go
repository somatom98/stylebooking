package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Surname  string             `bson:"surname,omitempty" json:"surname,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Phone    string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}

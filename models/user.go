package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User modelo de usuario
type User struct {
	ID         primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name       string             `bson:"name" json:"name,omitempty"`
	Lastname   string             `bson:"lastname" json:"lastname, omitempty"`
	Nacimiento time.Time          `bson:"nacimiento" json:"nacimiento, omitempty"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password,omitempty"`
	Avatar     string             `bson:"avatar" json:"avatar,omitempty"`
	Banner     string             `bson:"banner" json:"banner,omitempty"`
	Biografia  string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion  string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb   string             `bson:"sitioweb" json:"sitioweb,omitempty"`
}

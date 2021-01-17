package bd

import (
	"context"
	"time"

	"github.com/igorariza/Go-Backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoRegistro es el final con la BD para insertar user
func InsertoRegistro(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("GroupArPost")
	col := db.Collection("User")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}

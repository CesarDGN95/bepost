package models

import (
	// bson es un json encriptado para el id de usuario en mongo db - tipo de dato de mongoDB
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	Email                string             `json:"email"`
	ID                   primitive.ObjectID `bson:"_id" json:"_id",omitemmpty` // Cuando encuentre que este dato contiene omitempty, que lo omita.
	jwt.RegisteredClaims                    // El resto del jscon sera completado con RegisteredClaims que contiene cosas como la fecha de expiracion
}

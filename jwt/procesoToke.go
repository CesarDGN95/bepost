package jwt

import (
	"errors"
	"strings"

	"github.com/CesarDGN95/bepost/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token invalido")
	}

	// OBETENR EL VECTOR 1 SIN EL BEARER
	tk = strings.TrimSpace(splitToken[1])
	//REPROCESAR TK
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		// RUTINA QUE CHEQUEA CON LA BASE DE DATOS
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("token invalido")
	}

	return &claims, false, string(""), err
}

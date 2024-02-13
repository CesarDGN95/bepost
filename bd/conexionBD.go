package bd

import (
	"context"
	"fmt"

	"github.com/CesarDGN95/bepost/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// VARIABLES PUBLICAS
// Conexion abierta hacia MongoDB
var MongoCN *mongo.Client
var DataBaseName string

func ConectarDB(ctx context.Context) error {
	// (Tipo de dato y nombre -> (string) lo covierte a string)
	user := ctx.Value(models.Key("user")).(string)
	passwd := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)

	//ARMAR LA CADENA DE CONEXION
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, passwd, host)
	var clientOptions = options.Client().ApplyURI(connStr)

	//EJECUTAMOS LA CONEXION
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		// si hubo un error enviarlo a cloaudwatch
		fmt.Println(err.Error())
		return err
	}

	// REALIZAR UN PING PARA SABER SI LA CONEXION QUEDO ABIERTA O HUBO ALGUN ERROR
	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexion exitosa")

	// ENVIAR LA CONEXION A LA VARIABLE PUBLICA
	MongoCN = client
	DataBaseName = ctx.Value(models.Key("database")).(string)

	//RETORNAMOS NIL PORQUE NO HUBO ERROR
	return nil

}

func BaseConectada() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	// Si es igual retorna false, sino true
	return err == nil
}

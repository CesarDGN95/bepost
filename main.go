package bepost

import (
	"context"
	"os"
	"strings"

	"github.com/CesarDGN95/bepost/awsgo"
	"github.com/CesarDGN95/bepost/bd"
	"github.com/CesarDGN95/bepost/models"
	"github.com/CesarDGN95/bepost/secretmanager"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InicializoAws()

	// Si no se validan bien los parametros, haz lo siguiente:
	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. deben incluir 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-type": "application/json",
			},
		}
		return res, nil
	}
	// Si si se validan bien, realiza lo siguiente:
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName")) // Traemos la variale de entorno
	if err != nil {

		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de SecretName" + err.Error(),
			Headers: map[string]string{
				"Content-type": "application/json",
			},
		}
		return res, nil
	}
	//VAIRABLES DE EJECUCION EN EL CONTEXT
	path := strings.Replace(request.PathParameters["bepost"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path) // (EN DONDE, TIPO DE DATO, VALOR)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	//CHAQUEO CONEXION A LA DB OCONECTO LA DB
	err = bd.ConectarDB(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error  conectando en la base de datos" + err.Error(),
			Headers: map[string]string{
				"Content-type": "application/json",
			},
		}
		return res, nil
	}

}

func ValidoParametros() bool {
	//Chacate su trae valor o no
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro

	}

	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro {
		return traeParametro

	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro

	}

	return traeParametro

}

package bepost

import (
	"context"
	"os"

	"github.com/CesarDGN95/bepost/awsgo"
	"github.com/CesarDGN95/bepost/secretmanager"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, reques events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
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
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
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

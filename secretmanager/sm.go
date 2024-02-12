package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/CesarDGN95/bepost/awsgo"
	"github.com/CesarDGN95/bepost/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	// Valor que retornaremos
	var datosSecret models.Secret
	fmt.Println("> Pido secreto" + secretName)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{SecretId: aws.String(secretName)})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> Lectura de Secret OK " + secretName)
	return datosSecret, nil
}
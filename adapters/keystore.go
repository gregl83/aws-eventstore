package adapters

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

// keystore adapter provides an interface to aws secrets manager
type keystore struct {
	manager *secretsmanager.SecretsManager
}

// createSecretInput for secrets manager
func (ks keystore) createSecretInput(key string) *secretsmanager.GetSecretValueInput {
	return &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(key),
	}
}

// getSecret from secrets manager
func (ks *keystore) getSecret(input *secretsmanager.GetSecretValueInput) (string, error) {
	var secret string

	result, err := ks.manager.GetSecretValue(input)

	if err != nil {
		return secret, err
	}

	if result.SecretString != nil {
		secret = *result.SecretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		length, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			return secret, err
		}
		secret = string(decodedBinarySecretBytes[:length])
	}

	return secret, nil
}

// ReadKey value from keystore
func (ks *keystore) ReadKey(key string) (string, error) {
	return ks.getSecret(
		ks.createSecretInput(key),
	)
}

// NewKeyStore adapter
func NewKeyStore() *keystore {
	manager := secretsmanager.New(session.New())
	return &keystore{manager: manager}
}

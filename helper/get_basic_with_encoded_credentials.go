package helper

import (
	"encoding/base64"
	"fmt"
)

// GetBasicWithEncodedCredentials returns the basic authentication header with the encoded credentials.
func GetBasicWithEncodedCredentials(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", fmt.Errorf("The params 'username' and 'password' are required and cannot be empty")
	}

	auth := fmt.Sprintf("%s:%s", username, password)
	econdedAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	return fmt.Sprintf("Basic %s", econdedAuth), nil
}

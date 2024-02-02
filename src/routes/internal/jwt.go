package internal

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

// extractDataFromJWT extracts all the data that the JWT contains. In order to do that it uses some
// environment variables. An error is returned if for some reason, the data cannot be extracted
func extractDataFromJWT(tokenString string) (*jwtData, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("error token is missing")
	}

	secretKey := os.Getenv(secretEnvVar)
	if secretKey == "" {
		return nil, fmt.Errorf("error secret key is missing")
	}

	algorithm := os.Getenv(cryptographicAlgorithmEnvVar)
	if algorithm == "" {
		return nil, fmt.Errorf("error algorithm is missing")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Sanity check: the algorithm must be the same as the one in the env var
		if token.Method.Alg() != algorithm {
			return nil, fmt.Errorf("error unexpected signing method: %s", token.Method.Alg())
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing JWT: %v", err)
	}

	// Access the claims if the signature is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		telegramId := ""
		if claims["telegram_id"] != nil {
			telegramId = claims["telegram_id"].(string)
		}
		return &jwtData{
			userID:     claims["user_id"].(string),
			email:      claims["email"].(string),
			telegramID: telegramId,
		}, nil
	}

	return nil, fmt.Errorf("error invalid JWT")
}

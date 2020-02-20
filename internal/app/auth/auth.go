package auth

import (
	"errors"
	"time"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/store"

	jwt "github.com/dgrijalva/jwt-go"
)

// UserAuthenticate ...
func UserAuthenticate(s *store.Store, token string) bool {
	result, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(s.Config.KeyString), nil
	})

	if err != nil {
		return false
	}

	if result.Valid {
		return true
	}

	return false
}

// UserAuthorize ...
func UserAuthorize(s *store.Store, user string, password string) (string, error) {
	var errIncorrectEmailOrPassword = errors.New("Incorrect email or password")
	if user == s.Config.Admin && password == s.Config.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["authorized"] = true
		claims["user"] = "admin"
		claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

		tokenstring, err := token.SignedString([]byte(s.Config.KeyString))
		if err != nil {
			return "", err
		}

		return tokenstring, nil
	}

	return "", errIncorrectEmailOrPassword
}

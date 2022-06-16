package pkg

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/virhanali/koika-app/domain/schemas"
)

func SignToken(config *schemas.JWtMetaRequest) (string, error) {
	expiredAt := time.Now().Add(time.Duration(time.Minute) * config.Options.ExpiredAt).Unix()
	claims := jwt.MapClaims{}
	claims["jwt"] = config.Data
	claims["exp"] = (24 * 60) * expiredAt
	claims["audience"] = config.Options.Audience
	claims["authority"] = true

	to := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	accessToken, err := to.SignedString([]byte(config.SecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}
	return accessToken, nil
}

func VerifyToken(accessToken, SecretPublicKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretPublicKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

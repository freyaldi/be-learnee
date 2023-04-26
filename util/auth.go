package util

import (
	"time"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/config"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	GenerateToken(userId int, isAdmin bool) (string, error)
	ComparePassword(hashedPwd string, inputPwd string) bool
	HashPassword(pwd string) (string, error)
}

type authImpl struct{}

type AuthConfig struct{}

func NewAuth(cfg *AuthConfig) Auth {
	return &authImpl{}
}

type IdTokenClaims struct {
	jwt.RegisteredClaims
	UserId  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}

func (a *authImpl) GenerateToken(userId int, isAdmin bool) (string, error) {
	expireTime := time.Now().Add(1 * time.Hour)

	claims := &IdTokenClaims{}
	claims.UserId = userId
	claims.IsAdmin = isAdmin
	claims.IssuedAt = &jwt.NumericDate{Time: time.Now()}
	claims.ExpiresAt = jwt.NewNumericDate(expireTime)
	claims.Issuer = config.Issuer

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *authImpl) ComparePassword(hashedPwd string, inputPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
	return err == nil
}

func (a *authImpl) HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

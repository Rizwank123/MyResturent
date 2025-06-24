package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rizwank123/myResturent/internal/pkg/config"
)

const issuer = "myResturent"

// authClaims contains the claims of the jwt token
type authClaims struct {
	TokenMetaData
	jwt.RegisteredClaims
}

// jwtManager represents the jwt security manager
type jwtSecurityManager struct {
	cfg config.ResturantConfig
}

// NewJWTSecurityManager returns a new jwt security manager
func NewJWTSecurityManager(cfg config.ResturantConfig) Manager {
	return &jwtSecurityManager{
		cfg: cfg,
	}
}

// GenerateToken implements Manager.
func (s jwtSecurityManager) GenerateToken(tokenMetaData TokenMetaData) (token string, err error) {
	claims := &authClaims{
		TokenMetaData: tokenMetaData,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(s.cfg.AuthExpiryPeriod))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(s.cfg.AuthSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

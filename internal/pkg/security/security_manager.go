package security

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// TokenMetaData represents the metadata of the jwt token
type TokenMetaData struct {
	UserID      string `json:"user_id"`
	Role        string `json:"role"`
	ResturentID string `json:"resturent_id"`
}

// Manager represents the security manager
type Manager interface {
	GenerateToken(tokenMetaData TokenMetaData) (token string, err error)
}

// GetClaimsForContext returns the claims for the context
func GetClaimsForContext(ctx echo.Context) jwt.MapClaims {
	token := ctx.Get("user")
	if token != nil {
		u := token.(*jwt.Token)
		claims := u.Claims.(jwt.MapClaims)
		return claims
	}
	return nil

}

package jwt

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT ...
type JWT struct {
	tokenTTL time.Duration
	secret   []byte
}

// CreateToken to user.
func (j *JWT) CreateToken(ctx context.Context, id string) (token string, err error) {
	accessClaims := &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        id,
			ExpiresAt: time.Now().Unix() + int64(j.tokenTTL.Seconds()),
		},
	}

	token, err = jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), accessClaims).SignedString(j.secret)
	return
}

// Parse token
func (j *JWT) Parse(ctx context.Context, token string) (isValid bool, id string, err error) {
	var (
		jwtToken = &jwt.Token{}
		claims   = &tokenClaims{}
		keyFunc  = func(token *jwt.Token) (interface{}, error) {
			return j.secret, nil
		}
	)

	if jwtToken, err = jwt.ParseWithClaims(token, claims, keyFunc); jwtToken != nil {
		id = claims.Id
		isValid = jwtToken.Valid
	}
	return
}

// New ...
func New(
	tokenTTL time.Duration,
	secret []byte,
) *JWT {
	return &JWT{
		tokenTTL: tokenTTL,
		secret:   secret,
	}
}

package jwt

import (
	"context"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT ...
type JWT struct {
	tokenTTL time.Duration
	secret   []byte
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

// CreateToken to user.
func (j *JWT) CreateToken(ctx context.Context, id int) (string, error) {
	accessClaims := &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.Itoa(id),
			ExpiresAt: time.Now().Unix() + int64(j.tokenTTL.Seconds()),
		},
	}

	return jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), accessClaims).SignedString(j.secret)
}

// Parse token
func (j *JWT) Parse(ctx context.Context, token string) (isValid bool, id int, err error) {
	var (
		jwtToken *jwt.Token
		claims   = &tokenClaims{}
		keyFunc  = func(token *jwt.Token) (interface{}, error) {
			return j.secret, nil
		}
	)

	if jwtToken, err = jwt.ParseWithClaims(token, claims, keyFunc); jwtToken != nil {
		id, _ = strconv.Atoi(claims.Id)
		isValid = jwtToken.Valid
	}
	return
}

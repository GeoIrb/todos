package password

import (
	"context"
	"encoding/base64"
	"hash"
)

// Hash processor
type Hash struct {
	hash func() hash.Hash
	salt []byte
}

// New ...
func NewHash(
	hash func() hash.Hash,
	salt []byte,
) *Hash {
	return &Hash{
		hash: hash,
		salt: salt,
	}
}

// Password returns password hash with salt
func (h *Hash) Password(ctx context.Context, password string) string {
	hash := h.hash()
	passwordBytes := append([]byte(password), h.salt...)
	hash.Write(passwordBytes)
	hashedPasswordBytes := hash.Sum(nil)
	return base64.URLEncoding.EncodeToString(hashedPasswordBytes)
}

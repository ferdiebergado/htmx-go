package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

const (
	byteSize = 32
)

func GenerateSecureRandomBytes() (string, error) {
	bytes := make([]byte, byteSize)

	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

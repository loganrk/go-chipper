package cipher

import (
	"github.com/loganrk/go-cipher/aes"
)

type Cipher interface {
	Encrypt(text string) (string, error)
	Decrypt(cryptoText string) (string, error)
}

func New(cryptoKey string) Chipper {
	return aes.New(cryptoKey)

}

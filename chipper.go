package chipper

import (
	"github.com/loganrk/go-chipper/aes"
)

type Chipper interface {
	Encrypt(text string) (string, error)
	Decrypt(cryptoText string) (string, error)
}

func New(CryptoKey string) Chipper {
	return aes.New(CryptoKey)

}

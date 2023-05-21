package oauth

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"time"
)

const TokenTypeBearer = "Bearer"
const tokenRandomWordLength = 9

type Token struct {
	privateKey      []byte
	publicKey       []byte
	randomKey       []byte
	createdDateTime time.Time
}

func (t Token) Sign() []byte {
	signature := ed25519.Sign(t.privateKey, t.randomKey)

	hexHash := make([]byte, hex.EncodedLen(len(signature)))
	hex.Encode(hexHash, signature)

	base64Hash := make([]byte, base64.StdEncoding.EncodedLen(len(hexHash)))
	base64.StdEncoding.Encode(base64Hash, hexHash)

	return base64Hash
}

func (t Token) Verify(msg []byte) bool {
	base64Hash := make([]byte, base64.StdEncoding.DecodedLen(len(msg)))
	base64.StdEncoding.Decode(base64Hash, msg)

	hexHash := make([]byte, hex.DecodedLen(len(base64Hash)))
	hex.Decode(hexHash, base64Hash)

	return ed25519.Verify(t.publicKey, t.randomKey, hexHash)
}

var instance *Token

func NewToken() (*Token, error) {
	var err error
	if instance == nil {
		instance, err = newToken()

		return instance, err
	}

	// todo: correct the logic for every 24 hours
	if instance.createdDateTime.Nanosecond() < time.Now().Nanosecond() {
		instance, err = newToken()
	}

	return instance, err
}

func newToken() (*Token, error) {
	randomKey := make([]byte, tokenRandomWordLength)
	_, err := rand.Read(randomKey)
	if err != nil {
		return nil, err
	}

	base64RandomKey := make([]byte, base64.StdEncoding.EncodedLen(len(randomKey)))
	base64.StdEncoding.Encode(base64RandomKey, randomKey)

	public, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}

	return &Token{
		publicKey:       public,
		privateKey:      private,
		randomKey:       base64RandomKey,
		createdDateTime: time.Now(),
	}, nil
}

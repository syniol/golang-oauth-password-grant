package oauth

import (
	"bytes"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

type CredentialPassword struct {
	PublicKey      string `json:"publicKey"`
	PrivateKey     string `json:"privateKey"`
	HashedPassword string `json:"hashedPassword"`
}

func NewCredentialPassword(password string) (*CredentialPassword, error) {
	public, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal public key: %s", err.Error())
	}

	var pemEncodedPublicKey bytes.Buffer
	if err := pem.Encode(&pemEncodedPublicKey, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}); err != nil {
		return nil, err
	}

	prvKeyBytes, err := x509.MarshalPKCS8PrivateKey(private)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal private key: %s", err.Error())
	}

	var pemEncodedPrivateKey bytes.Buffer
	if err := pem.Encode(&pemEncodedPrivateKey, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: prvKeyBytes,
	}); err != nil {
		return nil, err
	}

	return &CredentialPassword{
		PublicKey:  encodeHash(pemEncodedPublicKey.Bytes()),
		PrivateKey: encodeHash(pemEncodedPrivateKey.Bytes()),
		HashedPassword: encodeHash(
			ed25519.Sign(private, []byte(password)),
		),
	}, nil
}

func (cred *CredentialPassword) VerifyPassword(inputPassword string) bool {
	return ed25519.Verify(
		decodePublicCert(decodeHash([]byte(cred.PublicKey))),
		[]byte(inputPassword),
		decodeHash([]byte(cred.HashedPassword)),
	)
}

func decodeHash(hash []byte) []byte {
	base64Hash := make([]byte, base64.StdEncoding.DecodedLen(len(hash)))
	base64.StdEncoding.Decode(base64Hash, hash)

	hexHash := make([]byte, hex.DecodedLen(len(base64Hash)))
	hex.Decode(hexHash, base64Hash)

	return hexHash
}

func encodeHash(plaintext []byte) string {
	hexHash := make([]byte, hex.EncodedLen(len(plaintext)))
	hex.Encode(hexHash, plaintext)

	base64Hash := make([]byte, base64.StdEncoding.EncodedLen(len(hexHash)))
	base64.StdEncoding.Encode(base64Hash, hexHash)

	return string(base64Hash)
}

func decodePublicCert(cert []byte) []byte {
	out, _ := pem.Decode(cert)

	parsedPublicKey, _ := x509.ParsePKIXPublicKey(out.Bytes)

	return parsedPublicKey.(ed25519.PublicKey)
}

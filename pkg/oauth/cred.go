package oauth

import (
	"bytes"
	"crypto/ed25519"
	"crypto/x509"
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
		return nil, fmt.Errorf("failed to write data to %s file: %s", public, err.Error())
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
		return nil, fmt.Errorf("failed to write data to %s file: %s", private, err.Error())
	}

	return &CredentialPassword{
		PublicKey:  pemEncodedPublicKey.String(),
		PrivateKey: pemEncodedPrivateKey.String(),
		HashedPassword: hex.EncodeToString(
			ed25519.Sign(private, []byte(password)),
		),
	}, nil
}

func (cred *CredentialPassword) PasswordVerify(inputPassword string) bool {
	return ed25519.Verify(
		decodePublicCert(cred.PublicKey),
		[]byte(inputPassword),
		decodeHash(cred.HashedPassword),
	)
}

func decodeHash(hash string) []byte {
	data := []byte(hash)
	dst := make([]byte, hex.DecodedLen(len(data)))
	hex.Decode(dst, data)

	return dst
}

func decodePublicCert(cert string) []byte {
	out, _ := pem.Decode(ed25519.PublicKey(cert))

	sss, _ := x509.ParsePKIXPublicKey(out.Bytes)

	return sss.(ed25519.PublicKey)
}

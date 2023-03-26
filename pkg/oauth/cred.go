package oauth

import (
	"bytes"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"log"
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

	pubBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		log.Fatalf("Unable to marshal public key: %v", err)
	}

	var buf bytes.Buffer
	if err := pem.Encode(&buf, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}); err != nil {
		log.Fatalf("Failed to write data to %s file: %s", public, err)
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(private)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}

	var keyOut bytes.Buffer
	if err := pem.Encode(&keyOut, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	}); err != nil {
		log.Fatalf("Failed to write data to %s file: %s", private, err)
	}

	return &CredentialPassword{
		PublicKey:  buf.String(),
		PrivateKey: keyOut.String(),
		HashedPassword: hex.EncodeToString(
			ed25519.Sign(private, []byte(password)),
		),
	}, nil
}

func PasswordVerify(inputPassword string, cred CredentialPassword) bool {
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

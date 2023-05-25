package hash

import (
	"encoding/base64"
	"encoding/hex"
)

func Encode(msg []byte) []byte {
	hexHash := make([]byte, hex.EncodedLen(len(msg)))
	hex.Encode(hexHash, msg)

	hash := make([]byte, base64.StdEncoding.EncodedLen(len(hexHash)))
	base64.StdEncoding.Encode(hash, hexHash)

	return hash
}

func Decode(msg []byte) []byte {
	base64Hash := make([]byte, base64.StdEncoding.DecodedLen(len(msg)))
	base64.StdEncoding.Decode(base64Hash, msg)

	hash := make([]byte, hex.DecodedLen(len(base64Hash)))
	hex.Decode(hash, base64Hash)

	return hash
}

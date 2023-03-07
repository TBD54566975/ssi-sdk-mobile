package ssi

import (
	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/TBD54566975/ssi-sdk/did"
	"github.com/goccy/go-json"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/pkg/errors"
)

type generateDIDKeyResult struct {
	DidKey string  `json:"didKey" validate:"required"`
	Jwk    jwk.Key `json:"jwk" validate:"required"`
}

// GenerateDIDKey takes in a key type value that this library supports and constructs a conformant did:key identifier.
// The function returns the associated private key value cast to the generic golang crypto.PrivateKey interface.
// To use the private key, it is recommended to re-cast to the associated type. For example, called with the input
// for a secp256k1 key:
// privKey, didKey, err := GenerateDIDKey(Secp256k1)
// if err != nil { ... }
// // where secp is an import alias to the secp256k1 library we use "github.com/decred/dcrd/dcrec/secp256k1/v4"
// secpPrivKey, ok := privKey.(secp.PrivateKey)
// if !ok { ... }
func GenerateDIDKey(kt string) ([]byte, error) {
	privateKey, didKey, err := did.GenerateDIDKey(stringToKeyType(kt))
	if err != nil {
		return nil, errors.Wrap(err, "generating did key")
	}
	jwk, err := crypto.PrivateKeyToJWK(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "creating jwk")
	}
	resultBytes, err := json.Marshal(generateDIDKeyResult{
		DidKey: string(*didKey),
		Jwk:    jwk,
	})
	if err != nil {
		return nil, errors.Wrap(err, "marshalling result")
	}

	return resultBytes, nil
}

// CreateDIDKey constructs a did:key from a specific key type and its corresponding public key
// This method does not attempt to validate that the provided public key is of the specified key type.
// A safer method is `GenerateDIDKey` which handles key generation based on the provided key type.
func CreateDIDKey(kt string, publicKey []byte) (string, error) {
	didKey, err := did.CreateDIDKey(stringToKeyType(kt), publicKey)
	return string(*didKey), err
}

type DecodedDIDKey struct {
	Data    []byte
	KeyType string
}

// DecodeDIDKey Decode takes a did:key and returns the underlying public key value as bytes, the LD key type, and a possible error
func DecodeDIDKey(d string) (*DecodedDIDKey, error) {
	data, _, keyType, err := did.DIDKey(d).Decode()
	return &DecodedDIDKey{
		Data:    data,
		KeyType: string(keyType),
	}, err
}

// ExpandDIDKey Expand turns the DID key into a compliant DID Document
func ExpandDIDKey(d string) ([]byte, error) {
	didDoc, err := did.DIDKey(d).Expand()
	if err != nil {
		return nil, err
	}

	didDocBytes, err := json.Marshal(didDoc)
	if err != nil {
		return nil, err
	}

	return didDocBytes, nil
}

package pkg

import (
	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/TBD54566975/ssi-sdk/did"
	"github.com/TBD54566975/ssi-sdk/util"
	"github.com/pkg/errors"
)

type DIDKeyWrapper struct {
	PrivateJSONWebKey []byte
	DIDKey            string
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
func GenerateDIDKey(kt string) (*DIDKeyWrapper, error) {
	privateKey, didKey, err := did.GenerateDIDKey(stringToKeyType(kt))
	if err != nil {
		return nil, errors.Wrap(err, "generating did key")
	}
	jwkKey, err := crypto.PrivateKeyToJWK(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "creating jwk")
	}
	data, err := json.Marshal(jwkKey)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling jwk")
	}
	return &DIDKeyWrapper{
		PrivateJSONWebKey: data,
		DIDKey:            string(*didKey),
	}, err
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
func ExpandDIDKey(d string) (*DIDDocumentMobile, error) {
	didDoc, err := did.DIDKey(d).Expand()
	if err != nil {
		return nil, err
	}
	contextArray, err := util.InterfaceToStrings(didDoc.Context)
	if err != nil {
		return nil, err
	}
	return &DIDDocumentMobile{
		Context:              &StringArray{Items: contextArray},
		ID:                   didDoc.ID,
		Controller:           didDoc.Controller,
		AlsoKnownAs:          didDoc.AlsoKnownAs,
		VerificationMethod:   &VerificationMethodArray{Items: didDoc.VerificationMethod},
		Authentication:       &VerificationMethodSetArray{Items: didDoc.Authentication},
		AssertionMethod:      &VerificationMethodSetArray{Items: didDoc.AssertionMethod},
		KeyAgreement:         &VerificationMethodSetArray{Items: didDoc.KeyAgreement},
		CapabilityInvocation: &VerificationMethodSetArray{Items: didDoc.CapabilityInvocation},
		CapabilityDelegation: &VerificationMethodSetArray{Items: didDoc.CapabilityDelegation},
	}, nil
}

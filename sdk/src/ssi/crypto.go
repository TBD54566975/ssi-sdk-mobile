package ssi

import (
	ssi "github.com/TBD54566975/ssi-sdk/crypto"
)

var (
	supportedKeyTypes = StringArray{items: []string{keyTypeToString(ssi.Ed25519), keyTypeToString(ssi.X25519),
		keyTypeToString(ssi.SECP256k1), keyTypeToString(ssi.P256),
		keyTypeToString(ssi.P384), keyTypeToString(ssi.P521), keyTypeToString(ssi.RSA)}}

	signatureAlgs = StringArray{items: []string{signatureToString(ssi.EdDSA), signatureToString(ssi.ES256K),
		signatureToString(ssi.ES256), signatureToString(ssi.ES384), signatureToString(ssi.PS256)}}
)

func IsSupportedKeyType(kt string) bool {
	supported := GetSupportedKeyTypes()
	for _, t := range supported.items {
		if kt == t {
			return true
		}
	}
	return false
}

func GetSupportedKeyTypes() *StringArray {
	return &supportedKeyTypes
}

func IsSupportedSignatureAlg(sa string) bool {
	supported := GetSupportedSignatureAlgs()
	for _, a := range supported.items {
		if sa == a {
			return true
		}
	}
	return false
}

func GetSupportedSignatureAlgs() *StringArray {
	return &signatureAlgs
}

// methods from crypto/keys.go

type CryptoKeyPair struct {
	KeyType string
	PrivKey []byte
	PubKey  []byte
}

func GenerateEd25519Key() (*CryptoKeyPair, error) {
	pubKey, privKey, err := ssi.GenerateEd25519Key()
	privKeyBytes, err := ssi.PrivKeyToBytes(privKey)
	if err != nil {
		return nil, err
	}
	pubKeyBytes, err := ssi.PubKeyToBytes(pubKey)
	if err != nil {
		return nil, err
	}
	return &CryptoKeyPair{
		KeyType: ssi.Ed25519.String(),
		PrivKey: privKeyBytes,
		PubKey:  pubKeyBytes,
	}, err
}

func GenerateX25519Key() (*CryptoKeyPair, error) {
	privKey, pubKey, err := ssi.GenerateX25519Key()
	privKeyBytes, err := ssi.PrivKeyToBytes(privKey)
	if err != nil {
		return nil, err
	}
	pubKeyBytes, err := ssi.PubKeyToBytes(pubKey)
	if err != nil {
		return nil, err
	}
	return &CryptoKeyPair{
		KeyType: ssi.X25519.String(),
		PrivKey: privKeyBytes,
		PubKey:  pubKeyBytes,
	}, err
}

func GenerateSecp256k1Key() (*CryptoKeyPair, error) {
	pubKey, privKey, err := ssi.GenerateSECP256k1Key()
	privKeyBytes, err := ssi.PrivKeyToBytes(privKey)
	if err != nil {
		return nil, err
	}
	pubKeyBytes, err := ssi.PubKeyToBytes(pubKey)
	if err != nil {
		return nil, err
	}
	return &CryptoKeyPair{
		KeyType: ssi.SECP256k1.String(),
		PrivKey: privKeyBytes,
		PubKey:  pubKeyBytes,
	}, err
}

func GenerateP256Key() (*CryptoKeyPair, error) {
	pubKey, privKey, err := ssi.GenerateP256Key()
	privKeyBytes, err := ssi.PrivKeyToBytes(privKey)
	if err != nil {
		return nil, err
	}
	pubKeyBytes, err := ssi.PubKeyToBytes(pubKey)
	if err != nil {
		return nil, err
	}
	return &CryptoKeyPair{
		KeyType: ssi.P256.String(),
		PrivKey: privKeyBytes,
		PubKey:  pubKeyBytes,
	}, err
}

func GenerateP384Key() (*CryptoKeyPair, error) {
	pubKey, privKey, err := ssi.GenerateP384Key()
	privKeyBytes, err := ssi.PrivKeyToBytes(privKey)
	if err != nil {
		return nil, err
	}
	pubKeyBytes, err := ssi.PubKeyToBytes(pubKey)
	if err != nil {
		return nil, err
	}
	return &CryptoKeyPair{
		KeyType: ssi.P384.String(),
		PrivKey: privKeyBytes,
		PubKey:  pubKeyBytes,
	}, err
}

func GenerateP521Key() (*CryptoKeyPair, error) {
	pubKey, privKey, err := ssi.GenerateP521Key()
	privKeyBytes, err := ssi.PrivKeyToBytes(privKey)
	if err != nil {
		return nil, err
	}
	pubKeyBytes, err := ssi.PubKeyToBytes(pubKey)
	if err != nil {
		return nil, err
	}
	return &CryptoKeyPair{
		KeyType: ssi.P521.String(),
		PrivKey: privKeyBytes,
		PubKey:  pubKeyBytes,
	}, err
}

func GenerateRSA2048Key() (*CryptoKeyPair, error) {
	pubKey, privKey, err := ssi.GenerateRSA2048Key()
	privKeyBytes, err := ssi.PrivKeyToBytes(privKey)
	if err != nil {
		return nil, err
	}
	pubKeyBytes, err := ssi.PubKeyToBytes(pubKey)
	if err != nil {
		return nil, err
	}
	return &CryptoKeyPair{
		KeyType: ssi.RSA.String(),
		PrivKey: privKeyBytes,
		PubKey:  pubKeyBytes,
	}, err
}

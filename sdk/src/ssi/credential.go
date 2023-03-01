package ssi

import (
	"github.com/TBD54566975/ssi-sdk/credential"
	"github.com/TBD54566975/ssi-sdk/credential/signing"
	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/goccy/go-json"
)

func SignVerifiableCredentialJWT(keyID, keyType string, privateKey, vcJSONBytes []byte) []byte {
	privKey, err := crypto.BytesToPrivKey(privateKey, crypto.KeyType(keyType))
	if err != nil {
		return nil
	}
	signer, err := crypto.NewJWTSigner(keyID, privKey)
	if err != nil {
		return nil
	}

	var cred credential.VerifiableCredential
	if err = json.Unmarshal(vcJSONBytes, &cred); err != nil {
		return nil
	}

	signedCredential, err := signing.SignVerifiableCredentialJWT(*signer, cred)
	if err != nil {
		return nil
	}
	return signedCredential
}

func VerifyVerifiableCredentialJWT(keyID, keyType string, publicKey, jwtBytes []byte) bool {
	pubKey, err := crypto.BytesToPubKey(publicKey, crypto.KeyType(keyType))
	if err != nil {
		return false
	}
	verifier, err := crypto.NewJWTVerifier(keyID, pubKey)
	if err != nil {
		return false
	}

	_, err = signing.VerifyVerifiableCredentialJWT(*verifier, string(jwtBytes))
	return err == nil
}

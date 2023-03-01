package ssi

import (
	"github.com/TBD54566975/ssi-sdk/credential"
	"github.com/TBD54566975/ssi-sdk/credential/signing"
	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/goccy/go-json"
)

// SignVerifiableCredentialJWT takes in a key ID, key type, private key, and a verifiable credential
// The keyID and KeyType are used to reconstruct a go-friendly private key to be used for signing
// the credential, which will be packaged as a JWT according to the VC-JWT 1.0 specification.
// The resulting is returned as a string representation of a JWT.
func SignVerifiableCredentialJWT(keyID, keyType string, privateKey, vcJSONBytes []byte) string {
	privKey, err := crypto.BytesToPrivKey(privateKey, crypto.KeyType(keyType))
	if err != nil {
		return ""
	}
	signer, err := crypto.NewJWTSigner(keyID, privKey)
	if err != nil {
		return ""
	}

	var cred credential.VerifiableCredential
	if err = json.Unmarshal(vcJSONBytes, &cred); err != nil {
		return ""
	}

	signedCredential, err := signing.SignVerifiableCredentialJWT(*signer, cred)
	if err != nil {
		return ""
	}
	return string(signedCredential)
}

// VerifyVerifiableCredentialJWT takes in a key ID, key type, public key, and a JWT string
// The keyID and KeyType are used to reconstruct a go-friendly public key to be used for verifying
// the JWT. The JWT is then decoded and verified, and the result is returned as a boolean.
func VerifyVerifiableCredentialJWT(keyID, keyType string, publicKey []byte, jwt string) bool {
	pubKey, err := crypto.BytesToPubKey(publicKey, crypto.KeyType(keyType))
	if err != nil {
		return false
	}
	verifier, err := crypto.NewJWTVerifier(keyID, pubKey)
	if err != nil {
		return false
	}

	_, err = signing.VerifyVerifiableCredentialJWT(*verifier, jwt)
	return err == nil
}

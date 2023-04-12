package ssi

import (
	"github.com/TBD54566975/ssi-sdk/credential"
	"github.com/TBD54566975/ssi-sdk/credential/signing"
	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/goccy/go-json"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/pkg/errors"
)

// SignVerifiableCredentialJWT takes in a did, key ID, private JWK, and a verifiable credential
// The keyID and privateJWK are used for signing the credential, which will be packaged as
// a JWT according to the VC-JWT 1.0 specification.
// The function returns a string representation of a JWT.
func SignVerifiableCredentialJWT(did string, keyID string, privateJSONWebKey []byte, vcJSONBytes []byte) (string, error) {
	key, err := jwk.ParseKey(privateJSONWebKey)
	if err != nil {
		return "", errors.Wrap(err, "parsing key")
	}

	signer, err := crypto.NewJWTSignerFromKey(did, keyID, key)
	if err != nil {
		return "", errors.Wrap(err, "creating signer")
	}

	var cred credential.VerifiableCredential
	if err = json.Unmarshal(vcJSONBytes, &cred); err != nil {
		return "", errors.Wrap(err, "unmarshalling vc")
	}

	signedCredential, err := signing.SignVerifiableCredentialJWT(*signer, cred)
	if err != nil {
		return "", errors.Wrap(err, "signing vc")
	}

	return string(signedCredential), nil
}

// VerifyVerifiableCredentialJWT takes in a did, key ID, public JWK, and a JWT string
// The keyID and publicJWK are used for verifying the JWT.
// The function returns the marshaled JSON representation of the verified Verifiable Credential.
func VerifyVerifiableCredentialJWT(did string, keyID string, publicJSONWebKey []byte, jwt string) ([]byte, error) {
	key, err := jwk.ParseKey(publicJSONWebKey)
	if err != nil {
		return nil, errors.Wrap(err, "parsing key")
	}

	verifier, err := crypto.NewJWTVerifierFromKey(did, keyID, key)
	if err != nil {
		return nil, errors.Wrap(err, "creating verifier")
	}

	_, _, vc, err := signing.VerifyVerifiableCredentialJWT(*verifier, jwt)
	if err != nil {
		return nil, errors.Wrap(err, "verifying jwt")
	}

	vcBytes, err := json.Marshal(vc)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling vc")
	}

	return vcBytes, nil
}

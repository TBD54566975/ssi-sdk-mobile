package ssi

import (
	"encoding/json"

	"github.com/TBD54566975/ssi-sdk/credential/exchange"
	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/pkg/errors"
)

/*
Parameters:

	keyID: id of key to sign resulting PresentationSubmission with
	privateJWKBytes: bytes of privateJWK to sign resulting PresentationSubmission with
	pdBytes: bytes of PresentationDefinition to build the PresentationSubmission for
	claimsBytes: bytes of an array of PresentationClaim bytes that are evaluated to potentially fulfill PresentationDefinition with
	embedTarget: target format to embed the resulting PresentationSubmission within

Returns:

	bytes of VerifiablePresentation, which embeds a PresentationSubmission within the provided embedTarget
*/
func BuildPresentationSubmission(keyID string, privateJWKBytes []byte, pdBytes []byte, claimsBytes []byte, embedTarget string) ([]byte, error) {
	key, err := jwk.ParseKey(privateJWKBytes)
	if err != nil {
		return nil, errors.Wrap(err, "parsing key")
	}

	signer, err := crypto.NewJWTSignerFromKey(keyID, key)
	if err != nil {
		return nil, errors.Wrap(err, "creating signer")
	}

	var pd exchange.PresentationDefinition
	if err := json.Unmarshal(pdBytes, &pd); err != nil {
		return nil, errors.Wrap(err, "unmarshalling PresentationDefinition")
	}

	if err := pd.IsValid(); err != nil {
		return nil, errors.Wrap(err, "invalid PresentationDefinition")
	}

	var claims []exchange.PresentationClaim
	if err := json.Unmarshal(claimsBytes, &claims); err != nil {
		return nil, errors.Wrap(err, "unmarshalling claims array")
	}

	return exchange.BuildPresentationSubmission(*signer, pd, claims, exchange.EmbedTarget(embedTarget))
}

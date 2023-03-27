package ssi

import (
	"fmt"

	"github.com/TBD54566975/ssi-sdk/credential/exchange"
	"github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/goccy/go-json"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/pkg/errors"
)

func VerifyPresentationRequestFromOIDCVPJWT(keyID string, publicJSONWebKey []byte, request []byte) ([]byte, error) {
	key, err := jwk.ParseKey(publicJSONWebKey)
	if err != nil {
		return nil, errors.Wrap(err, "parsing key")
	}

	verifier, err := crypto.NewJWTVerifierFromKey(keyID, key)
	if err != nil {
		return nil, errors.Wrap(err, "creating verifier")
	}

	pd, err := verifyPresentationDefinitionFromOIDCVPJWT(*verifier, request)
	if err != nil {
		return nil, errors.Wrap(err, "verifying presentation definition")
	}

	pdBytes, err := json.Marshal(pd)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling pd")
	}

	return pdBytes, nil
}

func verifyPresentationDefinitionFromOIDCVPJWT(verifier crypto.JWTVerifier, request []byte) (*exchange.PresentationDefinition, error) {
	parsed, err := verifier.VerifyAndParseJWT(string(request))
	if err != nil {
		return nil, errors.Wrap(err, "could not verify and parse jwt presentation request")
	}
	presDefGeneric, ok := parsed.Get("claims")
	if !ok {
		return nil, fmt.Errorf("key<%s> not found in token", "claims")
	}
	presDefGeneric, ok = presDefGeneric.(map[string]any)["vp_token"]
	if !ok {
		return nil, fmt.Errorf("key<%s> not found in token.claims", "vp_token")
	}
	presDefGeneric, ok = presDefGeneric.(map[string]any)["presentation_definition"]
	if !ok {
		return nil, fmt.Errorf("key<%s> not found in token.claims.vp_token", "presentation_definition")
	}
	presDefBytes, err := json.Marshal(presDefGeneric)
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal token into bytes for presentation definition")
	}
	var def exchange.PresentationDefinition
	if err := json.Unmarshal(presDefBytes, &def); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal token into presentation definition")
	}
	return &def, nil
}

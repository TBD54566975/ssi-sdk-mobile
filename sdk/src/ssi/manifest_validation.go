package ssi

import (
	"encoding/json"

	"github.com/TBD54566975/ssi-sdk/credential/manifest"
	"github.com/pkg/errors"
)

// IsValidCredentialApplicationForManifest validates the rules on how a credential manifest [cm] and credential
// application [ca] relate to each other https://identity.foundation/credential-manifest/#credential-application
//
// Parameters:
//
//	cmBytes: bytes of the CredentialManfiest being validated against
//	applicationAndCredsJSONBytes: bytes of the credential application and credentials as a JSON object
//
// Returns:
//
//	bytes of a JSON object, which outlines any unfulfilled inputDescriptors from the manifest
func IsValidCredentialApplicationForManifest(cmBytes []byte, applicationAndCredsJSONBytes []byte) ([]byte, error) {
	var cm manifest.CredentialManifest
	if err := json.Unmarshal(cmBytes, &cm); err != nil {
		return nil, errors.Wrap(err, "unmarshalling CredentialManifest")
	}

	applicationAndCredsJSON := make(map[string]any)
	if err := json.Unmarshal(applicationAndCredsJSONBytes, &applicationAndCredsJSON); err != nil {
		return nil, errors.Wrap(err, "unmarshalling applicationAndCredsJSON")
	}

	unfulfilledInputDescriptorsJSON, validationErr := manifest.IsValidCredentialApplicationForManifest(cm, applicationAndCredsJSON)

	unfulfilledInputDescriptorsBytes, err := json.Marshal(unfulfilledInputDescriptorsJSON)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling unfulfilledInputDescriptorsBytes")
	}

	return unfulfilledInputDescriptorsBytes, validationErr
}

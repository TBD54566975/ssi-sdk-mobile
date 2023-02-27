package ssi

import (
	"github.com/TBD54566975/ssi-sdk/credential"
)

type DIDDocumentMobile struct {
	Context              *StringArray                `json:"@context,omitempty"`
	ID                   string                      `json:"id,omitempty"`
	Controller           string                      `json:"controller,omitempty"`
	AlsoKnownAs          string                      `json:"alsoKnownAs,omitempty"`
	VerificationMethod   *VerificationMethodArray    `json:"verificationMethod,omitempty" validate:"dive"`
	Authentication       *VerificationMethodSetArray `json:"authentication,omitempty" validate:"dive"`
	AssertionMethod      *VerificationMethodSetArray `json:"assertionMethod,omitempty" validate:"dive"`
	KeyAgreement         *VerificationMethodSetArray `json:"keyAgreement,omitempty" validate:"dive"`
	CapabilityInvocation *VerificationMethodSetArray `json:"capabilityInvocation,omitempty" validate:"dive"`
	CapabilityDelegation *VerificationMethodSetArray `json:"capabilityDelegation,omitempty" validate:"dive"`
	Services             *ServiceSetArray            `json:"service,omitempty" validate:"dive"`
}

type VerifiableCredentialMobile struct {
	Context *StringArray `json:"@context" validate:"required"`
	ID      string       `json:"id,omitempty"`
	Type    *StringArray `json:"type" validate:"required"`

	// Mika: Problem type
	// Either a string or an object
	// Issuer         interface{} `json:"issuer" validate:"required"`

	IssuanceDate   string `json:"issuanceDate" validate:"required"`
	ExpirationDate string `json:"expirationDate,omitempty"`

	// Mika: Problem type
	// Requires id and type, but anything else is fair game
	// CredentialStatus  interface{}       `json:"credentialStatus,omitempty" validate:"omitempty,dive"`

	// Mika: Problem type
	// type is: map[string]interface{}
	// CredentialSubject CredentialSubject            `json:"credentialSubject" validate:"required"`

	// Mika: Problem type
	// These don't work, even though the structs only contain supported types
	CredentialSchema *credential.CredentialSchema `json:"credentialSchema,omitempty" validate:"omitempty,dive"`
	RefreshService   *credential.RefreshService   `json:"refreshService,omitempty" validate:"omitempty,dive"`

	// Mika: Problem type
	// TermsOfUse       []TermsOfUse                 `json:"termsOfUse,omitempty" validate:"omitempty,dive"`

	// Mika: Problem type
	// Evidence         []interface{}                `json:"evidence,omitempty" validate:"omitempty,dive"`

	// Mika: Problem type
	// Proof            *crypto.Proof                `json:"proof,omitempty"`
}

// Mika: Can I avoid needing to re-declare these structs?
// type CredentialSchema struct {
// 	ID   string `json:"id" validate:"required"`
// 	Type string `json:"type" validate:"required"`
// }
//
// type RefreshService struct {
// 	ID   string `json:"id" validate:"required"`
// 	Type string `json:"type" validate:"required"`
// }

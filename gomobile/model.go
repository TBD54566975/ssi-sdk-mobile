package gomobile

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

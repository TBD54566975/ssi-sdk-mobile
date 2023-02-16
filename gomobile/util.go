package gomobile

import (
	"strings"

	ssi "github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/TBD54566975/ssi-sdk/did"
)

type StringCollection interface {
	Add(s string) StringCollection
	Get(i int) string
	Size() int
}

// TODO solve this with generics
type StringArray struct {
	Items []string
}

func (sa StringArray) Add(s string) StringArray {
	sa.Items = append(sa.Items, s)
	return sa
}

func (sa StringArray) Get(i int) string {
	return sa.Items[i]
}

func (sa StringArray) Size() int {
	return len(sa.Items)
}

type VerificationMethodArray struct {
	Items []did.VerificationMethod
}

func (vma VerificationMethodArray) Add(item *did.VerificationMethod) VerificationMethodArray {
	vma.Items = append(vma.Items, *item)
	return vma
}

func (vma VerificationMethodArray) Get(i int) *did.VerificationMethod {
	return &vma.Items[i]
}

func (vma VerificationMethodArray) Size() int {
	return len(vma.Items)
}

type VerificationMethodSetArray struct {
	Items []did.VerificationMethodSet
}

func (vmsa VerificationMethodSetArray) Add(item *did.VerificationMethodSet) VerificationMethodSetArray {
	vmsa.Items = append(vmsa.Items, *item)
	return vmsa
}

func (vmsa VerificationMethodSetArray) Get(i int) *did.VerificationMethodSet {
	return &vmsa.Items[i]
}

func (vmsa VerificationMethodSetArray) Size() int {
	return len(vmsa.Items)
}

type ServiceSetArray struct {
	Items []did.Service
}

func (ssa ServiceSetArray) Add(item *did.Service) ServiceSetArray {
	ssa.Items = append(ssa.Items, *item)
	return ssa
}

func (ssa ServiceSetArray) Get(i int) *did.Service {
	return &ssa.Items[i]
}

func (ssa ServiceSetArray) Size() int {
	return len(ssa.Items)
}

func keyTypeToString(kt ssi.KeyType) string {
	return string(kt)
}

func stringToKeyType(s string) ssi.KeyType {
	switch strings.ToLower(s) {
	case "ed25519":
		return ssi.Ed25519
	case "x25519":
		return ssi.X25519
	case "secp256k1":
		return ssi.SECP256k1
	case "p-224":
		return ssi.P224
	case "p-256":
		return ssi.P256
	case "p-384":
		return ssi.P384
	case "p-521":
		return ssi.P521
	case "rsa":
		return ssi.RSA
	default:
		return ssi.KeyType(s)
	}
}

func signatureToString(s ssi.SignatureAlgorithm) string {
	return string(s)
}

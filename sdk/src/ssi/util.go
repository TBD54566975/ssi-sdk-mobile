package ssi

import (
	ssi "github.com/TBD54566975/ssi-sdk/crypto"
	"github.com/TBD54566975/ssi-sdk/did"
)

type StringArray struct {
	items []string
}

func (sa *StringArray) Add(s string) *StringArray {
	sa.items = append(sa.items, s)
	return sa
}

func (sa *StringArray) Get(i int) string {
	return sa.items[i]
}

func (sa *StringArray) Size() int {
	return len(sa.items)
}

func (sa *StringArray) toGoRepresentation() []string {
	return sa.items
}

type VerificationMethodArray struct {
	items []did.VerificationMethod
}

func (vma *VerificationMethodArray) Add(item *did.VerificationMethod) *VerificationMethodArray {
	vma.items = append(vma.items, *item)
	return vma
}

func (vma *VerificationMethodArray) Get(i int) *did.VerificationMethod {
	return &vma.items[i]
}

func (vma *VerificationMethodArray) Size() int {
	return len(vma.items)
}

type VerificationMethodSetArray struct {
	items []did.VerificationMethodSet
}

func (vmsa *VerificationMethodSetArray) Add(item *did.VerificationMethodSet) *VerificationMethodSetArray {
	vmsa.items = append(vmsa.items, *item)
	return vmsa
}

func (vmsa *VerificationMethodSetArray) Get(i int) *did.VerificationMethodSet {
	return &vmsa.items[i]
}

func (vmsa *VerificationMethodSetArray) Size() int {
	return len(vmsa.items)
}

type ServiceSetArray struct {
	items []did.Service
}

func (ssa *ServiceSetArray) Add(item *did.Service) *ServiceSetArray {
	ssa.items = append(ssa.items, *item)
	return ssa
}

func (ssa *ServiceSetArray) Get(i int) *did.Service {
	return &ssa.items[i]
}

func (ssa *ServiceSetArray) Size() int {
	return len(ssa.items)
}

func keyTypeToString(kt ssi.KeyType) string {
	return string(kt)
}

func stringToKeyType(s string) ssi.KeyType {
	return ssi.KeyType(s)
}

func signatureToString(s ssi.SignatureAlgorithm) string {
	return string(s)
}

package util

import (
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
	items []did.VerificationMethodSet
}

func (vmsa VerificationMethodSetArray) Add(item *did.VerificationMethodSet) VerificationMethodSetArray {
	vmsa.items = append(vmsa.items, *item)
	return vmsa
}

func (vmsa VerificationMethodSetArray) Get(i int) *did.VerificationMethodSet {
	return &vmsa.items[i]
}

func (vmsa VerificationMethodSetArray) Size() int {
	return len(vmsa.items)
}

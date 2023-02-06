package crypto

import (
	ssi "github.com/TBD54566975/ssi-sdk/crypto"
)

func KeyTypeToString(kt ssi.KeyType) string {
	return string(kt)
}

func StringToKeyType(s string) ssi.KeyType {
	return ssi.KeyType(s)
}

func SignatureToString(s ssi.SignatureAlgorithm) string {
	return string(s)
}

func StringToSignature(s string) ssi.SignatureAlgorithm {
	return ssi.SignatureAlgorithm(s)
}

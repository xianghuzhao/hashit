package kdfcrypt

import (
	"fmt"

	"golang.org/x/crypto/hkdf"
)

// HKDF params.
type HKDF struct {
	HashFunc string `param:"hash"`
	Info     string `param:"info"`
}

// SetDefaultParam sets the default param for hkdf
func (kdf *HKDF) SetDefaultParam() {
	if kdf.HashFunc == "" {
		kdf.HashFunc = "sha512"
	}
}

// Generate hash with hkdf.
func (kdf *HKDF) Generate(key, salt []byte, hashLength uint32) ([]byte, error) {
	hashFunc, ok := hashFuncMap[kdf.HashFunc]
	if !ok {
		return nil, fmt.Errorf("Hash func for HKDF is not valid: %s", kdf.HashFunc)
	}

	reader := hkdf.New(hashFunc, []byte(key), salt, []byte(kdf.Info))

	hashed := make([]byte, hashLength)
	_, err := reader.Read(hashed)

	return hashed, err
}
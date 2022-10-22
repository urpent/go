package randcode

import (
	"crypto/rand"
	"math/big"
)

const alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-"

// GenerateSecureCode generates random 15 digit alphanumeric code with crypto/rand
func GenerateSecureCode(codeLength int) string {
	ret := make([]byte, codeLength)
	for i := 0; i < codeLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphaNumeric))))
		if err != nil {
			panic("GenerateSecureCode " + err.Error()) // return "", err
		}
		ret[i] = alphaNumeric[num.Int64()]
	}

	return string(ret)
}

package randcode

import (
	"crypto/rand"
	"math/big"
)

type Combination string

const CombinationDefault Combination = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-"

// Secure generates random code with crypto/rand
func Secure(codeLength int) (string, error) {
	ret := make([]byte, codeLength)
	for i := 0; i < codeLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(CombinationDefault))))
		if err != nil {
			return "", err
		}
		ret[i] = CombinationDefault[num.Int64()]
	}

	return string(ret), nil
}

// CustomSecure generates random code with crypto/rand and custom combination
func CustomSecure(codeLength int, combination Combination) (string, error) {
	ret := make([]byte, codeLength)
	for i := 0; i < codeLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(combination))))
		if err != nil {
			return "", err
		}
		ret[i] = combination[num.Int64()]
	}

	return string(ret), nil
}

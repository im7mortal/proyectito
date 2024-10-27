package keygen

import (
	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

// PublicKeyGenerator defines the interface for key generation
type PublicKeyGenerator interface {
	GeneratePublicKey(privateKey string) (string, error)
	GeneratePrivateKey(name, email string, bitLength int) (string, error)
}

// KeyGenerator is the implementation of PublicKeyGenerator interface
type KeyGenerator struct{}

// NewKeyGenerator creates a new instance of KeyGenerator
func NewKeyGenerator() *KeyGenerator {
	return &KeyGenerator{}
}

// GeneratePublicKey takes an armored private key and returns an armored public key
func (kg *KeyGenerator) GeneratePublicKey(privateKey string) (string, error) {
	keyObj, err := crypto.NewKeyFromArmored(privateKey)
	if err != nil {
		return "", err
	}

	pubKeyArmored, err := keyObj.GetArmoredPublicKey()
	if err != nil {
		return "", err
	}

	return pubKeyArmored, nil
}

// GeneratePrivateKey generates a new RSA private key with given name, email, and bit length
func (kg *KeyGenerator) GeneratePrivateKey(name, email string, bitLength int) (string, error) {
	privateKeyObj, err := crypto.GenerateKey(name, email, "rsa", bitLength)
	if err != nil {
		return "", err
	}

	privateKeyArmored, err := privateKeyObj.Armor()
	if err != nil {
		return "", err
	}

	return privateKeyArmored, nil
}

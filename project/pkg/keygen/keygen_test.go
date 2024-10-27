package keygen

import (
	"testing"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	keyGen := NewKeyGenerator()

	privateKey, err := keyGen.GeneratePrivateKey("Test User", "test@example.com", 2048)
	assert.NoError(t, err, "Failed to generate private key")
	assert.NotEmpty(t, privateKey, "Expected non-empty private key")

	privateKeyObj, err := crypto.NewKeyFromArmored(privateKey)
	assert.NoError(t, err, "Failed to parse generated private key")
	assert.True(t, privateKeyObj.IsPrivate(), "Expected RSA private key")
}

func TestGeneratePublicKeyFromGeneratedPrivateKey(t *testing.T) {
	keyGen := NewKeyGenerator()

	privateKey, err := keyGen.GeneratePrivateKey("Test User", "test@example.com", 2048)
	assert.NoError(t, err, "Failed to generate private key")

	publicKey, err := keyGen.GeneratePublicKey(privateKey)
	assert.NoError(t, err, "Failed to generate public key")
	assert.NotEmpty(t, publicKey, "Expected non-empty public key")

	publicKeyObj, err := crypto.NewKeyFromArmored(publicKey)
	assert.NoError(t, err, "Failed to parse generated public key")
	assert.False(t, publicKeyObj.IsPrivate(), "Expected RSA public key")
}

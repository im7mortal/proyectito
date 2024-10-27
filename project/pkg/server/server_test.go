package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/im7mortal/project/pkg/keygen"
	"github.com/stretchr/testify/assert"
)

func setupRouter(pkGen keygen.PublicKeyGenerator) *gin.Engine {
	sdk := New(pkGen, "")
	router := sdk.GetMainEngine().(*gin.Engine)
	return router
}

func TestGeneratePrivateKeyHandler(t *testing.T) {
	pkGen := keygen.NewKeyGenerator()
	router := setupRouter(pkGen)

	requestBody := map[string]interface{}{
		"name":       "Test User",
		"email":      "test@example.com",
		"bit_length": 2048,
	}
	requestBodyJSON, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/v1/generate-private-key", bytes.NewBuffer(requestBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var responseBody map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	privateKey, exists := responseBody["private_key"]
	assert.True(t, exists, "Expected 'private_key' in response")
	assert.NotEmpty(t, privateKey, "Expected non-empty private key")
}

func TestExtractPublicKeyHandler(t *testing.T) {
	pkGen := keygen.NewKeyGenerator()
	router := setupRouter(pkGen)

	privateKey, err := pkGen.GeneratePrivateKey("Test User", "test@example.com", 2048)
	assert.NoError(t, err, "Failed to generate private key")

	requestBody := map[string]interface{}{
		"private_key": privateKey,
	}
	requestBodyJSON, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/v1/extract-public-key", bytes.NewBuffer(requestBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var responseBody map[string]string
	err = json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	publicKey, exists := responseBody["public_key"]
	assert.True(t, exists, "Expected 'public_key' in response")
	assert.NotEmpty(t, publicKey, "Expected non-empty public key")
}

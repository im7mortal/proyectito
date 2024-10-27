package server

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/im7mortal/project/pkg/keygen"
)

type sdkGin struct {
	pkGen    keygen.PublicKeyGenerator
	sitePath string
}

type MainEngine interface {
	GetMainEngine() http.Handler
}

// Create sdk server
func New(pkGen keygen.PublicKeyGenerator, sitePath string) MainEngine {
	return &sdkGin{
		pkGen:    pkGen,
		sitePath: sitePath,
	}
}

func (sdk *sdkGin) GetMainEngine() http.Handler {
	router := gin.Default()
	router.Use(setCORS())

	// k8s probes
	router.GET("/readiness", sdk.getProbe("readiness"))
	router.GET("/liveness", sdk.getProbe("liveness"))

	// API group
	v1 := router.Group("/v1")
	v1.POST("/extract-public-key", sdk.extractPublicKeyHandler)
	v1.POST("/generate-private-key", sdk.generatePrivateKeyHandler)

	router.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(sdk.sitePath, "index.html"))
	})

	return router
}

// Set CORS policy
func setCORS() gin.HandlerFunc {
	d := cors.Config{
		AllowMethods:     []string{http.MethodOptions, "OPTIONS", "GET", "HEAD", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept-Language"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}
	return cors.New(d)
}

func (sdk *sdkGin) getProbe(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}

// Handler to extract public key from provided private key
func (sdk *sdkGin) extractPublicKeyHandler(c *gin.Context) {
	var requestBody struct {
		PrivateKey string `json:"private_key"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	pubKey, err := sdk.pkGen.GeneratePublicKey(requestBody.PrivateKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate public key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"public_key": pubKey})
}

// New handler to generate a new private key
func (sdk *sdkGin) generatePrivateKeyHandler(c *gin.Context) {
	var requestBody struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		BitLength int    `json:"bit_length"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	privateKey, err := sdk.pkGen.GeneratePrivateKey(requestBody.Name, requestBody.Email, requestBody.BitLength)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate private key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"private_key": privateKey})
}

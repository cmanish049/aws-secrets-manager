package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"secrets-manager-platform/internal/services"
)

type SecretsHandler struct {
	service *services.SecretsService
}

func NewSecretsHandler(service *services.SecretsService) *SecretsHandler {
	return &SecretsHandler{service: service}
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// SuccessResponse represents a success message response
type SuccessResponse struct {
	Message string `json:"message" example:"operation successful"`
}

// ListSecrets godoc
// @Summary      List all secrets
// @Description  Returns a list of all secret names and descriptions from AWS Secrets Manager
// @Tags         secrets
// @Accept       json
// @Produce      json
// @Success      200  {array}   services.Secret
// @Failure      500  {object}  ErrorResponse
// @Router       /secrets [get]
func (h *SecretsHandler) ListSecrets(c *gin.Context) {
	secrets, err := h.service.ListSecrets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, secrets)
}

// GetSecret godoc
// @Summary      Get a secret value
// @Description  Returns the value of a specific secret by name
// @Tags         secrets
// @Accept       json
// @Produce      json
// @Param        name  path      string  true  "Secret name"
// @Success      200   {object}  services.Secret
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /secrets/{name} [get]
func (h *SecretsHandler) GetSecret(c *gin.Context) {
	name := strings.TrimPrefix(c.Param("name"), "/")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "secret name is required"})
		return
	}

	secret, err := h.service.GetSecret(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, secret)
}

// CreateSecretRequest represents the request body for creating a secret
type CreateSecretRequest struct {
	Name        string `json:"name" binding:"required" example:"prd/database"`
	Value       string `json:"value" binding:"required" example:"USERNAME=admin\nPASSWORD=secret123"`
	Description string `json:"description" example:"Production database credentials"`
}

// CreateSecret godoc
// @Summary      Create a new secret
// @Description  Creates a new secret in AWS Secrets Manager
// @Tags         secrets
// @Accept       json
// @Produce      json
// @Param        secret  body      CreateSecretRequest  true  "Secret to create"
// @Success      201     {object}  SuccessResponse
// @Failure      400     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /secrets [post]
func (h *SecretsHandler) CreateSecret(c *gin.Context) {
	var req CreateSecretRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	secret := services.Secret{
		Name:        req.Name,
		Value:       req.Value,
		Description: req.Description,
	}

	if err := h.service.CreateSecret(c.Request.Context(), secret); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "secret created successfully"})
}

// UpdateSecretRequest represents the request body for updating a secret
type UpdateSecretRequest struct {
	Value string `json:"value" binding:"required" example:"USERNAME=admin\nPASSWORD=newsecret456"`
}

// UpdateSecret godoc
// @Summary      Update a secret
// @Description  Updates the value of an existing secret in AWS Secrets Manager
// @Tags         secrets
// @Accept       json
// @Produce      json
// @Param        name    path      string               true  "Secret name"
// @Param        secret  body      UpdateSecretRequest  true  "New secret value"
// @Success      200     {object}  SuccessResponse
// @Failure      400     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /secrets/{name} [put]
func (h *SecretsHandler) UpdateSecret(c *gin.Context) {
	name := strings.TrimPrefix(c.Param("name"), "/")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "secret name is required"})
		return
	}

	var req UpdateSecretRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateSecret(c.Request.Context(), name, req.Value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "secret updated successfully"})
}

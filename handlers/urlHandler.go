package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pedrolfr404/encurtador-url-go/models"
	"github.com/pedrolfr404/encurtador-url-go/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type shortenRequest struct {
	URL string `json:"url" binding:"required"`
}

type URLHandler struct {
	Collection *mongo.Collection
	BaseURL    string
}

func (h *URLHandler) ShortenURL(c *gin.Context) {
	var req shortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload"})
		return
	}

	if !utils.IsValid(req.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	if utils.IsAlreadyShortened(req.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL Already Shortened or blocked domain"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existingURL models.URL
	err := h.Collection.FindOne(ctx, bson.M{"original_url": req.URL}).Decode(&existingURL)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"short_url":    h.BaseURL + existingURL.ID,
			"original_url": existingURL.OriginalUrl,
			"reused":       true,
		})
		return
	}

	shortID, err := utils.GeneratorShortId(6, req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	newURL := models.URL{
		ID:          shortID,
		OriginalUrl: req.URL,
		CreatedAt:   time.Now(),
	}

	_, err = h.Collection.InsertOne(ctx, newURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving to the database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_url":    h.BaseURL + newURL.ID,
		"original_url": newURL.OriginalUrl,
		"reused":       false,
	})

}

func (h *URLHandler) Redirect(c *gin.Context) {
	shortID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var urlDoc models.URL
	err := h.Collection.FindOne(ctx, bson.M{"shortenUrl": shortID}).Decode(&urlDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Shorten URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving the URL"})
		return
	}

	c.Redirect(http.StatusFound, urlDoc.OriginalUrl)
}

package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pedrolfr404/encurtador-url-go/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

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

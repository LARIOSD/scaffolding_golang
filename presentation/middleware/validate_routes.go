package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Validatable interface {
	Validate() error
}

func ValidateBody(obj Validatable) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalid", "details": err.Error(), "success": false, "status": http.StatusBadRequest})
			c.Abort()
			return
		}
		if err := obj.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed validation", "details": err.Error(), "success": false, "status": http.StatusBadRequest})
			c.Abort()
			return
		}
		c.Set("body", obj)
		c.Next()
	}
}

func ValidateQuery(obj Validatable) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindQuery(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Query invalid", "details": err.Error(), "success": false, "status": http.StatusBadRequest})
			c.Abort()
			return
		}
		if err := obj.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed query validation", "details": err.Error(), "success": false, "status": http.StatusBadRequest})
			c.Abort()
			return
		}
		c.Set("query", obj)
		c.Next()
	}
}

func ValidateParams(obj Validatable) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindUri(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters", "details": err.Error(), "success": false, "status": http.StatusBadRequest})
			c.Abort()
			return
		}
		if err := obj.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter validation failed", "details": err.Error(), "success": false, "status": http.StatusBadRequest})
			c.Abort()
			return
		}
		c.Set("params", obj)
		c.Next()
	}
}

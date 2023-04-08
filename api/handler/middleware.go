package handler

import (
	"app/pkg/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// r.Use(AuthMiddleware())

func (h *Handler) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		value := c.GetHeader("Authorization")

		info, err := helper.ParseClaims(value, h.cfg.SecretKey)
		fmt.Println(value)
		fmt.Println(info)
		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}

		c.Set("Auth", info)

		c.Next()
	}
}

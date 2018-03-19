package content

import (
	"github.com/gin-gonic/gin"
)

func ArticlesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(200, gin.H{
				"id": "No ID provided",
			})
		} else {
			c.JSON(200, gin.H{
				"id": id,
			})
		}
	}
}

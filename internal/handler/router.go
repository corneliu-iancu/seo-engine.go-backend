// todo: add docs.
package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ServerInterface interface {
	AddRule(c *gin.Context)
	FindAllSegments(c *gin.Context)
	GetRules(c *gin.Context)
	GetMatch(c *gin.Context)
}

// Gin Hanlder function - applies our app routing.
func GinHandler(si ServerInterface) func(*gin.Engine) {
	return func(router *gin.Engine) {
		fmt.Println("[DEBUG] 💡 Register handlers on gin engine.")
		// setup router paths.
		v1 := router.Group("/api/v1")
		{ // v1 block.
			// Handler for retrieving all rules in a human readable format.
			v1.Handle("GET", "/rules", si.GetRules)

			// Handler for retrieving all segments stored in the database.
			v1.Handle("GET", "/rules/segments", si.FindAllSegments)

			// Handler for registering a new seo rule based on a URI parameter.
			v1.Handle("GET", "/rules/add", si.AddRule)

			// Handler for registering a new seo rule based on a URI parameter.
			v1.Handle("GET", "/match", si.GetMatch)
		}
	}
}

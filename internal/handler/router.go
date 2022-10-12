// todo: add docs.
package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Gin Hanlder function - applies our app routing.
func GinHandler() func(*gin.Engine) {
	return func(router *gin.Engine) {
		fmt.Println("[DEBUG] 💡 Register handlers on gin engine.")
		// setup router paths.
		v1 := router.Group("/api/v1")
		{ //v1 block.
			// Handler for searching a matching rule based on a URI parameter.
			v1.Handle("GET", "/search", SearchRuleMatch)

			// Handler for registering a new seo rule based on a URI parameter.
			v1.Handle("POST", "/rule", AddSeoRule)
		}
	}
}

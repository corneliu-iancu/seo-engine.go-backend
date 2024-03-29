// todo: add docs.
package handler

import (
	"github.com/gin-gonic/gin"
)

type ServerInterface interface {
	AddRule(c *gin.Context)
	GetRules(c *gin.Context)
	GetMatch(c *gin.Context)
	GetUrlBySegmentId(c *gin.Context)
}

// Gin Hanlder function - applies our app routing.
func GinHandler(si ServerInterface) func(*gin.Engine) {
	return func(router *gin.Engine) {
		v1 := router.Group("/api/v1")
		{ // v1 block.
			// Handler for retrieving all rules in a human readable format.
			v1.Handle("GET", "/rules", si.GetRules)

			v1.Handle("GET", "/rules/by-segment-id", si.GetUrlBySegmentId)

			// Handler for registering a new seo rule based on a URI parameter.
			v1.Handle("POST", "/rules", si.AddRule)

			// Handler for registering a new seo rule based on a URI parameter.
			v1.Handle("GET", "/rules/match", si.GetMatch)
		}
	}
}

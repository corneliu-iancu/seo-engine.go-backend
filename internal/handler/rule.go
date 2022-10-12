// Rule controller.
// Implements custom logic for the SEO Engine.

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/metadata"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
)

// Adds a new record to the Seo Rules Database.
// todo: add docs.
func AddSeoRule(ctx *gin.Context) {
	// c.IndentedJSON(http.StatusOK, "OK")
	metadata.MetadataPack = append(metadata.MetadataPack, metadata.Metadata{
		Title:        "title",
		Header:       "header",
		Description:  "this is description",
		CanonicalURL: "this is canonical url",
		TextBlock:    "this si a text block",
	})
	ctx.IndentedJSON(http.StatusOK, "{\"SEO\"}")
}

// Search a Rule based on a request query string.
// todo: add docs.
func SearchRuleMatch(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, rule.RulePack)
}

// Based on input, returns the selected SEO rule.
// todo: add docs.
func Metadata(ctx *gin.Context) {
	// @todo: implement.
}

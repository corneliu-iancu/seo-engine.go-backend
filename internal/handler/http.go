// ================================================
// HTTP Handlers
// Implements router.go ServerInterface methods.
// ================================================
package handler

import (
	"fmt"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/metadata"
	"net/http"
	"net/url"

	"github.com/corneliu-iancu/seo-engine.go-backend/internal/app"
	"github.com/gin-gonic/gin"
)

type HttpControllers struct {
	app app.BusinessLogic
}

type RulePayload struct {
	Uri      string             `json:"uri"`
	Title    string             `json:"title"`
	MetaTags []metadata.MetaTag `json:"metaTags"`
}

// Creates new http controllers instance.
func NewHttpControllers(app app.BusinessLogic) *HttpControllers {
	return &HttpControllers{
		app: app,
	}
}

// Persist one rule to the database.
func (hc HttpControllers) AddRule(ctx *gin.Context) {
	// let's read out the input of the user by get parameter
	var rulePayload RulePayload
	if err := ctx.ShouldBindJSON(&rulePayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := url.Parse(rulePayload.Uri)
	if err != nil {
		panic(err) //@todo: handle errors mechanism.
	}

	meta := metadata.Metadata{
		Title:    rulePayload.Title,
		MetaTags: rulePayload.MetaTags,
	}

	result, err := hc.app.CreateRule(u, meta)
	if err != nil {
		fmt.Println(err)
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

// Returns rules that matches a given uri as a GET param.
func (hc HttpControllers) GetMatch(ctx *gin.Context) {
	uri, _ := ctx.GetQuery("uri")
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}

	result, _ := hc.app.GetMatch(u)
	ctx.IndentedJSON(http.StatusOK, result)
}

// Returns all rules
func (hc HttpControllers) GetRules(ctx *gin.Context) {
	rules, _ := hc.app.GetAllRules()
	ctx.IndentedJSON(http.StatusOK, rules)
}

func (hc HttpControllers) GetUrlBySegmentId(ctx *gin.Context) {
	segmentId, exists := ctx.GetQuery("id")
	if exists != true {
		panic("missing query identifier")
	}

	segments, err := hc.app.GetURLBySegmentId(segmentId)
	if err != nil {
		panic(err)
	}

	ctx.IndentedJSON(http.StatusOK, segments)
}

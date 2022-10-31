// todo: add docs.
package handler

import (
	"net/http"
	"net/url"

	"github.com/corneliu-iancu/seo-engine.go-backend/internal/app"
	"github.com/gin-gonic/gin"
)

type HttpControllers struct {
	app app.App
}

// Creates new http controllers instance.
func NewHttpControllers(app app.App) *HttpControllers {
	// todo: implement me.
	return &HttpControllers{
		app: app,
	}
}

// Returns all db segments data.
func (hc HttpControllers) FindAllSegments(ctx *gin.Context) {
	segments, _ := hc.app.GetAllSegments() // @todo: rename to find all segments.
	ctx.IndentedJSON(http.StatusOK, segments)
}

func (hc HttpControllers) GetRules(ctx *gin.Context) {

	// demo purposes.

	rules, _ := hc.app.GetAllRules()
	ctx.IndentedJSON(http.StatusOK, rules)
}

// Persist one rule to the database.
func (hc HttpControllers) AddRule(ctx *gin.Context) {
	// let's read out the input of the user by get parameter
	uri, _ := ctx.GetQuery("uri")
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}

	// @todo: only send the URI. The parse should go in app.
	result, _ := hc.app.AddRule(u)

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

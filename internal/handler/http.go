// ================================================
// HTTP Handlers
// Implements router.go ServerInterface methods.
// ================================================
package handler

import (
	"net/http"
	"net/url"

	"github.com/corneliu-iancu/seo-engine.go-backend/internal/app"
	"github.com/gin-gonic/gin"
)

type HttpControllers struct {
	app app.BusinessLogic
}

// Creates new http controllers instance.
func NewHttpControllers(app app.BusinessLogic) *HttpControllers {
	return &HttpControllers{
		app: app,
	}
}

// Returns all rules
func (hc HttpControllers) GetRules(ctx *gin.Context) {
	rules, _ := hc.app.GetAllRules()
	ctx.IndentedJSON(http.StatusOK, rules)
}

// Persist one rule to the database.
func (hc HttpControllers) AddRule(ctx *gin.Context) {
	// let's read out the input of the user by get parameter
	uri, _ := ctx.GetQuery("uri")
	u, err := url.Parse(uri)
	if err != nil {
		panic(err) //@todo: handle errors mechanism.
	}

	result, _ := hc.app.CreateRule(u)
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

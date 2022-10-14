// todo: add docs.
package handler

import (
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/app"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (hc HttpControllers) GetAllRules(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, rule.SegmentPack)
}

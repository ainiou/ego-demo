package http

import (
	"ego-demo/internal/server/service/example"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Option struct {
	Service *example.Service
}
type Handler struct {
	*Option
}

func NewHandler(opt *Option) *Handler {
	return &Handler{opt}
}

func (h *Handler) RegisterRouter(g gin.IRouter) {
	g.GET("/server", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"errmsg": "success"})
	})

}

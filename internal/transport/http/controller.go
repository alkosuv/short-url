package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gen95mis/short-url/internal/model"
	"github.com/gen95mis/short-url/internal/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	s *service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{
		s: s,
	}
}

func (c *Controller) AddNewLink(ctx *gin.Context) {
	url := new(model.URL)
	if err := json.NewDecoder(ctx.Request.Body).Decode(url); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	hash, err := c.s.Set(url.Original)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	shortenedUrl := fmt.Sprintf("http://%s/%s", ctx.Request.Host, hash)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, shortenedUrl)
}

func (c *Controller) GetLink(ctx *gin.Context) {
	hash, ok := ctx.Params.Get("hash")
	if !ok {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	oridinal, err := c.s.Get(hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.Redirect(http.StatusFound, oridinal)
}

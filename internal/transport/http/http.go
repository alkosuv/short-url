package http

import (
	"github.com/gen95mis/short-url/internal/service"
	"github.com/gin-gonic/gin"
)

func Service(s *service.Service) error {
	c := NewController(s)
	r := gin.Default()

	r.POST("/short", c.AddNewLink)
	r.GET("/:hash", c.GetLink)

	return r.Run(":80")
}

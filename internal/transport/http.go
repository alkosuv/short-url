package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gen95mis/short-url/internal/model"
	"github.com/gen95mis/short-url/internal/service"
	"github.com/gin-gonic/gin"
)

func Service(s *service.Service) error {
	r := gin.Default()

	r.POST("/short", func(ctx *gin.Context) {
		url := new(model.URL)
		if err := json.NewDecoder(ctx.Request.Body).Decode(url); err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		hash, err := s.Set(url.Original)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		shortenedUrl := fmt.Sprintf("http://localhost/%s", hash)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		ctx.JSON(http.StatusOK, shortenedUrl)
	})

	r.GET("/:hash", func(ctx *gin.Context) {
		hash, ok := ctx.Params.Get("hash")
		if !ok {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		oridinal, err := s.Get(hash)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		ctx.Redirect(http.StatusFound, oridinal)
	})

	return r.Run(":80")
}

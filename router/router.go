package router

import (
	"net/http"

	"github.com/anugrahwl/nearby-places/controller"
	"github.com/anugrahwl/nearby-places/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const SWAGGER_URL_REMOTE = "https://gist.githubusercontent.com/anugrahwl/5e96b1e599eeb56ad3de5553739718a7/raw/4feb86dddec748bd05c3d273ebc855bd5fc6b906/swagger.yaml"

// const SWAGGER_LOCAL = "./swagger-config.yaml"

func SetupRouter(data []models.CityPlace) *gin.Engine {
	r := gin.Default()

	// attach BatchPlace to gin Context
	r.Use(func(ctx *gin.Context) {
		ctx.Set("places", data)
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})

	r.GET("/search", controller.GetNearby)

	// load swagger from remote url
	conf := ginSwagger.URL(SWAGGER_URL_REMOTE)

	// swagger route
	r.GET("/swagger/*any", ginSwagger.
		WrapHandler(
			swaggerFiles.Handler,
			conf,
		))

	return r
}

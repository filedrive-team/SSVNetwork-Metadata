package routers

import (
	"ssv_operator_metadata/controllers"
	"ssv_operator_metadata/middleware"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter(control *controllers.Control) *gin.Engine {
	router := gin.Default()

	// router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(middleware.Cors())

	appV1 := router.Group("/api/v1")
	{
		appV1.GET("/operator/getoperator", control.GetOperator)
		appV1.GET("/operator/list", control.ListOperator)
		appV1.POST("/operator/update", control.UpdateOperator)

		appV1.POST("/operator/uploadimage", control.UploadImage)
		appV1.GET("/operator/gettimestamp", control.GetTimestamp)
		appV1.GET("/operator/getoperatorbykeyword", control.GetOperatorByKeyword)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, controllers.CommonOkResp{Code: 404, Msg: "Not found"})
	})

	return router
}

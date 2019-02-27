package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"svc-china-divisions/apis/district"
	_ "svc-china-divisions/docs"
)

const districtDataDir = "./district_data/dist/"

// @title 🇨🇳中国行政区查询服务
// @version 1.0
// @description 中国行政区域查询，数据来源：https://github.com/modood/Administrative-divisions-of-China
func main() {
	r := setupRouter()
	r.Run(":9001")
}

func setupRouter() *gin.Engine {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "行政区划微服务",
			"docs":    "/docs/index.html",
		})
	})

	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.GET("/district", district.Get)
	return route
}

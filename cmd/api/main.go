package main

import (
	"fmt"
	getGeoIp "pebu-go-demo/api/geo-ip"
	_ "pebu-go-demo/docs"
	"pebu-go-demo/external-services/maxmind"
	"pebu-go-demo/internal/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Example Gin Swagger API
// @version 1.0
// @description Sample Swagger integration with Gin
// @host localhost:8080
// @BasePath /api/v1
func main() {
	logger := logger.New()
	r := gin.Default()

	// Swagger integration
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoints handlers
	api := r.Group("/api/v1")

	// GeoIp
	maxmind := maxmind.New(logger)
	getGeoIpEndpoint := getGeoIp.New(logger, maxmind)
	api.GET(fmt.Sprintf("/%v", getGeoIp.EndpointName), getGeoIpEndpoint.Handler)

	r.Run(":8080")
}

package iplookup

import (
	"fmt"
	"net/http"

	"pebu-go-demo/external-services/maxmind"
	"pebu-go-demo/internal/logger"

	"github.com/gin-gonic/gin"
)

type GetGeoIp struct {
	logger  logger.Logger
	maxmind maxmind.Maxmind
}

func New(logger logger.Logger, maxmind maxmind.Maxmind) GetGeoIp {
	return GetGeoIp{
		logger:  logger,
		maxmind: maxmind,
	}
}

const IpAddressParameterName = "ip-address"

var EndpointName = fmt.Sprintf("ip-lookup/:%v", IpAddressParameterName)

// @Summary Lookup IP Address
// @Description Returns a geoIp result
// @Tags example
// @Accept json
// @Produce json
// @Param ip-address path int true "ip-address"
// @Success 200 {object} map[string]string
// @Router /ip-lookup/{ip-address} [get]
func (g *GetGeoIp) Handler(httpContext *gin.Context) {
	g.getGeoIp(httpContext)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (g *GetGeoIp) getGeoIp(httpContext *gin.Context) {
	ctx := httpContext.Request.Context()
	ipAddress := httpContext.Param(IpAddressParameterName)
	bearerToken := httpContext.Request.Header.Get("Authorization")

	g.logger.Info(ctx, "user looking up ip address", map[string]any{
		"ipAddress": ipAddress,
	})

	response, err := g.maxmind.GetGeoIpData(ctx, ipAddress, bearerToken)
	if err != nil {
		httpContext.IndentedJSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}

	httpContext.IndentedJSON(http.StatusOK, response)
}

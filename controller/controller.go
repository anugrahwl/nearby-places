package controller

import (
	"net/http"
	"strconv"

	"github.com/anugrahwl/nearby-places/models"
	"github.com/gin-gonic/gin"
)

func GetNearby(ctx *gin.Context) {
	places := ctx.MustGet("places").([]models.CityPlace)

	lat, _ := strconv.ParseFloat(ctx.Query("latitude"), 64)
	lon, _ := strconv.ParseFloat(ctx.Query("longitude"), 64)
	cid, _ := strconv.ParseUint(ctx.Query("category_id"), 10, 8)

	q := models.WebQuery{
		Latitude:   lat,
		Longitude:  lon,
		CategoryId: uint(cid),
	}

	result, _ := models.GetNearbyPlaces(q, places)

	ctx.JSON(http.StatusOK, gin.H{
		"data":  result,
		"total": len(result),
	})
}

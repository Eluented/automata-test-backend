package main

import (
	"github.com/gin-gonic/gin"
)

type incubator struct {
	Microplates        string `json:microplates`
	TempRange          string `json:tempRange`
	HumidityRange      string `json:humidityRange`
	IncubatorPlacement string `json:incubatorPlacement`
	ExternalControlBox string `json:externalControlBox`
}

// setup router and handle mainpoints of API
func main() {
	router := gin.Default()
	router.Run("localhost:8080")
}

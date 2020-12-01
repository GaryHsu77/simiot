package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/MOXA-ISD/sim-iot-city/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var general = `
{
	"description": "myDevice",
	"modelName": "UC-8220-T-LX",
	"serialNumber": "ABC123",
	"firmwareVersion": "1.0",
	"cpu": "ARMv7 Processor rev 5 (v7l)",
	"memorySize": 2116042752,
	"disk": [
	  {
		"name": "System",
		"mount": "/"
	  }
	]
}
`

func getGeneral(c *gin.Context) {
	c.String(http.StatusOK, general)
}

func setGeneral(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(server.HTTPErrorBadRequest("invalid request, err: %s", err.Error()))
	}

	var ii interface{}
	if err := json.Unmarshal(data, &ii); err != nil {
		panic(server.HTTPErrorBadRequest("invalid request, err: %s", err.Error()))
	}

	general = string(data)
	c.String(http.StatusOK, general)
}

func main() {
	route := gin.Default()
	server := &http.Server{Handler: route}
	server.Addr = ":59000"

	api := route.Group("api/v1")
	{
		api.GET("/device/general", getGeneral)
		api.PUT("/device/general", setGeneral)
	}

	logrus.Infoln("Listening and serving HTTP on tcp:", server.Addr)
	server.ListenAndServe()
}

package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/maiacodes/slashy/analytics"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/maiacodes/slashy/commands"
	"github.com/maiacodes/slashy/interaction"
)

func main() {
	logrus.Info("Starting...")
	go commands.InitialiseCommands()
	go analytics.Connect()
	e := echo.New()

	e.POST("/hook", hook)

	e.HideBanner = true
	e.Start(":" + os.Getenv("PORT"))
}

func hook(c echo.Context) error {
	rawBody, _ := ioutil.ReadAll(c.Request().Body)

	// Process timestamp
	timestamp := c.Request().Header.Get("X-Signature-Timestamp")
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return echo.NewHTTPError(500, "cannot process timestamp")
	}
	tm := time.Unix(i, 0)
	if time.Now().Sub(tm).Seconds() > 10 {
		return echo.NewHTTPError(401, "bad timestamp")
	}

	// Verify signature
	if !verify(c.Request().Header.Get("X-Signature-Ed25519"), timestamp+string(rawBody), os.Getenv("public_key")) {
		return echo.NewHTTPError(401, "bad signature")
	}

	// Decode body
	var body partialEvent
	if json.Unmarshal(rawBody, &body) != nil {
		return c.NoContent(500)
	}

	if body.Type == 2 {
		var event interaction.Event
		err := json.Unmarshal(rawBody, &event)
		if err != nil {
			logrus.Info("ouch " + err.Error())
			return c.NoContent(500)
		}
		response := commands.Router(&event)
		if response != nil {
			event.Responded = true
			return c.JSON(200, response)
		}
	}

	return c.JSON(200, map[string]interface{}{
		"type": 1,
	})
}

type partialEvent struct {
	Type int `json:"type"`
}

package server

import (
	"encoding/json"
	"net/http"

	"io"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/thehivecorporation/raccoon"
	"github.com/thehivecorporation/raccoon/parser"
)

//REST is the server that is launched when a user selects the "server" option
//in the CLI
func REST(c *cli.Context) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/", func(c echo.Context) error {
		req, err := parseRequest(c.Request().Body())

		jobParser := parser.JobParser{Dispatcher: new(raccoon.SimpleDispatcher)}

		taskList, err := jobParser.ParseTaskList(req.TaskList)
		if err != nil {
			return err
		}

		jobs := jobParser.BuildJobList(req.Infrastructure, taskList)

		//Send jobs to dispatcher
		jobParser.Dispatcher.Dispatch(*jobs)

		rsp := struct {
			Status string
		}{
			Status: "ok",
		}

		return c.JSON(http.StatusOK, rsp)
	})

	log.WithFields(log.Fields{
		"port": c.String("port"),
	}).Info("Starting Raccoon server...")

	e.Run(standard.New(":" + c.String("port")))
}

func parseRequest(r io.Reader) (*raccoon.JobRequest, error) {
	req := raccoon.JobRequest{}

	err := json.NewDecoder(r).Decode(&req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

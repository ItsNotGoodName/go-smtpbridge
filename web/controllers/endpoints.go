package controllers

import (
	"net/http"

	"github.com/ItsNotGoodName/smtpbridge/internal/core"
	"github.com/ItsNotGoodName/smtpbridge/internal/procs"
	"github.com/ItsNotGoodName/smtpbridge/web/helpers"
	"github.com/gofiber/fiber/v2"
)

func Endpoints(c *fiber.Ctx, cc *core.Context) error {
	// Execute
	ends, err := procs.EndpointList(cc)
	if err != nil {
		return helpers.Error(c, err)
	}

	// Response
	return c.Render("endpoints", fiber.Map{
		"Endpoints": ends,
	})
}

func EndpointsTest(c *fiber.Ctx, cc *core.Context, id int64) error {
	// Execute
	err := procs.EndpointTest(cc, id)
	if err != nil {
		return helpers.Error(c, err)
	}

	// Response
	return c.SendStatus(http.StatusNoContent)
}

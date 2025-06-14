package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gxjakkap/dekcpe.link/model"
	"github.com/gxjakkap/dekcpe.link/utils"
)

func (h *Handler) RedirectToHome(c *fiber.Ctx) error {
	return c.Redirect("https://dash.dekcpe.link/", fiber.StatusFound)
}

func (h *Handler) RedirectToLink(c *fiber.Ctx) error {
	slug := c.Params("slug")

	if slug == "favicon.ico" {
		return c.SendStatus(fiber.StatusNoContent)
	}

	link, err := h.linkStore.GetLinkBySlug(slug)

	if err != nil || link == nil {
		return c.Status(fiber.StatusNotFound).Render("not-found", fiber.Map{
			"Slug": slug,
		})
	}

	ip, err := utils.GetIPFromHeaders(c)
	if err != nil {
		log.Printf("failed to get ip from req")
	}

	go func(ip, ua, utmSource string, linkID int) {
		log.Printf("ip: %s, link: %d", ip, linkID)
		var geo *model.GeoLocation
		if geo, err = utils.GetGeoFromIP(ip); err != nil {
			geo = &model.GeoLocation{
				Country:  "Unknown",
				Region:   "Unknown",
				City:     "Unknown",
				Timezone: "Unknown",
			}
		}

		click := &model.Click{
			LinkID:      linkID,
			UserAgent:   ua,
			UTMSource:   &utmSource,
			Geolocation: *geo,
		}

		if err := h.clicksStore.Create(click); err != nil {
			log.Printf("Failed to create click: %v", err)
		}
	}(ip, c.Get("User-Agent"), c.Query("utm_source"), link.ID)
	return c.Redirect(link.URL, fiber.StatusFound)
}

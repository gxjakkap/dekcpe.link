package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Register(r *fiber.App) {
	r.Get("/", h.RedirectToHome)
	r.Get("/:slug", h.RedirectToLink)
}

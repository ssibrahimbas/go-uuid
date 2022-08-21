package internal

import fiber "github.com/gofiber/fiber/v2"

type Handler struct {
	srv *Service
	app *fiber.App
}

func NewHandler(srv *Service, app *fiber.App) *Handler {
	return &Handler{
		srv: srv,
		app: app,
	}
}

func (h *Handler) Init() {
	h.initV1()
}

func (h *Handler) initV1() {
	v1 := h.app.Group("users/v1")
	v1.Post("/create", h.createUser)
	v1.Get("/detail/:id", h.getUserDetail)
	v1.Get("/all", h.getAllUsers)
}

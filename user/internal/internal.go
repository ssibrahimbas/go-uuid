package internal

import (
	pg "github.com/go-pg/pg/v10"
	fiber "github.com/gofiber/fiber/v2"
)

type UserModule struct {
	handler *Handler
}

func New(db *pg.DB, app *fiber.App) *UserModule {
	r := NewRepository(db)
	srv := NewService(r)
	h := NewHandler(srv, app)
	return &UserModule{
		handler: h,
	}
}

func (um *UserModule) Init() {
	um.handler.Init()
}

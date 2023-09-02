package handlers

import "github.com/gofiber/fiber/v2"

type CallbackHandlers struct{}

func NewCallbackHandlers() *CallbackHandlers {
	return &CallbackHandlers{}
}

func (h *CallbackHandlers) Callback(ctx *fiber.Ctx) error {
	return nil
}

package rest

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func (api *Service) CreateTask(ctx *fiber.Ctx) error {
	id, err := api.app.CreateTask(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create task",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func (api *Service) GetTaskStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	task, err := api.app.GetTaskStatus(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	resp := fiber.Map{
		"id":         task.ID,
		"status":     task.Status,
		"created_at": task.CreatedAt.Format(time.RFC3339),
		"duration":   mapDuration(task.StartedAt, task.CompletedAt),
		"result":     task.Result,
	}
	return ctx.JSON(resp)
}

func (api *Service) DeleteTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := api.app.DeleteTask(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

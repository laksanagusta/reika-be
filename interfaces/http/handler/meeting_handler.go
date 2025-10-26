package handler

import (
	"sandbox/application/dto"
	"sandbox/application/usecase"

	"github.com/gofiber/fiber/v2"
)

type MeetingHandler struct {
	createMeetingUseCase *usecase.CreateMeetingUseCase
}

func NewMeetingHandler(createMeetingUseCase *usecase.CreateMeetingUseCase) *MeetingHandler {
	return &MeetingHandler{
		createMeetingUseCase: createMeetingUseCase,
	}
}

func (h *MeetingHandler) CreateMeeting(c *fiber.Ctx) error {
	var reqBody dto.CreateMeetingRequest

	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body format",
			"error":   err.Error(),
		})
	}

	// Validate request using DTO validation method
	if err := reqBody.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	// Get context from fiber
	ctx := c.Context()

	response, err := h.createMeetingUseCase.Execute(ctx, reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	if !response.Success {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

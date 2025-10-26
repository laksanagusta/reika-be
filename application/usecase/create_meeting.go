package usecase

import (
	"context"
	"fmt"

	"sandbox/application/dto"
	"sandbox/domain/meeting"
)

type CreateMeetingUseCase struct {
	meetingService *meeting.Service
}

func NewCreateMeetingUseCase(meetingService *meeting.Service) *CreateMeetingUseCase {
	return &CreateMeetingUseCase{
		meetingService: meetingService,
	}
}

func (uc *CreateMeetingUseCase) Execute(ctx context.Context, req dto.CreateMeetingRequest) (*dto.CreateMeetingResponse, error) {
	// Convert to domain (validation is now handled in handler layer)
	meetingEntity, err := req.ToDomain()
	if err != nil {
		return &dto.CreateMeetingResponse{
			Success: false,
			Message: fmt.Sprintf("Invalid request format: %v", err),
		}, nil
	}

	// Create meeting
	result, err := uc.meetingService.CreateMeeting(ctx, *meetingEntity)
	if err != nil {
		return &dto.CreateMeetingResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to create meeting: %v", err),
		}, nil
	}

	responseData := &dto.MeetingResponseData{
		Meeting:          result.Meeting,
		DriveFolderURL:   result.DriveFolderURL,
		AbsenceFormURL:   result.AbsenceFormURL,
		NotificationSent: result.NotificationSent,
	}

	return &dto.CreateMeetingResponse{
		Success: true,
		Message: "Meeting created successfully",
		Data:    responseData,
	}, nil
}

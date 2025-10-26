package meeting

import (
	"context"
	"fmt"
	"time"

	"sandbox/domain/meeting"
	"sandbox/infrastructure/drive"
	"sandbox/infrastructure/notification"
	"sandbox/infrastructure/zoom"
)

type Repository struct {
	zoomClient         *zoom.Client
	driveClient        *drive.Client
	notificationClient *notification.Client
}

func NewRepository(
	zoomClient *zoom.Client,
	driveClient *drive.Client,
	notificationClient *notification.Client,
) meeting.Repository {
	return &Repository{
		zoomClient:         zoomClient,
		driveClient:        driveClient,
		notificationClient: notificationClient,
	}
}

func (r *Repository) CreateZoomMeeting(ctx context.Context, meeting meeting.Meeting) (*meeting.Meeting, error) {
	return r.zoomClient.CreateZoomMeeting(ctx, meeting)
}

func (r *Repository) CreateDriveFolder(ctx context.Context, parentFolderID, folderName string) (string, error) {
	return r.driveClient.CreateFolder(ctx, parentFolderID, folderName)
}

func (r *Repository) DuplicateAbsenceForm(ctx context.Context, templateID, folderID string) (string, error) {
	if templateID == "" {
		return "", fmt.Errorf("template ID is required")
	}

	// Generate a unique name for the new form
	newFileName := fmt.Sprintf("Absence Form - %s", time.Now().Format("2006-01-02-15-04"))

	return r.driveClient.DuplicateFile(ctx, templateID, folderID, newFileName)
}

func (r *Repository) SendNotification(ctx context.Context, opts meeting.NotificationOpts, meetingURL string) error {
	return r.notificationClient.SendNotification(ctx, opts, meetingURL)
}
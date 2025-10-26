package meeting

import "context"

type Repository interface {
	CreateZoomMeeting(ctx context.Context, meeting Meeting) (*Meeting, error)
	CreateDriveFolder(ctx context.Context, parentFolderID, folderName string) (string, error)
	DuplicateAbsenceForm(ctx context.Context, templateID, folderID string) (string, error)
	SendNotification(ctx context.Context, opts NotificationOpts, meetingURL string) error
}
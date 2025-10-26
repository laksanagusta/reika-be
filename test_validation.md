# Testing Validation with DTO Validation Methods

## Arsitektur Validasi Baru

- **DTO Layer:** Berisi fungsi `Validate()` dengan spesifikasi validasi lengkap
- **Handler Layer:** Memanggil `req.Validate()` dari DTO
- **UseCase Layer:** Business logic murni tanpa validasi input
- **Domain Layer:** Hanya business entities dan core business logic

## Meeting Validation Test

### Request with Invalid Data:
```bash
curl -X POST http://localhost:3000/api/meetings \
  -H "Content-Type: application/json" \
  -d '{
    "title": "",
    "description": "Test meeting with invalid title",
    "start_time": "2024-01-01T10:00:00",
    "timezone": "UTC",
    "duration_minutes": 500,
    "host_user_id": "user123",
    "options": {
      "duplicate_absence_form": true,
      "absence_form_template_id": "",
      "notify": {
        "send_email": true,
        "channels": [],
        "message": ""
      },
      "zoom": {
        "auto_recording": "invalid"
      }
    },
    "metadata": {
      "tags": ["tag_that_is_too_long_to_be_valid_and_should_fail_validation"]
    }
  }'
```

### Expected Response:
```json
{
  "success": false,
  "message": "Validation failed",
  "error": "Title: cannot be blank; Duration: must be no more than 480; absence_form_template_id: absence_form_template_id is required when duplicate_absence_form is true; channels: channels are required when send_email is true; message: message is required when send_email is true; auto_recording: auto_recording must be one of: none, local, cloud; Tags: all values must be at most 50 characters"
}
```

## Transaction Validation Test

### Request with Invalid Data:
```bash
curl -X POST http://localhost:3000/api/transactions/generate-excel \
  -H "Content-Type: application/json" \
  -d '{
    "startDate": "32 Oktober 2025",
    "endDate": "30 Februari 2025",
    "activityPurpose": "",
    "destinationCity": "",
    "spdDate": "tanggal salah",
    "departureDate": "25 Desember 2025",
    "returnDate": "24 Desember 2025",
    "assignees": []
  }'
```

### Expected Response:
```json
{
  "error": "Validation failed",
  "details": "StartDate: does not match pattern '^\\d{1,2}\\s+(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember)\\s+\\d{4}$'; EndDate: does not match pattern '^\\d{1,2}\\s+(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember)\\s+\\d{4}$'; ActivityPurpose: cannot be blank; DestinationCity: cannot be blank; SpdDate: does not match pattern '^\\d{1,2}\\s+(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember)\\s+\\d{4}$'; assignees: at least one assignee is required; date_range: start date must be before or equal to end date; travel_dates: departure date must be before or equal to return date"
}
```

## Valid Request Test

### Meeting Creation with Valid Data:
```bash
curl -X POST http://localhost:3000/api/meetings \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Valid Meeting Title",
    "description": "This is a valid meeting description",
    "start_time": "25 Oktober 2025 10:00:00",
    "timezone": "Asia/Jakarta",
    "duration_minutes": 60,
    "host_user_id": "host123",
    "options": {
      "create_drive_folder": true,
      "duplicate_absence_form": false,
      "notify": {
        "send_email": false
      },
      "zoom": {
        "waiting_room": true,
        "require_password": true,
        "auto_recording": "cloud"
      }
    },
    "metadata": {
      "project_id": "proj123",
      "tags": ["important", "client-meeting"]
    }
  }'
```

### Transaction Report with Valid Indonesian Date Format:
```bash
curl -X POST http://localhost:3000/api/transactions/generate-excel \
  -H "Content-Type: application/json" \
  -d '{
    "startDate": "1 November 2025",
    "endDate": "5 November 2025",
    "activityPurpose": "Meeting Client di Surabaya",
    "destinationCity": "Surabaya",
    "spdDate": "1 November 2025",
    "departureDate": "2 November 2025",
    "returnDate": "4 November 2025",
    "receiptSignatureDate": "5 November 2025",
    "assignees": [
      {
        "name": "Budi Santoso",
        "spdNumber": "SPD-001/2025",
        "employee_id": "EMP001",
        "position": "Software Engineer",
        "rank": "Junior",
        "transactions": [
          {
            "name": "Hotel Surabaya",
            "type": "accommodation",
            "amount": 500000,
            "subtotal": 500000,
            "payment_type": "transfer",
            "description": "Malam 1-3 November"
          }
        ]
      }
    ]
  }'
```

Both requests should pass validation and proceed to the business logic layer.
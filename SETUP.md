# Setup Guide

## Panduan Setup Project Domain-Driven Design

### Prerequisites

- **Go** versi 1.25 atau lebih tinggi
- **Google Gemini API Key** (dapatkan dari [Google AI Studio](https://makersuite.google.com/app/apikey))
- **Git** (untuk version control)

### Langkah-langkah Setup

#### 1. Clone Project

```bash
git clone <repository-url>
cd sandbox
```

#### 2. Install Dependencies

```bash
go mod download
go mod tidy
```

#### 3. Setup Environment Variables

Buat file `.env` di root project:

```bash
# Buat file .env
touch .env
```

Isi file `.env` dengan konfigurasi berikut:

```env
# Server Configuration
PORT=5002

# Gemini API Configuration
# PENTING: Ganti dengan API key Anda sendiri!
GEMINI_API_KEY=your_actual_api_key_here

# CORS Configuration
CORS_ALLOW_ORIGINS=http://localhost:3000
```

**Cara mendapatkan GEMINI_API_KEY:**

1. Kunjungi [Google AI Studio](https://makersuite.google.com/app/apikey)
2. Login dengan akun Google
3. Klik "Create API Key"
4. Copy API key dan paste ke file `.env`

#### 4. Verify Build

```bash
# Build project
go build -o sandbox

# Atau langsung run
go run main.go
```

Jika berhasil, Anda akan melihat output:

```
🚀 Server running on port 5002
📝 Environment: development
```

#### 5. Test API

Buka terminal baru dan test endpoint:

```bash
# Health check
curl http://localhost:5002/api/health
```

Response yang diharapkan:

```json
{
  "status": "healthy"
}
```

### Testing Upload Endpoint

#### Menggunakan cURL

```bash
curl -X POST http://localhost:5002/api/upload \
  -F "file=@/path/to/your/receipt.jpg" \
  -F "file=@/path/to/another/document.pdf"
```

#### Menggunakan Postman

1. Method: `POST`
2. URL: `http://localhost:5002/api/upload`
3. Body:
   - Type: `form-data`
   - Key: `file` (type: File)
   - Value: Select your image/PDF file
   - (Bisa tambahkan multiple files dengan key yang sama)

### Struktur Project

```
sandbox/
├── domain/                     # Layer Domain
│   ├── transaction/
│   │   ├── entity.go          # Entity Transaction
│   │   ├── repository.go      # Interface Repository
│   │   └── service.go         # Domain Service
│   └── errors/
│       └── errors.go          # Domain Errors
│
├── application/               # Layer Aplikasi
│   ├── usecase/
│   │   └── extract_transactions.go
│   └── dto/
│       └── transaction_dto.go
│
├── infrastructure/            # Layer Infrastruktur
│   ├── gemini/
│   │   └── client.go
│   └── file/
│       └── processor.go
│
├── interfaces/               # Layer Interface
│   └── http/
│       ├── handler/
│       ├── middleware/
│       └── router/
│
├── config/                   # Konfigurasi & DI
│   ├── config.go
│   └── container.go
│
├── main.go                   # Entry point
├── go.mod                    # Dependencies
├── .env                      # Environment variables (buat sendiri)
├── .gitignore               # Git ignore rules
├── README.md                # Dokumentasi utama
├── ARCHITECTURE.md          # Dokumentasi arsitektur
└── SETUP.md                 # Panduan setup (file ini)
```

### Troubleshooting

#### Error: "GEMINI_API_KEY environment variable is required"

**Solusi**:

- Pastikan file `.env` sudah dibuat
- Pastikan `GEMINI_API_KEY` sudah diisi dengan API key yang valid
- Restart aplikasi setelah mengubah `.env`

#### Error: "failed to call Gemini API"

**Solusi**:

- Check koneksi internet
- Verify API key masih valid
- Check quota API key di Google AI Studio

#### Port sudah digunakan

**Solusi**:

```bash
# Ganti port di .env
PORT=8080
```

#### Build error: "package not found"

**Solusi**:

```bash
# Clean dan reinstall dependencies
go clean
go mod tidy
go mod download
```

### Development Workflow

#### 1. Menjalankan Development Server

```bash
# Dengan auto reload (install air terlebih dahulu)
go install github.com/cosmtrek/air@latest
air

# Atau manual
go run main.go
```

#### 2. Running Tests

```bash
# Run all tests
go test ./...

# Run tests dengan coverage
go test -cover ./...

# Run tests untuk specific package
go test ./domain/transaction/...
```

#### 3. Build untuk Production

```bash
# Build binary
go build -o sandbox

# Build dengan optimizations
go build -ldflags="-s -w" -o sandbox

# Run binary
./sandbox
```

### Environment Variables Reference

| Variable             | Deskripsi             | Default               | Required |
| -------------------- | --------------------- | --------------------- | -------- |
| `PORT`               | Port server HTTP      | 5002                  | No       |
| `GEMINI_API_KEY`     | Google Gemini API key | -                     | **Yes**  |
| `CORS_ALLOW_ORIGINS` | CORS allowed origins  | http://localhost:3000 | No       |

### Production Deployment

#### Docker (Optional)

Buat `Dockerfile`:

```dockerfile
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o sandbox

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/sandbox .
COPY .env .env

EXPOSE 5002
CMD ["./sandbox"]
```

Build dan run:

```bash
docker build -t sandbox-app .
docker run -p 5002:5002 --env-file .env sandbox-app
```

### Security Checklist

- [ ] ✅ `.env` file tidak di-commit ke Git
- [ ] ✅ API key disimpan di environment variables
- [ ] ✅ CORS dikonfigurasi dengan proper origins
- [ ] ✅ File upload size limits diterapkan
- [ ] ✅ Input validation di semua endpoints
- [ ] ✅ Error messages tidak mengekspos informasi sensitif

### Next Steps

1. Baca [README.md](README.md) untuk overview project
2. Baca [ARCHITECTURE.md](ARCHITECTURE.md) untuk memahami arsitektur
3. Explore code di setiap layer
4. Mulai development! 🚀

### Support

Jika mengalami masalah:

1. Check dokumentasi di README.md dan ARCHITECTURE.md
2. Verify environment variables sudah benar
3. Check logs untuk error messages
4. Ensure semua dependencies ter-install dengan baik

Happy coding! 🎉

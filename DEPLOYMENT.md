# Reika Backend - Docker Deployment Guide

## Build & Push Docker Image

### Prerequisites
- Docker installed
- Access to container registry (Docker Hub, AWS ECR, GCR, etc.)
- `.env` file with required environment variables

### Quick Start (Recommended)

1. **Build Docker Image**
   ```bash
   # Build with latest tag
   ./build-and-push.sh latest your-registry.com

   # Build with version tag
   ./build-and-push.sh v1.0.0 docker.io/username
   ```

2. **Login to Registry**
   ```bash
   # Docker Hub
   docker login

   # Private registry
   docker login your-registry.com
   ```

3. **Push Image**
   ```bash
   # Push to registry
   ./push.sh latest your-registry.com
   ```

### Manual Commands

#### Build Only
```bash
docker build -t reika-backend:latest .
docker tag reika-backend:latest your-registry.com/reika-backend:latest
```

#### Push Only
```bash
docker push your-registry.com/reika-backend:latest
```

#### Run Locally
```bash
docker run -p 5002:5002 --env-file .env your-registry.com/reika-backend:latest
```

## Cloud Deployment

### Option 1: Using Docker (Linux Server)

1. **Install Docker on server:**
```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
sudo usermod -aG docker $USER
```

2. **Pull and run:**
```bash
docker pull your-dockerhub-username/reika:latest
docker run -d \
  --name reika-app \
  -p 5002:5002 \
  --restart unless-stopped \
  your-dockerhub-username/reika:latest
```

### Option 2: Using Docker Compose (Production)

Create `docker-compose.prod.yml`:

```yaml
version: '3.8'

services:
  app:
    image: your-dockerhub-username/reika:latest
    ports:
      - "5002:5002"
    environment:
      - GIN_MODE=release
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:5002/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s

  # Optional: Nginx reverse proxy
  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro
      - ./ssl:/etc/nginx/ssl:ro
    depends_on:
      - app
    restart: unless-stopped
```

Run with:
```bash
docker-compose -f docker-compose.prod.yml up -d
```

### Option 3: Kubernetes Deployment

Create `k8s-deployment.yaml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: reika-app
  labels:
    app: reika
spec:
  replicas: 3
  selector:
    matchLabels:
      app: reika
  template:
    metadata:
      labels:
        app: reika
    spec:
      containers:
      - name: reika
        image: your-dockerhub-username/reika:latest
        ports:
        - containerPort: 5002
        env:
        - name: GIN_MODE
          value: "release"
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 5002
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 5002
          initialDelaySeconds: 5
          periodSeconds: 5

---
apiVersion: v1
kind: Service
metadata:
  name: reika-service
spec:
  selector:
    app: reika
  ports:
  - protocol: TCP
    port: 80
    targetPort: 5002
  type: LoadBalancer
```

Deploy with:
```bash
kubectl apply -f k8s-deployment.yaml
```

## Environment Variables

### Required Environment Variables

Create `.env` file:
```env
GEMINI_API_KEY=your-gemini-api-key
PORT=5002
CORS_ALLOW_ORIGINS=http://localhost:3000
```

| Variable | Default | Description |
|----------|---------|-------------|
| `GEMINI_API_KEY` | Required | Google Gemini API key for transaction extraction |
| `PORT` | `5002` | Server port |
| `CORS_ALLOW_ORIGINS` | `http://localhost:3000` | Allowed CORS origins |

## Monitoring & Health Checks

### Health Check Endpoint

- **URL:** `/api/health`
- **Method:** GET
- **Response:** `{"status": "healthy"}`

### Container Health Check

The Docker image includes built-in health checks:

```bash
# Check container health
docker ps

# View health check logs
docker inspect reika-app | grep Health -A 20
```

## CI/CD Integration

### GitHub Actions Example

Create `.github/workflows/docker.yml`:

```yaml
name: Build and Push Docker Image

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - run: go mod download
      - run: go test -v ./...

  build-and-push:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    steps:
      - uses: actions/checkout@v3
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2
      - name: Login to DockerHub
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}
      - name: Build and push
        uses: docker/build-push-action@v4
        with:
          context: .
          push: true
          tags: |
            your-dockerhub-username/reika:latest
            your-dockerhub-username/reika:${{ github.sha }}
```

## Troubleshooting

### Container Won't Start

```bash
# Check logs
docker logs reika-app

# Check if port is available
netstat -tulpn | grep :5002

# Run with different port
docker run -d -p 8080:5002 your-dockerhub-username/reika:latest
```

### Health Check Failing

```bash
# Test health endpoint manually
curl http://localhost:5002/health

# Check if application is running
docker exec -it reika-app ps aux
```

### Build Issues

```bash
# Clean build cache
docker builder prune -a

# Rebuild without cache
docker build --no-cache -t your-dockerhub-username/reika:latest .
```

## Security Considerations

1. **Non-root user:** The container runs as non-root user (`appuser`)
2. **Minimal base image:** Uses Alpine Linux for smaller attack surface
3. **Read-only filesystem:** Consider adding `--read-only` flag in production
4. **Resource limits:** Set memory and CPU limits in production
5. **Network security:** Use firewall rules to restrict access

## Scaling

### Horizontal Scaling with Docker Compose

```yaml
# In docker-compose.yml
services:
  app:
    image: your-dockerhub-username/reika:latest
    deploy:
      replicas: 3
    ports:
      - "5002-5004:5002"
```

### Load Balancing

Use nginx, HAProxy, or cloud load balancer to distribute traffic across multiple instances.

## Support

For deployment issues, check:

1. Container logs: `docker logs reika-app`
2. Health status: `docker ps`
3. Network connectivity: `curl http://localhost:5002/health`
4. Resource usage: `docker stats reika-app`
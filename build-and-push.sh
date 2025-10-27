#!/bin/bash

# Build and Push Docker Image Script
# Usage: ./build-and-push.sh [version] [registry]

set -e

# Default values
VERSION=${1:-latest}
REGISTRY=${2:-"your-registry.com"}
IMAGE_NAME="reika-backend"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}🚀 Building Docker image for Reika Backend${NC}"
echo -e "${YELLOW}Version: ${VERSION}${NC}"
echo -e "${YELLOW}Registry: ${REGISTRY}${NC}"
echo ""

# Build the Docker image
echo -e "${GREEN}📦 Building image...${NC}"
docker build -t ${IMAGE_NAME}:${VERSION} .
docker tag ${IMAGE_NAME}:${VERSION} ${REGISTRY}/${IMAGE_NAME}:${VERSION}

if [ "$VERSION" != "latest" ]; then
    docker tag ${IMAGE_NAME}:${VERSION} ${REGISTRY}/${IMAGE_NAME}:latest
fi

echo -e "${GREEN}✅ Build completed successfully!${NC}"
echo ""

# Show image info
echo -e "${YELLOW}📋 Image Information:${NC}"
docker images | grep ${IMAGE_NAME}

echo ""
echo -e "${GREEN}🎯 To push to registry, run:${NC}"
echo "docker push ${REGISTRY}/${IMAGE_NAME}:${VERSION}"
if [ "$VERSION" != "latest" ]; then
    echo "docker push ${REGISTRY}/${IMAGE_NAME}:latest"
fi

echo ""
echo -e "${GREEN}🐳 To run locally:${NC}"
echo "docker run -p 5002:5002 --env-file .env ${IMAGE_NAME}:${VERSION}"
#!/bin/bash

# Push Docker Image to Registry
# Usage: ./push.sh [version] [registry]

set -e

VERSION=${1:-latest}
REGISTRY=${2:-"your-registry.com"}
IMAGE_NAME="reika-backend"

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}📤 Pushing Docker image to registry${NC}"
echo -e "${YELLOW}Image: ${REGISTRY}/${IMAGE_NAME}:${VERSION}${NC}"
echo ""

# Check if logged in to registry
echo -e "${YELLOW}🔐 Checking registry authentication...${NC}"
if ! docker info | grep -q "Username"; then
    echo -e "${RED}❌ Not logged in to Docker registry${NC}"
    echo -e "${YELLOW}Please run: docker login ${REGISTRY}${NC}"
    exit 1
fi

# Push the image
echo -e "${GREEN}🚀 Pushing image...${NC}"
docker push ${REGISTRY}/${IMAGE_NAME}:${VERSION}

if [ "$VERSION" != "latest" ]; then
    echo -e "${GREEN}🚀 Pushing latest tag...${NC}"
    docker push ${REGISTRY}/${IMAGE_NAME}:latest
fi

echo -e "${GREEN}✅ Image pushed successfully!${NC}"
echo ""
echo -e "${YELLOW}📋 Deploy commands:${NC}"
echo "docker pull ${REGISTRY}/${IMAGE_NAME}:${VERSION}"
echo "docker run -p 5002:5002 --env-file .env ${REGISTRY}/${IMAGE_NAME}:${VERSION}"
# Build Stage for Backend
FROM golang:1.23.2-alpine3.20 AS backend-builder

WORKDIR /app

# Copy Go modules separately to leverage caching
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the entire backend source code
COPY backend/ .

# Build the Go application
RUN go build -o main main.go

# Build Stage for Frontend (Using Node.js v20.11.0)
FROM node:20.11.0 AS frontend-builder

WORKDIR /frontend

# Copy package.json and install dependencies
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

# Copy frontend source code and build
COPY frontend/ .
RUN npm run build

# Run Stage for Backend
FROM alpine:3.14.2 AS backend

WORKDIR /app

# Install required dependencies (e.g., for SSL, networking, debugging)
RUN apk --no-cache add ca-certificates

# Copy the backend binary
COPY --from=backend-builder /app/main .

# Copy environment file
COPY backend/app.env .

# Expose backend port
EXPOSE 8080

# Start backend
CMD ["./main"]

# Run Stage for Frontend with Nginx 
FROM nginx:alpine AS frontend

WORKDIR /usr/share/nginx/html

# Remove default Nginx static files
RUN rm -rf ./*

# Copy frontend build output
COPY --from=frontend-builder /frontend/dist ./

# Expose frontend port
EXPOSE 3000

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]

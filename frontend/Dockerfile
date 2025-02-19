# Build Stage
FROM node:20.11.0 AS builder
WORKDIR /app
COPY . .
RUN npm install && npm run build

# Run Stage
FROM nginx:alpine
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /app/dist ./
EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]
# Build.
FROM golang:latest AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./app

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage /app ./app
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["./app/app"]

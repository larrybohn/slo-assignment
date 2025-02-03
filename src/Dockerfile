FROM golang:1.23 AS build

WORKDIR /app

COPY . ./

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /store

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build /store /store
COPY --from=build /app/index.html /index.html

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/store"]
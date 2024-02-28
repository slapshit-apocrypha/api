FROM golang:1.22.0 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o api

FROM gcr.io/distroless/base-debian12:nonroot
COPY --from=builder /app/api /api

ENTRYPOINT [ "/api" ]
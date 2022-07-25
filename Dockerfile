FROM golang:1.18-bullseye as builder

WORKDIR /app

COPY ./ ./

RUN go mod download -x

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /wicked cmd/wicked/main.go


FROM gcr.io/distroless/static

COPY --from=builder /wicked /

CMD ["/wicked"]

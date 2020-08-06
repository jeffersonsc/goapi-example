FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

ENV TZ=America/Fortaleza

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app cmd/api/main.go

FROM alpine

ENV PORT=3333

COPY --from=builder /go/bin/app /go/bin/app

EXPOSE 3333

ENTRYPOINT ["/go/bin/app"]
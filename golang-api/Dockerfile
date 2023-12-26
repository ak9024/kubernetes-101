FROM golang:1.21 as builder
WORKDIR /
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o web

FROM alpine
RUN apk add curl
COPY --from=builder /web .
ENTRYPOINT [ "/web" ]


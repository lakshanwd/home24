# build stage
FROM golang:alpine
RUN apk add --no-cache git build-base
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o home24
EXPOSE 3000
ENTRYPOINT /go/src/app/home24
LABEL name=home24-web-page-analyzer
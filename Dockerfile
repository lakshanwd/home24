# build stage
FROM golang:alpine as build
RUN apk add --no-cache git build-base
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o home24

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /go/src/app/home24 /app
COPY --from=build /go/src/app/static /static
ENTRYPOINT [ "/app" ]
LABEL name=home24
EXPOSE 3000
FROM golang:1.16-alpine AS build

WORKDIR /go/src

COPY ./go/src /go/src

RUN go build -ldflags "-w" -o /go/app/go-hello-codeeducation .


FROM alpine:3.5
COPY --from=build /go/app/go-hello-codeeducation /app/go-hello-codeeducation

CMD ["/app/go-hello-codeeducation"]

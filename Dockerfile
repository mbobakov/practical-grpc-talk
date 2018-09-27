
FROM golang:1.11.0-alpine as builder
WORKDIR /project
ENV GO111MODULE=on
RUN apk add --no-cache git gcc musl-dev
COPY --chown=root:root . /project
RUN go build -o time-machine cmd/time-machine/*.go

FROM alpine:3.8
EXPOSE 6060
EXPOSE 50051
COPY --from=builder /project/time-machine /server
CMD [ "/time-machine" ]

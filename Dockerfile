
FROM golang:1.11.0-alpine as builder
WORKDIR /project
ENV GO111MODULE=on
RUN apk add --no-cache git gcc musl-dev
COPY --chown=root:root . /project
RUN go build -o client cmd/client/*.go
RUN go build -o server cmd/server/*.go

FROM alpine:3.8
EXPOSE 8080
EXPOSE 6060
EXPOSE 50051
COPY --from=builder /project/server /server
COPY --from=builder /project/client /client
CMD [ "/server" ]

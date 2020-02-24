FROM golang:1.13-alpine as build
RUN apk add git
RUN go get github.com/githubnemo/CompileDaemon
ADD . /app
WORKDIR /app
RUN go build -o dependent-server .

FROM alpine as app
COPY --from=build /app/dependent-server /app/
CMD ["/app/dependent-server"]

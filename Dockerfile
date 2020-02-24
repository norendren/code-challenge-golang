FROM golang:1.13-alpine as build
RUN apk add git
RUN go get github.com/securego/gosec/cmd/gosec/...
RUN go get github.com/githubnemo/CompileDaemon
ADD . /app
WORKDIR /app
RUN go build -o code-challenge-golang .

FROM alpine as app
COPY --from=build /app/code-challenge-golang /app/
WORKDIR /app
CMD ["./code-challenge-golang"]

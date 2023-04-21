FROM golang:1.20.3-alpine3.16 AS build

WORKDIR /go/src/github.com/rizqo46/go-todo-list

COPY . .

RUN CGO_ENABLED=0 go build -mod=vendor -ldflags="-w -s" -o /binary

FROM scratch

COPY --from=build /binary /binary

EXPOSE 3030

ENTRYPOINT ["/binary"]
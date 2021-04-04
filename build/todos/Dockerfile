FROM golang:1.16-alpine AS build

WORKDIR /go/src/todos
COPY . ./

RUN go build \
    -o /out/service \
    ./cmd/todos/main.go

FROM alpine
RUN adduser --system appuser
WORKDIR /app
COPY --from=build /out/service /app/service
USER appuser
CMD ["/app/service"]
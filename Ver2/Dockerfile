FROM golang:1-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE 8080
CMD [ "./main" ]
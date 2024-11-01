FROM golang:1.22-bullseye as local
RUN apt-get update && apt-get upgrade -y
WORKDIR /app
COPY . .
RUN go install github.com/air-verse/air@v1.52.3
CMD ["air", "-c", ".air.toml"]

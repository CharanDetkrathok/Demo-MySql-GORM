FROM golang:1.19.3-bullseye 
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o main .
COPY config/config.yaml /config/
EXPOSE 3001
CMD [ "./main" ]

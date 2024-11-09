FROM golang:1.22.5
WORKDIR /online_chat
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /online-chat
EXPOSE 8080
CMD ["/online-chat"]
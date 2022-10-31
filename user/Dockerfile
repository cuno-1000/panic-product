FROM golang:1.16.10-alpine3.14
# WORKDIR /app
RUN export GO111MODULE=on
COPY ./go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o /app
CMD [ "/app" ]
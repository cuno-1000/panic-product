FROM alpine
ADD api /api
EXPOSE 8080
ENTRYPOINT [ "/api" ]

# FROM golang:1.16.10-alpine3.14
# RUN export GO111MODULE=on
# COPY ./go.mod go.sum ./
# RUN go mod download && go mod verify
# COPY . .
# RUN go build -o /app
# EXPOSE 8080
# CMD [ "/app" ]
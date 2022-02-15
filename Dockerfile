FROM golang:1.16.14-alpine3.15
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
RUN go build -o /chaostheory-task

EXPOSE 8080

CMD [ "/chaostheory-task" ]


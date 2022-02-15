FROM golang:1.16.14-alpine3.15
RUN apk add --no-cache gcc musl-dev

WORKDIR $GOPATH/src
COPY . .

RUN go mod download

RUN go build -o /chaostheory-task
RUN go build -o /go/bin/chaostheory-task

EXPOSE 3000

CMD [ "./chaostheory-task" ]


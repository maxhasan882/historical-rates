FROM golang:alpine as build-env

RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN mkdir /historical_rates
WORKDIR /historical_rates/
ENV HOME .
ADD . /historical_rates
COPY go.mod go.sum /historical_rates/

RUN go mod download

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app cmd/app/main.go
COPY .env /go/bin/
FROM scratch
COPY --from=build-env /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]
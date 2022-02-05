FROM golang:1.17 as build-env

WORKDIR /go/src/app
COPY *.go .
COPY .env .

RUN go mod init
RUN go get -d -v ./...
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static

COPY --from=build-env /go/bin/app /
COPY --from=build-env /go/src/app/.env /
# port 23023 should be opened
EXPOSE 23023
CMD ["/app"]

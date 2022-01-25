FROM golang:1.17

WORKDIR /go/src/app
COPY . .

# run go list ./... to get a sense for './...'
# it seems like it matches all of the *.go files within a module.
# it could also be that it matches only the main packages.
RUN go get -d -v ./...
RUN go install -v ./...

# my app would produce 'backend' and 'dinny' as executables.
# I choose to only run the backend as dinny is more of a cli application
CMD ["backend"]
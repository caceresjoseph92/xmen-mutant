FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/xmen-mutant
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/xmen-mutant/cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/xmen-mutant /go/bin/xmen-mutant
ENTRYPOINT ["/go/bin/xmen-mutant"]

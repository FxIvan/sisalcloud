FROM golang:alpine AS build

ENV GOPROXY=https://proxy.golang.org
ENV GOPATH=/go

WORKDIR /go/src/servergo
#quiero que copie mi directorio excercise/twoParameter dentro de la GOPATH del mi contenedor

COPY ./excercise/twoParameter.go /usr/local/go/src/excercise/twoParameter.go
COPY . . 

RUN GOOS=linux go build -o /go/bin/servergo server.go

FROM alpine
COPY --from=build /go/bin/servergo /go/bin/servergo
ENTRYPOINT ["/go/bin/servergo"]
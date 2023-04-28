FROM golang:alpine AS build

ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/excelgo
COPY . .
COPY cmd/myapp/ . 

RUN GOOS=linux go build -o /go/bin/excelgo server.go

FROM alpine
COPY --from=build /go/bin/excelgo /go/bin/excelgo
ENTRYPOINT ["/go/bin/excelgo"]
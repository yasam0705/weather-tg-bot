FROM golang:1.18 as builder

WORKDIR /app
ADD . /app

COPY . ./

RUN make build-linux

FROM alpine:latest

COPY --from=builder app/bin/app ./app

RUN chmod +x ./app 

CMD ["./app"]

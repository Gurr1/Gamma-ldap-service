FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN go install -v

ENV ADDRESS http://0.0.0.0:8081/api
ENV API_KEY api_key

CMD ["app"]
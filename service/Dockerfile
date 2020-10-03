FROM golang

WORKDIR /go/src/github.com/chimano/water-those-service
COPY . .
RUN go get github.com/chimano/water-those-service
RUN go build github.com/chimano/water-those-service

CMD ["water-those-service"]
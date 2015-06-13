FROM golang:1.4.2

COPY . /go/src/app
WORKDIR /go/src/app

RUN make dependency
RUN make install

CMD ["app"]


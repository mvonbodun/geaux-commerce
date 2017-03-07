FROM golang:1.7.4-onbuild

COPY . /go/src/app
RUN go get -u -d -v
RUN go install -v

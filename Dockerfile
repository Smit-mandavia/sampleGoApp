#FROM golang:latest
#
#COPY go.mod
#RUN go mod download
#ADD ./* /goapp
#EXPOSE 8080
#RUN go build -o goapp /goapp/bin/app
#CMD ["/goapp/bin/app"]
#

#FROM  golang:1.8.3-alpine3.6
FROM golang:latest
#
#ARG GOWS=/home/go
#ARG GOPKG=sampleApp
#ENV GOBIN=${GOWS}/bin
#ENV GOPATH=${GOWS}
#
#RUN mkdir -p ${GOBIN}
#
## Add and build Go source
#ADD . ${GOWS}/src/${GOPKG}/
#ADD go.mod .
#RUN go mod download
#RUN ls
#RUN go build -o ${GOBIN}/sampleApp ${GOPKG}
#
## Set the default CMD to run io-apiserver
#CMD GOBIN/sampleApp


ADD . /go/src/sampleApp
WORKDIR /go/src/sampleApp
RUN go get sampleApp
RUN go install
ENTRYPOINT ["/go/bin/sampleApp"]
FROM golang:1.11 as mustache
RUN go get github.com/cbroglie/mustache/...

FROM centos:7
COPY --from=mustache /go/bin/mustache /usr/bin/mustache
COPY ./bin/oc /usr/bin/oc
COPY ./build.sh /usr/bin/build
ENTRYPOINT ["/usr/bin/build"]

FROM golang:1.11 as build
COPY . src/app
RUN ls / && pwd
RUN go build -o /build app/cmd

FROM centos:7
COPY ./bin/oc /usr/bin/oc
COPY --from=build ./build /usr/bin/build
ENTRYPOINT ["/usr/bin/build"]

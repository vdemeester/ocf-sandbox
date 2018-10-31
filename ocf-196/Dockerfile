FROM centos:7
RUN yum -y install git golang && yum clean all
COPY . /src/app
ENV GOPATH=/
RUN go build -o /usr/bin/build app/cmd
COPY ./bin/oc /usr/bin/oc
COPY ./spec.mustache /spec.mustache
ENTRYPOINT ["/usr/bin/build"]

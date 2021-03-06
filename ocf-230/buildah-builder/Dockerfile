FROM centos:7 as base
RUN \
  # Install buildah dependencies.
  yum -y install \
    make \
    golang \
    bats \
    btrfs-progs-devel \
    device-mapper-devel \
    glib2-devel \
    gpgme-devel \
    libassuan-devel \
    libseccomp-devel \
    ostree-devel \
    git \
    bzip2

FROM base as runc
ARG RUNC_REVISION="master"

RUN yum -y install runc
RUN mkdir ~/runc && \
  cd ~/runc && \
  export GOPATH=`pwd` && \
  git clone https://github.com/opencontainers/runc ./src/github.com/opencontainers/runc && \
  cd $GOPATH/src/github.com/opencontainers/runc && \
  git checkout "${RUNC_REVISION}" && \
  make runc && \
  mv runc /usr/bin/runc

FROM base as buildah
ARG BUILDAH_REVISION="dockerfile-leading-args"
RUN yum -y install \
    go-md2man \
    runc \
    skopeo-containers
RUN mkdir ~/buildah && \
  cd ~/buildah && \
  export GOPATH=`pwd` && \
  cd $GOPATH/ && \
  git clone https://github.com/vdemeester/buildah ./src/github.com/containers/buildah && \
  cd $GOPATH/src/github.com/containers/buildah && \
  git checkout "${BUILDAH_REVISION}" && \
  make && \
  make install

FROM centos:7
RUN yum -y install libarchive ostree lzo libseccomp libedit
COPY --from=runc /usr/bin/runc /usr/bin/runc
COPY --from=buildah /usr/local/bin/buildah /usr/bin/buildah
COPY --from=buildah /etc/containers /etc/containers
COPY push.sh /usr/bin/push
RUN chmod +x /usr/bin/push
ENV BUILDAH_ISOLATION chroot
ENV STORAGE_DRIVER vfs


FROM golang:1.12.3 as builder
LABEL app="rivet" by="ennoo"
ENV REPO=$GOPATH/src/github.com/ennoo/rivet
WORKDIR $REPO
RUN git clone https://github.com/ennoo/rivet.git ../rivet && \
 go build -o $REPO/bow/runner/bow $REPO/bow/runner/bow.go
FROM centos:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/ennoo/rivet/bow/runner/bow .
EXPOSE 19219
CMD ./bow

# Build-time metadata as defined at http://label-schema.org
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL io.github.ennoo.build-date=$BUILD_DATE \
      io.github.ennoo.name="Rivet Bow Image" \
      io.github.ennoo.description="Gateway developed with golang" \
      io.github.ennoo.url="https://github.com/ennoo/rivet" \
      io.github.ennoo.license="Apache License 2.0" \
      io.github.ennoo.docker.dockerfile="/bow/runner/Dockerfile" \
      io.github.ennoo.vcs-ref=$VCS_REF \
      io.github.ennoo.vcs-type="Git" \
      io.github.ennoo.vcs-url="https://github.com/ennoo/rivet.git" \
      io.github.ennoo.vendor="ENNOO" \
      io.github.ennoo.version=$VERSION
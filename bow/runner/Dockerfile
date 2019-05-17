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
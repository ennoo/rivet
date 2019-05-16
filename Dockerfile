FROM golang:1.12.3 as builder
LABEL app="rivet" by="ennoo"
ENV REPO=$GOPATH/src/github.com/ennoo/rivet
WORKDIR $REPO
RUN go get github.com/ennoo/rivet
RUN CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $REPO/rivet/bow/runner/bow_darwin_amd64 $REPO/rivet/bow/runner

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder $REPO/rivet/bow/runner/bow_darwin_amd64 .
EXPOSE 19219
CMD ["./bow_darwin_amd64"]
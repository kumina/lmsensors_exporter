#!/bin/sh

docker run -i -v `pwd`:/lmsensors_exporter alpine:3.8 /bin/sh << 'EOF'
set -ex

# Install prerequisites for the build process.
apk update
apk add ca-certificates git go libc-dev
update-ca-certificates

# Build the lmsensors_exporter.
cd /lmsensors_exporter/cmd/lmsensors_exporter/
export GOPATH=/gopath
go get -d ./...
go build --ldflags '-extldflags "-static"'
strip lmsensors_exporter
mv /lmsensors_exporter/cmd/lmsensors_exporter/lmsensors_exporter /lmsensors_exporter/lmsensors_exporter
EOF

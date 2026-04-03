FROM docker.io/library/golang:1.25
RUN apt-get update && apt-get install -y osslsigncode openssl bash

WORKDIR /output
WORKDIR /build

# copy module files and download deps first for better caching
COPY go.mod go.sum* ./

# copy source (including tests) and then resolve/check dependencies
COPY . .
RUN go mod tidy
RUN go test ./... -v

# build the binaries
RUN GOOS=windows GOARCH=amd64 go build -o /build/rudder-multitool-unsigned.exe .
RUN GOOS=linux GOARCH=amd64 go build -o /output/rudder-multitool-linux .
RUN GOOS=darwin GOARCH=amd64 go build -o /output/rudder-multitool-mac .

# sign the windows version and place signed binary into /output
ARG SIGNING_PASSWORD
RUN osslsigncode sign \
    -pkcs12 ruddervirt.pfx \
    -pass "${SIGNING_PASSWORD}" \
    -n "Ruddervirt" \
    -i "https://ruddervirt.com" \
    -in /build/rudder-multitool-unsigned.exe \
    -out /output/rudder-multitool.exe && \
    rm /build/rudder-multitool-unsigned.exe /build/ruddervirt.pfx
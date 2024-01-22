############################
# STEP 1 build executable binary
############################
FROM golang:1.21-alpine AS build

# Install git + SSL ca certificates.
## Git is required for fetching the dependencies.
## Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Set docker environment
ARG env
ARG GITLAB_USER
ARG GITLAB_PASSWORD

# For get private gitlab repository
RUN echo "machine gitlab.com login ${GITLAB_USER} password ${GITLAB_PASSWORD}" > ~/.netrc && chmod 600 ~/.netrc

# Set Golang Env
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    env=${env} \
    GOPRIVATE=gitlab.com

WORKDIR /go-app

COPY . .

# Fetch dependencies.
## Using go get.
## RUN go get -d -v

## Using go mod.
RUN go mod download
RUN go mod verify

# Build the binary.
RUN go build -a -ldflags="-w -s" -o goapp .

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /go-app

# Copy our static executable
COPY --from=build /go-app .

# Open Port
EXPOSE 8080

# Run the binary.
ENTRYPOINT ["./goapp"]

ARG  BUILDER_IMAGE=golang
FROM ${BUILDER_IMAGE} AS builder

# Set destination for COPY
WORKDIR $GOPATH/src/app/

# Download Go modules
COPY go.mod ./
RUN go mod download

COPY . ./

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -o /go/bin/app .

FROM scratch
EXPOSE 8080

# Copy our static executable
COPY --from=builder /go/bin/app /go/bin/app

# Run the app binary.
CMD ["/go/bin/app"]
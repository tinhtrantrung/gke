############################
# STEP 1 build executable binary
############################
# golang alpine 1.11.5
FROM golang@sha256:8dea7186cf96e6072c23bcbac842d140fe0186758bcc215acb1745f584984857 as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

# Fetch dependencies.

# Using go get.
RUN go get -d -v

# Using go mod.
# RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/myapp .

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable
COPY --from=builder /go/bin/myapp /go/bin/myapp

# Use an unprivileged user.
USER appuser

# Run the myapp binary.
ENTRYPOINT ["/go/bin/myapp"]
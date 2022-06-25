############################
# STEP 1 build executable binary
############################
FROM golang:1.18-alpine as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Configure git to use token to fetch private dependencies
ARG github_token=empty
RUN git config --global url."https://$github_token:@github.com/".insteadOf "https://github.com/"

# Create appuser
ENV USER=appuser
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}" $USER

WORKDIR /workdir

COPY go.mod go.sum ./

RUN go mod download

COPY . .

#
# Build the binary
RUN GOOS=linux GOARCH=386 go build -ldflags="-w -s" -o /go/bin/app ./cmd/gqlgen_example/

############################
# STEP 2 build a small image
############################

FROM scratch

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/bin/app /go/bin/app

# Use an unprivileged user.
USER appuser:appuser

# Run the app binary.
ENTRYPOINT ["/go/bin/app"]

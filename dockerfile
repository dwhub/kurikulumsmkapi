# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.11

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/dwhub/kurikulumsmkapi

# Copy the local package files to the container's workspace.
# ADD . /go/src/github.com/dwhub/kurikulumsmkapi

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...
#RUN go install github.com/dwhub/kurikulumsmkapi

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/kurikulumsmkapi

# Document that the service listens on port 7000.
EXPOSE 7000
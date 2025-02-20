# Use the official Go image as a base image
FROM golang:1.23

# Install required packages: git, make, zsh, and curl
RUN apt-get update && apt-get install -y \
  git \
  make \
  zsh \
  curl \
  tmux \
  openjdk-17-jre-headless \
  && rm -rf /var/lib/apt/lists/*

# Install Delve, a Go debugging tool
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Install Reflex, a watch task tool
RUN go install github.com/cespare/reflex@latest

# Install cobra-cli for scaffolding CLI applications
RUN go install github.com/spf13/cobra-cli@latest

# Install gotestsum, an alternate testing tool with prettier output
RUN go install gotest.tools/gotestsum@latest

# Install testing library
RUN go install gotest.tools/gotestsum@latest

# Install Oh My Zsh non-interactively
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# Install OpenAPI Generator CLI:
# 1. Download the JAR file (using version 6.6.0 in this example; update as needed).
# 2. Create a wrapper script to run the generator via 'java -jar'.
RUN curl -L https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.6.0/openapi-generator-cli-6.6.0.jar -o /usr/local/bin/openapi-generator-cli.jar && \
  chmod +x /usr/local/bin/openapi-generator-cli.jar && \
  echo '#!/bin/sh\njava -jar /usr/local/bin/openapi-generator-cli.jar "$@"' > /usr/local/bin/openapi-generator && \
  chmod +x /usr/local/bin/openapi-generator

# Set the working directory inside the container
WORKDIR /workspace
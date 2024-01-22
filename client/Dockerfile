FROM golang:latest AS build-env


# Add the private key arg
ARG SSH_PRIVATE_KEY
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

# Make the root foler for our ssh 
RUN mkdir -p /root/.ssh && \
chmod 0700 /root/.ssh && \
ssh-keyscan github.com > /root/.ssh/known_hosts

# Define a work directory
WORKDIR /go/src
RUN mkdir -p /go/bin
ENV GOBIN = /go/bin

# Copy our files
COPY . .

# put the contents of the private key argument into id_rsa inside our ssh folder and give permissions
# Then download our modules and private modules, finally delete the ssh key.
RUN echo "$SSH_PRIVATE_KEY" > /root/.ssh/id_rsa && \
chmod 600 /root/.ssh/id_rsa && \
go mod download && \
rm -rf /root/.ssh/

# Build our go binary
RUN go build -o client


EXPOSE 8000

ENTRYPOINT ["/go/src/client"]
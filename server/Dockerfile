FROM golang:latest as builder

ENV GOPATH="/RUBICON-CHATROOM"
ENV GOBIN="/RUBICON-CHATROOM/bin"
ENV PATH=$PATH:$GOBIN
ENV GOPRIVATE="github.com/NicNedwob/RUBICON-CHATROOM"

RUN mkdir -p /RUBICON-SERVER
RUN mkdir /RUBICON-SERVER/src
RUN mkdir /RUBICON-SERVER/bin

WORKDIR /RUBICON-SERVER


COPY go.mod ./
RUN go mod download

# Copy dir
COPY main.go ./src
COPY internal /src

# RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
RUN CGO_ENABLED=0 GOOS=linux go build src/main.go 
#-o /RUBICON-SERVER/bin/server


EXPOSE 8080

CMD ["/bin/server"]
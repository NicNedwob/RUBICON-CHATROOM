FROM golang:1.26 as builder

ENV GOPATH="/RUBICON-CHATROOM"
ENV GOBIN="/RUBICON-CHATROOM/bin"
ENV PATH=$PATH:$GOBIN
ENV GOPRIVATE="github.com/NicNedwob/RUBICON-CHATROOM"

RUN mkdir -p /RUBICON-CLIENT

WORKDIR /RUBICON-CLIENT
RUN mkdir /src
RUN mkdir /bin

COPY go.mod ./
RUN go mod download

# Copy cmd dir
COPY main.go ./src
COPY internal /src

# RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
RUN CGO_ENABLED=0 GOOS=linux go build src/main.go -o /bin/client


EXPOSE 8080

CMD ["/bin/client"]
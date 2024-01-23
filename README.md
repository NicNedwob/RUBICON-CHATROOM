# RUBICON TCP Chatroom
This repository contains two dockerised Go Modules. The first module is for a TCP server that registers new users as clients to the server and broadcasts a message from any registered client to all other registered clients. The server module is run on the desired server machine. You need this machine's IPv4 address. Use ifconfig (linux) or ipconfig (windows) to find the W-LAN Adapter Wifi IPv4 Address. 

The second module is the client module. It connects to the server using the server's IPv4 address and fascilitates a CLI for welcoming a new client to the chat, setting that client's username and writing messages to the server. 

Each module has a docker file to containerise the module. The docker files allow the Go source code to be built into executables within a docker image. Running the docker image creates a container to run the compiled Go code. The server and client modules will run and connect provided Docker is running and the correct server ip address is provided. The server and client can also run on the same machine using multiple terminals as shown below:

<p float="center">
  <img src="docs/resources/Working Example.png" />
</p>

## Setup
First, clone into the repo to get the source code and docker files.   
You might need to provide me with a public SSH key so I can add it to the repo in order for you to build the docker images. 

Clone the repo:
```bash
git clone git@github.com:NicNedwob/RUBICON-CHATROOM.git
```
For the next steps the Docker Engine on your system needs to be running.

## Run Docker Images
The dockerBuildAndRun.md files located in the server and client directories not only detail the docker build command but also the docker run command for each module. But here they are as well for easy access:

### Server
To build, nagivate to server directory and run:
```bash
docker build -t server .
```

Run:
```bash
docker run -it --network=bridge -p <server_host_ipaddrs>:8000:8000 server
```
Change the server_host_ipaddrs as needed. It will be IPv4 address on the current machine.

### Client
To build, nagivate to client directory and run:

```bash
docker build -t client .
```

Run:
```bash
docker run -it --network=host client
```
Runnign the the above command in terminal will open a terminal session giving instructions. At this point enter the IPv4 address of the server (i.e. 192.168.1.171 or whatever was determined by running ipconfig/ifconfig on the server host machine). 


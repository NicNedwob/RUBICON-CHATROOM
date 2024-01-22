# RUBICON TCP Chatroom
Contains two Go Modules: One for client and one for server. Each module has a docker file to containerise the module.
## Setup

Clone the repo:
```bash
git clone git@github.com:NicNedwob/RUBICON-CHATROOM.git
```

You need to provide me a public SSH key so you can build the docker images. 

### Server
While in repository directory:
```bash
cd server/
```
Open the dockerBuildAndRun.md file and run the first command to build the docker image for the server. 

### Client
Likewise, while in the root repository directory:
```bash
cd client/
```
Open the dockerBuildAndRun.md file and run the first command to build the docker image for the client. 

## Usage
Use docker to create the containers for the client and server by using the docker run commands listed in the dockerBuildAndRun.md files for each module. 

# Server Docker Build and Run Commands

## Server Build Command
```bash
docker build -t server .
```

## Server Run Command
```bash
docker run -it --network=bridge -p <server_host_ipaddrs>:8000:8000 server
```
For example:
```bash
docker run -it --network=bridge -p 192.168.1.171:8000:8000 server
```
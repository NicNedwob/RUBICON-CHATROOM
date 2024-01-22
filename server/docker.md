docker build --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" -t server .
docker run -it -p 8080:8080 server
docker save client:latest | gzip > client_latest.tar.gz
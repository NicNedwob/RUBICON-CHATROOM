docker build --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" -t server .
docker run --name rubiserver -p 8080:8080 server
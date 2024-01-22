docker build --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" -t client .
docker run -it -p 8080:8080 client
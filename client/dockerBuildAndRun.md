docker build --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" -t client .
docker run -it client
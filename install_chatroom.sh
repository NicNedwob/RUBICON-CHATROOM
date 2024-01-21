# Setup Git config for SSH
# git config --global url."ssh://git@github.com".insteadOf "https://github.com"
# git config --list --show-origin

# Set GOPRIVATE Environment variable
go env -w GOPRIVATE="github.com/NicNedwob/RUBICON-CHATROOM"

# install server and client executables into ~/go/bin
go install github.com/NicNedwob/RUBICON-CHATROOM/cmd/server@latest
go install github.com/NicNedwob/RUBICON-CHATROOM/cmd/client@latest

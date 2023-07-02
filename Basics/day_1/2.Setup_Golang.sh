#!/bin/bash

# Update the system
sudo dnf update -y

# Install Visual Studio Code
sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc
sudo sh -c 'echo -e "[code]\nname=Visual Studio Code\nbaseurl=https://packages.microsoft.com/yumrepos/vscode\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/vscode.repo'
sudo dnf install code -y

# Install Go
sudo dnf install golang -y

# Set up Go environment variables
echo "export GOPATH=\$HOME/go" >> ~/.bashrc
echo "export PATH=\$PATH:\$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc

# Install the Go extension for Visual Studio Code
code --install-extension ms-vscode.Go

# Install commonly used Go libraries
go get -u github.com/gin-gonic/gin
go get -u github.com/go-sql-driver/mysql
go get -u github.com/golang-jwt/jwt
go get -u github.com/sirupsen/logrus

# Print installation completion message
echo "Installation completed successfully!"

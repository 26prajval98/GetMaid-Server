@echo off
setlocal enabledelayedexpansion
REM get all go packages required
echo "Getting all dependencies please wait....";

echo 0
go get golang.org/x/crypto/

echo 1
go get github.com/go-sql-driver/mysql

echo 2
go get github.com/qor/auth

echo 3
go get golang.org/x/oauth2

echo 4
go get cloud.google.com/go/compute/metadata

echo "All dependencies set";
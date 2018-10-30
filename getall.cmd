@echo off
setlocal enabledelayedexpansion
REM get all go packages required
echo "Getting all dependencies please wait....";

echo 0
go get golang.org/x/crypto/

echo 1
go get github.com/go-sql-driver/mysql

REM echo 2
REM go get github.com/qor/auth

REM echo 3
REM go get golang.org/x/oauth2

REM echo 4
REM go get cloud.google.com/go/compute/metadata

REM echo 5
REM go get golang.org/x/oauth2

REM echo 6
REM go get github.com/dchest/uniuri

echo 7
go get -u github.com/gbrlsnchs/jwt

echo 8
go get github.com/subosito/twilio

echo 9
go get github.com/github.com/rs/cors

echo "All dependencies set";
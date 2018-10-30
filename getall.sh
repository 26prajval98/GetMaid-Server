#!/usr/bin/env bash
#@echo off
#setlocal enabledelayedexpansion
#REM get all go packages required
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

echo 5
go get golang.org/x/oauth2

echo 6
go get github.com/dchest/uniuri

echo 7
go get -u github.com/gbrlsnchs/jwt

echo 8
go get github.com/rs/cors
echo "All dependencies set";

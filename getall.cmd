@echo off
setlocal enabledelayedexpansion
REM get all go packages required
echo "Getting all dependencies please wait....";

go get golang.org/x/crypto/
go get github.com/go-sql-driver/mysql

echo "All dependencies set";
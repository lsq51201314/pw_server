@echo off
cd %cd%
echo ���ڱ���Windowsƽ̨ ������
go build -ldflags "-s -w" main.go
echo ���ڱ���Linuxƽ̨ ������
SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=linux
go build -ldflags "-s -w" main.go
echo ������� ������
pause
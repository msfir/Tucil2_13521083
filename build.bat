@echo off

cd src
go build
cd ..
if not exist bin md bin
move src\pairit.exe bin 1>NUL
set GOPROXY=https://goproxy.cn
echo Building Kernel
set GOOS=windows
set GOARCH=amd64
go version
go build -v -o electron/bnd2.exe -ldflags "-s -w -H=windowsgui"
if "%errorlevel%" == "1" goto :errorend

set GOOS=darwin
set GOARCH=amd64
go build -v -o electron/bnd2 -ldflags "-s -w"
if "%errorlevel%" == "1" goto :errorend

echo Building UI
cd electron
node -v
call npm -v
call npm install && npm run dist
cd ..

:errorend
echo "Error in go build"

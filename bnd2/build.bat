set GOPROXY=https://goproxy.io
echo Building Kernel
set GOOS=windows
set GOARCH=amd64
go version
go build -v -o electron/bnd2.exe -ldflags "-s -w -H=windowsgui"

set GOOS=darwin
set GOARCH=amd64
go build -v -o electron/bnd2 -ldflags "-s -w"

echo Building UI
cd electron
node -v
call npm -v
call npm install && npm run dist
cd ..

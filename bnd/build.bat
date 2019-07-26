set CGO_ENABLED=1
set CC=D:\mingw-w64\i686-8.1.0-posix-dwarf-rt_v6-rev0\mingw32\bin\gcc.exe
set CXX=D:\mingw-w64\i686-8.1.0-posix-dwarf-rt_v6-rev0\mingw32\bin\g++.exe
set GOARCH=386
set PATH=%GOROOT%\bin;D:\mingw-w64\i686-8.1.0-posix-dwarf-rt_v6-rev0\mingw32\bin;%PATH%

go version
go build -v -o bnd.exe -ldflags "-s -w -H=windowsgui"

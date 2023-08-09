# Hashdump

Sliver `hashdump` extension.

## Build

Use the Makefile, artifacts will be stored in `dist/`:

```shell
$ make
rm -rf dist
mkdir -p dist
CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o ./dist/hashdump.x64.dll -buildmode=c-shared ./dll/
CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o ./dist/hashdump.x86.dll -buildmode=c-shared ./dll/
GOOS=windows GOARCH=amd64 go build -o ./dist/hashdump.x64.exe main.go
GOOS=windows GOARCH=386 go build -o ./dist/hashdump.x86.exe main.go
$ tree dist
dist
├── hashdump.x64.dll
├── hashdump.x64.exe
├── hashdump.x64.h
├── hashdump.x86.dll
├── hashdump.x86.exe
└── hashdump.x86.h

1 directory, 6 files
```

The Makefile has a target for each architecture (x86/x64) and builds an executable to test the code outside of Sliver.
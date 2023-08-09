GO ?= go

CC_X64 = x86_64-w64-mingw32-gcc
CXX_X64 = x86_64-w64-mingw32-g++

CC_X86 = i686-w64-mingw32-gcc
CXX_X86 = i686-w64-mingw32-g++

dist_dir = ./dist

# DLLs
DLL_X64 := $(dist_dir)/hashdump.x64.dll
DLL_X86 := $(dist_dir)/hashdump.x86.dll
# Executables (for local testing)
EXE_X64 := $(dist_dir)/hashdump.x64.exe
EXE_X86 := $(dist_dir)/hashdump.x86.exe


all: clean $(dist_dir) dll-x64 dll-x86 exe-x64 exe-x86

$(dist_dir):
	mkdir -p $@

dll-x64:
	CC=$(CC_X64) CXX=$(CXX_X64) CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GO) build -o $(DLL_X64) -buildmode=c-shared ./dll/

dll-x86:
	CC=$(CC_X86) CXX=$(CXX_X86) CGO_ENABLED=1 GOOS=windows GOARCH=386 $(GO) build -o $(DLL_X86) -buildmode=c-shared ./dll/

exe-x64:
	GOOS=windows GOARCH=amd64 $(GO) build -o $(EXE_X64) main.go

exe-x86:
	GOOS=windows GOARCH=386 $(GO) build -o $(EXE_X86) main.go

clean:
	rm -rf dist

.PHONY: all
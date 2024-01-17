# Makefile demo

#伪目标：不会去检查是否存在一个叫做build的文件
.PHONY: build

#define
MF_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
THIS_DIR := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MF_PATH))))
BUILD_DIR := ./build
OUT_NAME := $(CURRENT_DIR)

#默认不开启cgo,需要开启命令行参数CGO=1
CGO := 0
#默认GOOS=linux
OS := linux

ifeq ($(OS),windows)
	OUT_NAME := $(CURRENT_DIR).exe
endif

#项目依赖的静态资源
ETC_DIR := ./etc
STATIC_DIR := ./static
bpath:
	mkdir -p $(BUILD_DIR)
	cp -f $(ETC_DIR) $(BUILD_DIR) -r
	cp -f $(STATIC_DIR) $(BUILD_DIR) -r

build:bpath
	CGO_ENABLED=$(CGO) GOOS=$(OS) GOARCH=amd64 go build -o $(BUILD_DIR)/$(OUT_NAME)

debug:bpath
	CGO_ENABLED=$(CGO) GOOS=$(OS) GOARCH=amd64 go build -tags debug -o $(BUILD_DIR)/$(OUT_NAME)

release:bpath
	CGO_ENABLED=$(CGO) GOOS=$(OS) GOARCH=amd64 go build -ldflags '-s -w --extldflags "-static -fpic"' -o $(BUILD_DIR)/$(OUT_NAME)

current:
	CGO_ENABLED=$(CGO) GOOS=$(OS) GOARCH=amd64 go build -ldflags '-s -w --extldflags "-static -fpic"' -o ./$(OUT_NAME)

chmod: $(BUILD_DIR)/$(OUT_NAME)
	chmod a+x $(BUILD_DIR)/$(OUT_NAME)

clean:
	rm -rf $(BUILD_DIR)/*

env:
	go env

dir:
	@echo $(MF_PATH)
	@echo $(THIS_DIR)
	@echo $(CURRENT_DIR)

protoc:
	 protoc --proto_path=proto --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative proto/*.proto

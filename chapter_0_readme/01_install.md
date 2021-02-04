```
> @Descripttion: GO的安装
> @Author: tacks321@qq.com
> @Date: 2021-02-04 10:08:19
> @LastEditTime: 2021-02-04 10:40:25
```
[toc]


### [1] 设置环境变量 .bashrc

| 环境变量 | 作用  |
| :----: | :----:  |
| GOROOT | `go`的安装目录 自定义环境变量  |
| PATH |  加入PATH，相关文件在文件系统的任何地方都能被调用  |
| GOPATH | `go`  工作目录  |

```shell 
# 打开 vim ~/.bashrc 再最后进行追加
export GOROOT=/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/tacks/go

# 并且生效
source .bashrc
```

 
### [2] 下载go源代码 

- [官网下载地址](https://golang.org/dl/)

[root@Centos7 source]# wget https://storage.googleapis.com/golang/go1.14.6.src.tar.gz
 
### [3] 解压
```
[root@Centos7 source]# tar -zxvf go1.14.6.src.tar.gz  
```

### [4] 移动到GOROOT的位置
```
[root@Centos7 source]#　mv go $GOROOT
```

### [5] 构建 Go
```shell
[root@Centos7 go]# cd $GOROOT/src
[root@Centos7 src]# ./all.bash

[root@Centos7 go]# cd $GOROOT/src
[root@Centos7 src]# ./all.bash
# 报错
./make.bash: line 165: /root/go1.4/bin/go: No such file or directory
Building Go cmd/dist using /root/go1.4. ()
ERROR: Cannot find /root/go1.4/bin/go.
Set $GOROOT_BOOTSTRAP to a working Go tree >= Go 1.4.

```

- 需要先安装 `GOROOT_BOOTSTRAP1.4`

### [6] 安装 `GOROOT_BOOTSTRAP1.4` 
```shell
[root@Centos7 source] wget https://storage.googleapis.com/golang/go1.4-bootstrap-20170531.tar.gz 
[root@Centos7 source] tar -xf go1.4-bootstrap-20170531.tar.gz 
[root@Centos7 source] cd go/src 
[root@Centos7 source/go/src ] ./make.bash

Installed Go for linux/amd64 in /root/source/go
Installed commands in /root/source/go/bin
...

//重命名1.4文件夹，并添加 GOROOT_BOOTSTRAP 环境变量 
[root@Centos7 source] mv go go1.4

# 打开 vim ~/.bashrc 再最后进行追加
export GOROOT_BOOTSTRAP=/root/source/go1.4
# 并且生效
source .bashrc
 
```

### [7] 再次编译go 
```shell
[root@Centos7 go]# cd $GOROOT/src
[root@Centos7 src]# ./all.bash
...
ALL TESTS PASSED
---
Installed Go for linux/amd64 in /go
Installed commands in /go/bin

```

### [8] 检查
```shell
# go的环境
[root@Centos7 ~]# go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/root/.cache/go-build"
GOENV="/root/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/qianxin/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build592703533=/tmp/go-build -gno-record-gcc-switches"

# 版本
[root@Centos7 go]# go version
go version go1.14.6 linux/amd64
```

### [9] 写个demo
```
package main func main() { println("Hello", "world") }

[root@Centos7 go]# go run hello_world.go
hello world!
```

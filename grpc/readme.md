# GRPC 
> 在于微服务架构中，各个服务的代码库之间是相互独立运行的，彼此通信就可以借助`rpc`, 那么`grpc`是谷歌推推行的`rpc`框架，是一套开源成熟的结构数据序列化机制。 `grpc` 可以让客户端程序直接调用另一台服务端应用程序上的方法，主要用作与微服务中接口调用。通过proto3协议缓冲区定义数据格式，定义参数以及返回类型的调用方法，对应服务端实现接口运行grpc服务器，等待客户端的调用。可以跨平台、跨语言、使用`protocol buffers`作为数据交换格式，支持HTTP2，双向传输，认证等。

## 协议缓冲区 Protocol Buffer 

- 官网 (https://developers.google.com/protocol-buffers/docs/proto3)
- 相关资源
    - （http://www.mooyle.com/golang/gRPC/001-proto3.html）
    - （https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/）

### 定义消息类型

- 文件后缀是 `.proto`
- 开头非注释行必须是 `syntax = "proto3";` 指明协议缓冲区的编辑器，默认为 `proto2`。
- 通过 `package` 定义包名，用来区分不同包下的message。（可省略）
- 通过 `import` 可以引入其他的 proto文件。（可省略）
- 定义 `message` 消息类型。
- 同一个 `message的` 每个字段都有唯一一个编号，并且建议终生这个编号都不要改变。
- 通过协议缓冲区编译 `protoc` 命令 使用`--go_out`参数指定生成`go`语言文件，将proto3文件转化为对应的pb.go文件。
  
    ```shell
    syntax = "proto3";
    
    # 定义一个  SearchRequest 消息类型
    message SearchRequest {
      # 格式定义 type fieldName = fieldNumber
      string query = 1;
      int32 page_number = 2;
      int32 result_per_page = 3;
    }
    ```
  
### 相关编译器

- `protoc` grpc通用编译器
  ```gotemplate
  # 下载地址
  https://github.com/protocolbuffers/protobuf/releases
  # 添加到环境变量中（windows / linux）
  # 使用命令 （protoc包以及protoc-gen-go插件的作用）
  # 编译Product.proto之后输出到service文件夹
  protoc --go_out=../service Product.proto
  # 编译
  protoc --go_out=plugins=grpc:../service Product.proto
  ```
- `protoc-gen-go` 编译器插件 go专用的protoc的生成器
  ```gotemplate
    go get github.com/golang/protobuf/protoc-gen-go
  ```

### `go path` & `go mod`

- `go path`

> GO1.1之前都是`GOPATH`模式，进行包依赖管理，项目必须运行在`GOPATH`下，并且需要创建`$GOPATH/src`用来存放 `go get`的下载的包。
```gotemplate
  【1】 安装grpc客户端
        （1）官方推荐，需要翻墙
                go get -u google.golang.org/grpc
        （2）github仓库
            进入GOPATH目录下，src下
            git clone https://github.com/grpc/grpc-go GOPATH/src/google.golang.org/grpc
  【2】 安装protoc-get-go工具
        然后复制到bin下
  
  【3】 安装protobuf
    git clone https://github.com/golang/protobuf.git $GOPATH/src/github.com/golang/protobuf
  【4】 安装genproto
        git clone https://github.com/google/go-genproto.git  $GOPATH/src/google.golang.org/genproto
  【5】 安装
    git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
    git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
    git clone https://github.com/google/go-genproto.git  $GOPATH/src/google.golang.org/genproto
    git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc



  GoCode\grpc\go-path-demo\grpc-client\service\Product.pb.go:15:2: cannot find package "google.golang.org/grpc" in any of:
  D:\Go\src\google.golang.org\grpc (from $GOROOT)
  D:\Vwork\ATacks\GoCode\src\google.golang.org\grpc (from $GOPATH)
  GoCode\grpc\go-path-demo\grpc-client\service\Product.pb.go:16:2: cannot find package "google.golang.org/grpc/codes" in any of:
  D:\Go\src\google.golang.org\grpc\codes (from $GOROOT)
  D:\Vwork\ATacks\GoCode\src\google.golang.org\grpc\codes (from $GOPATH)
  GoCode\grpc\go-path-demo\grpc-client\service\Product.pb.go:17:2: cannot find package "google.golang.org/grpc/status" in any of:
  D:\Go\src\google.golang.org\grpc\status (from $GOROOT)
  D:\Vwork\ATacks\GoCode\src\google.golang.org\grpc\status (from $GOPATH)
  GoCode\grpc\go-path-demo\grpc-client\service\Product.pb.go:18:2: cannot find package "google.golang.org/protobuf/reflect/protoreflect" in any of:
  D:\Go\src\google.golang.org\protobuf\reflect\protoreflect (from $GOROOT)
  D:\Vwork\ATacks\GoCode\src\google.golang.org\protobuf\reflect\protoreflect (from $GOPATH)
  GoCode\grpc\go-path-demo\grpc-client\service\Product.pb.go:19:2: cannot find package "google.golang.org/protobuf/runtime/protoimpl" in any of:
  D:\Go\src\google.golang.org\protobuf\runtime\protoimpl (from $GOROOT)
  D:\Vwork\ATacks\GoCode\src\google.golang.org\protobuf\runtime\protoimpl (from $GOPATH)
```



- `go mod`

> GO1.3之后，推荐使用GOMOD,项目不需要再放到GOPATH下面，每一个项目都是一个Module。

```shell
// 查看环境配置
# go env
// [GO111MODULE] 用来设置是否使用Go mod， 可以设置 on/off/auto
# go env -w GO111MODULE=on
// 配置镜像 就是配置mod的代理。 direct如果代理拿不到，再通过源地址go get获取
# go env -w GOPROXY=https://goproxy.cn,direct

// 恢复初始设置
# go env -u


```


### 程序代码

```gotemplate
- grpc-client
    - service 
        - Product.pb.go [复制服务端生成对应的数据结构]
    - go_client.go [GRPC 客户端主程序]
- grpc-service
    - protofiles
        - Product.proto [编写protobuf文件 定义message]
    - service
        - Product.pb.go [利用protoc工具生成的]
        - ProductService.go [方法封装 实现proto定义的类的方法 也就是实现接口方法]
    - go_service.go [GRPC 服务端主程序]
```
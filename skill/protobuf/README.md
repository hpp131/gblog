1. protoc(protobuf compiler)：protobuf的编译器，可以生成java代码，c++代码，Go代码等。  
[下载并安装protoc](https://github.com/protocolbuffers/protobuf/releases)

protoc命令行参数分析：
- `-I 即--protoc-path`:protoc从这个参数指定的路径下寻找待编译的.proto文件；另外如果带编译的文件中包括import语句，那么import语句后面的文件也会从这个参数指定的路径下面寻找。
```sh
protoc --proto_path=./  --proto_path=/Users/han/practice/go13project/gblog/skill/protobuf/test/common --go_out=./ --go_opt=module="github.com/hpp131/gblog/pb"  hello.proto  common.proto
```




2. protoc-go插件
Protobuf核心的工具集是C++语言开发的，在官方的protoc编译器中并不支持Go语言。要想基于上面的hello.proto文件生成相应的Go代码，需要安装相应的插件
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
执行go install后插件的二进制程序默认安装在GOPATH下，需要把GOPATH添加到系统的PATH变量中，否则执行protoc编译时会报错找不到可执行文件，报错如下：
```sh
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.
```
添加GOPATH到系统PATH中（macos）:
```sh
vim /User/xx/.zprofile
```


3. 关于go_package, go_out, go_opt的用法，推荐官方文档：
[package](https://protobuf.dev/reference/go/go-generated/#package)
- 实测发现.proto文件中的package字段的作用仅仅是防止不同proto文件中message字段的冲突。protobuf中的package并不是Go中package的概念
- 在使用import导入其他proto文件时，import后的.proto文件路径需要在--proto_path指定的路径下能找到！！！


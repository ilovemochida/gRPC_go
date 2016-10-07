
# gRPC-go　環境構築

## gRPCの環境構築

### 注意

GOの環境構築は整っている前提

```
$ cd $GOPATH
$ go get google.golang.org/grpc
```
## protoファイルからgoファイルを生成するために

### 注意　protoDir/test.proto ファイルがある設定

```
$ cd $GOPATH
$ git clone https://github.com/google/protobuf
$ cd protobuf/
$ ./autogen.sh
$ ./configure
$ make
$ make check
$ sudo make install
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
$ export PATH=$PATH:$GOPATH/bin
$ cd /path/to/protoDir/
$ protoc -I ./ ./test.proto --go_out=plugins=grpc:test
```

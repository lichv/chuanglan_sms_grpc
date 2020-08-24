# 创蓝短信 grpc 

### 安装protoc-gen-go

```
go get github.com/golang/protobuf/protoc-gen-go

export PATH="$PATH:$(go env GOPATH)/bin"
```

### 安装 protoc-gen-go-grpc

```
git clone -b v1.31.0 https://github.com/grpc/grpc-go
```

```
cd cmd/protoc-gen-go-grpc && go install .
```

### 生成grpc代码

```shell
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative pb/*.proto
```


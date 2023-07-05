#  libp2p通讯示例

更多[examples](https://github.com/libp2p/go-libp2p/tree/master/examples)

## 运行

在终端执行以下命令

```shell
go run chat.go -sp 6006
```

然后根据提示信息在另一个终端执行对应指令，比如

```shell
go run chat.go -d /ip4/127.0.0.1/tcp/6003/p2p/QmVyDyfBkTSjNuN1J2eVEY5KZCgbY5QL8WvXU14SDcU1MZ
```

接下来就可以在两个终端聊天了
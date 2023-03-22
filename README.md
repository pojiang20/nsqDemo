
#### 安装
安装客户端`go get https://github.com/nsqio/go-nsq`
拉取镜像`docker pull nsqio/nsq`

启动lookupd `docker run --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd`
再启动nsqd 
```docker
docker run --name nsqd -p 4150:4150 -p 4151:4151 \
    nsqio/nsq /nsqd \
    --broadcast-address=<host> \
    --lookupd-tcp-address=<host>:<port>
```

#### 运行结果
可以看到消费者会接收到生产者的消息。
```text
2023/03/22 16:49:36 INF    1 [test/test] (127.0.0.1:4150) connecting to nsqd
2023/03/22 16:49:36 Start consume
2023/03/22 16:49:36 INF    2 [test/lc] (127.0.0.1:4150) connecting to nsqd
2023/03/22 16:49:36 Start consume
{hello}
{hello}
```
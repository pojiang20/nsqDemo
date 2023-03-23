
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

### 2 producer 1 consumer
在`2p1c`中，构建了两个生产者一个消费者，每一个生产者配备一个nsqd。
消费者有两种方式消费消息：
- ConnectToNSQDs，多个地址直连，对应`TestNSQ1`。
- ConnectToNSQLookupd，消费者连接nsqLookupd，对应`TestNSQ2`。
可以看到日志中，消费者的消费情况是不相同的，这是因为nsq会对消费者进行负载均衡，也就是消费者的消费是有策略的。
#### topic与channel
整个模型中，生产者关系的是topic，即将消息`Publish`到哪个`topic`中。而消费者关心的是`channel`,`nsq.NewConsumer("test", "channel1", cfg)`他需要消费哪个`topic`下的`channel`。
#### 组播
而topic到channel的组播则是由nsq完成，并且消息的消费原则是至少被消费一次。
#### lookupd
其作用仅是维护一组活跃的nsqd，并没有路由等其他功能。`ConnectToNSQDs`与`ConnectToNSQLookupd`的区别仅在于后者的hosts是动态的。

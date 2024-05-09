# proxy

A主机————>B主机————>C主机

程序功能主要包括三个：

1、代理转发metrics请求到第三台主机

2、代理转发压缩文件到第三台主机

3、代理转发可执行文件/脚本到第三台主机执行

用法例子：

1、获取metrics信息

go run main.go -m metrics -ip 192.168.145.102:9292

2、发送文件自动压缩为后缀为.gz文件(默认发送到/tmp下)

go run main.go -m upload -s /root/test.tgz -ip 192.168.145.102

3、发送脚本文件到目标端执行(目前只支撑shell)

go run main.go -m script -s /root/test.sh -ip 192.168.145.102

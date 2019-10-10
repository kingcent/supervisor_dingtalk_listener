# supervisor_dingtalk_listener
supervisor process status event listener and send notification to dingtalk group 

监听supervisor的进程状态（运行 停止 异常 退出），发送钉钉通知。通常运用于系统脚本监控

本脚本只监控了进程的部分状态。supervisor相关event见[官方文档](http://supervisord.org/events.html?highlight=tick#event-types)

### 如何使用

1 创建钉钉群组自定义机器人，获取accesstoken 见[官方文档](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)

2 build
```
go build main.go
```
3 添加supervisor event lisenter section 配置
```
    command=/path/to/main xxxxxxaccesstokenxxxxxxx
    process_name=%(program_name)s ; process_name expr (default %(program_name)s)
    ;numprocs=1                    ; number of processes copies to start (def 1)
    events=PROCESS_STATE           ; event notif. types to subscribe to (req'd)
```

4 重新载入配置
```
/usr/bin/supervisorctl update
```


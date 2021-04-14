# grafana webhook to wechat [重构-练手项目]   
  


功能：grafana webhook 通过企业微信的Bot机器人 发送报警通知

#### Go run
<pre>
go build main.go  

[GIN-debug] Listening and serving HTTP on :8081  

</pre>

#### 创建企业微信bot

![](/image/01.png)

例如： https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=`xxxx-xxx-xxx-xxx-xxxx`

#### Alerting Grafana

![](/image/02.png)

#### alerting demo

![](/image/03.png)

#### LOG
<pre>
127.0.0.1 - [Wed, 14 Apr 2021 16:00:21 CST] "POST /eXXX-XXX-XXX-XXX HTTP/1.1 200 37.3801ms "Grafana" "
127.0.0.1 - [Wed, 14 Apr 2021 16:04:12 CST] "POST /eXXX-XXX-XXX-XXX HTTP/1.1 200 28.3171ms "Grafana" "
</pre>

参考项目：https://github.com/n0vad3v/g2ww  





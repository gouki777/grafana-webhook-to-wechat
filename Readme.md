# grafana webhook to wechat [practice]   
  


function：grafana webhook Alarm notification is sent through the BOT robot of enterprise WeChat

#### Go run
<pre>
go build main.go  
</pre>

#### Get webchat-bot key

![](/image/01.png)

https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=`xxxx-xxx-xxx-xxx-xxxx`

#### Alerting Grafana

![](/image/02.png)

#### alerting demo

![](/image/03.png)

#### LOG
<pre>
127.0.0.1 - [Wed, 14 Apr 2021 16:00:21 CST] "POST /eXXX-XXX-XXX-XXX HTTP/1.1 200 37.3801ms "Grafana" "
127.0.0.1 - [Wed, 14 Apr 2021 16:04:12 CST] "POST /eXXX-XXX-XXX-XXX HTTP/1.1 200 28.3171ms "Grafana" "
</pre>

reference：https://github.com/n0vad3v/g2ww  





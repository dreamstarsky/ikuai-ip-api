### ikuai-ip-api


将ikuai网口的ip通过api暴露出来供其他应用使用，解决ikuai内置ddns不好用的问题。



勉强使用，url后+网络接口名称即可以字符串形式返回ip



#### 安装

```sh
git clone https://github.com/dreamstarsky/ikuai-ip-api.git
cd ikuai-ip-api
docker build -t "ikuai-ip-api" .
```

编辑compose.yaml

```yaml
services:
  ikuai-api:
    image: ikuai-ip-api  
    container_name: ikuai-ip-api 
    restart: unless-stopped
    ports:
      - "8080:8080" 
    environment:
      - IKUAI_URL=http://192.168.5.11/  # 请在这里填写你的 iKuai 路由器的地址
      - IKUAI_USERNAME=testuser  # 请在这里填写你的 iKuai 路由器的用户名
      - IKUAI_PASSWORD=114514abc  # 请在这里填写你的 iKuai 路由器的密码
```

这里的用户只要有ikuai中的状态监控-线路监控的访问权限即可

然后

```sh
docker compose up -d
```


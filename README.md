## 启动

```bash
# 启动必要服务
docker-compose up -d mysql redis micro

# 启动业务服务
docker-compose up -d user role api
```

## 接口地址(测试)

`http://localhost:8082/api`

## 后台页面

1. 首先注册一个后台用户

```bash
curl --location --request POST 'localhost:8082/api/v1/user/register' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'email=sfk@live.cn' \
--data-urlencode 'password=123456' \
--data-urlencode 'username=sfk'
```

2. 启动后台
```bash
git clone https://github.com/isfk/aio-web.git
cd aio-web && yarn && yarn dev
```
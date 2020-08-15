# aliddns

自动更新 alidns 记录

# 使用方法

## Docker 运行

```sh
docker run registry.cn-chengdu.aliyuncs.com/quteam/aliddns
```

## 配置文件

```json
{
  "accessKey": "",
  "accessKeySecret": "",
  "domain": "",             // 域名
  "rr": "",                 // 记录
  "interval": 30            // 运行周期（秒）
}
```

## 环境变量配置
环境变量优先级高于配置文件

```sh
export ACCESS_KEY=abc
export ACCESS_KEY_SECRET=123
export DOMAIN=quteam.com    # 域名
export RR=@                 # 记录
export INTERVAL=30          # 运行周期（秒）
```

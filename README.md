# seichi-api
圣地信息api

## 编译打包
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build seichi-api
```
## 部署启动
```
chmod +x seichi-api
nohup ./seichi-api &
```

## 查看是否启动成功
```
netstat -tunlp |grep 8896
```

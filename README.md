# Payment Service

```
micro new order
sudo docker pull micro/micro
sudo docker run --rm -v $(pwd):$(pwd) -w $(pwd) micro/micro new order
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```

## consul key/value设置

micro/config/mysql

{
"host":"0.0.0.0",
"user":"root",
"password":"wq131415",
"database":"micro",
"port":3306
}

#Cart Service
##安装工具protobuf 相关工具
``` 
go get -u github.com/golang/protobuf/protoc-gen-go go-micro
```
##自己的生成工具
```
go get github.com/micro/protoc-gen-micro/v2
```

## docker 安装 consul
```
docker pull consul
docker run -d -p 8500:8500 consul/consul
```

## consul配置和注册安装
```
github.com/micro/go-plugins/config/source/consul/v2
github.com/micro/go-plugins/registry/consul/v2
```

##jaeger镜像安装
```
docker pull jaegertracing/all-in-one
docker run -d -e --name -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:latest
```
```
github.com/uber/jaeger-client-go
github.com/opentracing/opentracing-go
github.com/micro/go-plugins/wrapper/trace/opentracing/v2
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```
## 限流ratelimiter
```
go get github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2
```

## qiankun前端微服务
https://github.com/xushanpei/qiankun_template
module github.com/cuno-1000/panic-product/api-gateway

go 1.16

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.45.0

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/asim/go-micro/v3 v3.5.2-0.20210630062103-c13bb07171bc
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/orcaman/concurrent-map v1.0.0
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

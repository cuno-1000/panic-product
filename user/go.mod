module github.com/cuno-1000/panic-product/user

go 1.16

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.45.0

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210601052333-ca2014bf8e50
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.7.0
	github.com/asim/go-micro/v3 v3.5.2-0.20210630062103-c13bb07171bc
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/micro/v3 v3.2.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/shopspring/decimal v1.3.1
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.27.1
)

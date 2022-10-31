module github.com/cuno-1000/panic-product/repayment_record

go 1.16

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.45.0

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/gorm v1.9.16
	github.com/shopspring/decimal v1.3.1
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.26.0
)

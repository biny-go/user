module user

go 1.24.4

// replace github.com/biny-go/cLoong => ../cLoong // 替换为实际本地路径
replace github.com/biny-go/loong => ../loong // 替换为实际本地路径

require (
	github.com/biny-go/loong v1.1.2
	github.com/go-kratos/kratos/v2 v2.8.4
	github.com/google/wire v0.6.0
	github.com/jinzhu/copier v0.4.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250603155806-513f23925822
	google.golang.org/grpc v1.73.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

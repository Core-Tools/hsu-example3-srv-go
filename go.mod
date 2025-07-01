module github.com/core-tools/hsu-example3-srv-go

go 1.22.3

replace github.com/core-tools/hsu-core => github.com/core-tools/hsu-core/go v0.0.0-20250629192131-f2790593d496

replace github.com/core-tools/hsu-echo => github.com/core-tools/hsu-example3-common/go v0.0.0-20250701175645-62020145072b

replace github.com/core-tools/hsu-echo-simple => .

require (
	github.com/core-tools/hsu-echo v0.0.0-00010101000000-000000000000
	github.com/core-tools/hsu-echo-simple v0.0.0-00010101000000-000000000000
)

require (
	github.com/core-tools/hsu-core v0.0.0-00010101000000-000000000000 // indirect
	github.com/jessevdk/go-flags v1.6.1 // indirect
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

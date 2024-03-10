package main

// go install github.com/hnlq715/struct2interface@master
// go install -v github.com/google/wire/cmd/wire@v0.5.0
//go:generate struct2interface -d internal/pkg/mysql/example
////go:generate struct2interface -d internal/pkg/pgsql/example
//go:generate struct2interface -d internal/pkg/redis/match
//go:generate struct2interface -d internal/pkg/redis/user
////go:generate struct2interface -d internal/pkg/http/example
//go:generate struct2interface -d internal/pkg/grpc/example
//go:generate struct2interface -d internal/pkg/rocketmq/example
//go:generate wire ./...

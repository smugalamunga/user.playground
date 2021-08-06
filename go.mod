module github.com/smugalamunga/user.playground

go 1.16

require (
	github.com/brianvoe/gofakeit/v6 v6.5.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/nats-io/nats.go v1.11.0 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/reactivex/rxgo/v2 v2.5.0 // indirect
	github.com/smugalamunga/instrumentation.playground v0.0.0-20210614205803-9eff369501a8 // indirect
	github.com/smugalamunga/messaging.playground v0.0.0-20210528144138-a0ced6d0fcbe // indirect
	github.com/smugalamunga/playground v0.0.0-20210512110832-eaf2ab01155d
	github.com/smugalamunga/workflow.playground v0.0.0-20210615172537-f209777b5a27 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	go.mongodb.org/mongo-driver v1.5.2
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/grpc v1.38.0

	//github.com/brianvoe/gofakeit/v6 v6.5.0
	//github.com/kr/text v0.2.0 // indirect
	//github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	//github.com/onsi/ginkgo v1.16.4
	//github.com/onsi/gomega v1.13.0
	//github.com/smugalamunga/playground v0.0.0-20210512110832-eaf2ab01155d
	//github.com/stretchr/testify v1.7.0
	//go.mongodb.org/mongo-driver v1.5.2
	//golang.org/x/crypto v0.0.0-20210314154223-e6e6c4f2bb5b // indirect
	//golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	//golang.org/x/net v0.0.0-20210610132358-84b48f89b13b // indirect
	//golang.org/x/sys v0.0.0-20210611083646-a4fc73990273 // indirect
	//golang.org/x/tools v0.1.3 // indirect
	//google.golang.org/grpc v1.38.0
	//gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

replace github.com/smugalamunga/playground => ../playground

replace github.com/smugalamunga/messaging.playground => ../messaging.playground

replace github.com/smugalamunga/workflow.playground => ../workflow.playground

replace github.com/smugalamunga/instrumentation.playground => ../instrumentation.playground

//# git config --global --add url."git@github.com:".insteadOf "https://github.com/"
//# export GOPRIVATE=github.com/smugalamunga/*
//# go get github.com/smugalamunga/playground
//# go mod edit -replace github.com/smugalamunga/playground=../playground
// export GOPRIVATE=github.com/smugalamunga/*
// go get github.com/smugalamunga/workflow.playground
// go mod edit -replace github.com/smugalamunga/workflow.playground=../workflow.playground
//# git remote add origin https://github.com/smugalamunga/user.playground.git
//# go get github.com/golang/protobuf
//# go get google.golang.org/protobuf
//# go get -u google.golang.org/grpc
//# go mod tidy

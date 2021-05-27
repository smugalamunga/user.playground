module github.com/smugalamunga/user.playground

go 1.16

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/smugalamunga/playground v0.0.0-20210512110832-eaf2ab01155d
	github.com/stretchr/testify v1.7.0 // indirect
	go.mongodb.org/mongo-driver v1.5.2
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/sys v0.0.0-20210521203332-0cec03c779c1 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/grpc v1.38.0
)

replace github.com/smugalamunga/playground => ../playground

//# git config --global --add url."git@github.com:".insteadOf "https://github.com/"
//# export GOPRIVATE=github.com/smugalamunga/*
//# go get github.com/smugalamunga/playground
//# go mod edit -replace github.com/smugalamunga/playground=../playground
//# git remote add origin https://github.com/smugalamunga/user.playground.git
//# go get github.com/golang/protobuf
//# go get google.golang.org/protobuf
//# go get -u google.golang.org/grpc
//# go mod tidy

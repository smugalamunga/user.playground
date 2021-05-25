module github.com/smugalamunga/user.playground

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	go.mongodb.org/mongo-driver v1.5.2 // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210521203332-0cec03c779c1 // indirect
	google.golang.org/genproto v0.0.0-20210524171403-669157292da3 // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
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

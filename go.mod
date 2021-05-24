module github.com/smugalamunga/user.playground

go 1.16

require github.com/smugalamunga/playground v0.0.0-20210512110832-eaf2ab01155d // indirect


replace github.com/smugalamunga/playground => ../playground

//# git config --global --add url."git@github.com:".insteadOf "https://github.com/"
//# export GOPRIVATE=github.com/smugalamunga/*
//# go get github.com/smugalamunga/playground
//# go mod edit -replace github.com/smugalamunga/playground=../playground
//# git remote add origin https://github.com/smugalamunga/user.playground.git


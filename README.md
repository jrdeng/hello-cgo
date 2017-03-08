# Hello CGo

leaning cgo

1. get this project

        go get -d github.com/jrdeng/hello-cgo

2. build 3rd-party C code (as dynamically linked library)

        cd $GOPATH/src/github.com/jrdeng/hello-cgo/cwrapper && make

3. using cwrapper in Go

        cd $GOPATH/src/github.com/jrdeng/hello-cgo
        go run main.go

    or

        go install github.com/jrdeng/hello-cgo



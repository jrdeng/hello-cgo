package main

import "github.com/jrdeng/hello-cgo/cwrapper"

func main() {
    str := "Hello, CGO!";
    cwrapper.Println(str);

    cwrapper.MyFoo();
}


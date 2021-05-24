package main

import (
    "fmt"
)

func hello(phrase string) {
    fmt.Println(phrase)
    fmt.Println("new line")
}

func main() {
    var qoses = [3]string{"0", "1", "2"}
    go func(phrase string) {
	for  _,qos := range qoses {
            hello(phrase+qos)
	}
    }("Hello World")
}

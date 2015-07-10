package main

import "fmt"

func Foo(bar int) int {
    if bar > 0 {
        return 123
    } else {
        return 456
    }
}

func main() {
	fmt.Println("Hello, 世界")
}


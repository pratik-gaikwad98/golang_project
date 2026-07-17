package main

import (
	"fmt"
	goroutinesandchannels "my-goworkspace/Programs/goroutines-and-channels"
	interfaceandstruct "my-goworkspace/Programs/interface-and-struct"
)

func main() {
	fmt.Println("Hello, World!")

	interfaceandstruct.InterfaceProgram()

	goroutinesandchannels.EnableGoroutine()
	
}

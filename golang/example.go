package main

import "fmt"

func main() {
	device := NewCPCOE_Device(2, `\\.\COM13`)
	fmt.Println(device.Address)
}

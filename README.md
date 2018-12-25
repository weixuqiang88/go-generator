# easy-go
package main

import (
	"easy"
	"fmt"
)

func main() {
	ip_seg := "192.168.1.0/24"
	ip_gener := easy.NewIPGener(ip_seg)
	ip_iter := easy.NewIterData(ip_gener.IPList)
	for i := 0; i < 10; i++ {
		fmt.Println(ip_iter.Next())
	}
}

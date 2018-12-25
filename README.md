# easy-go
```
a simple go library, some go common functions.

package main

import (
	"github.com/weixuqiang88/easy-go"
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
```

package cmd

import (
	"fmt"
	"time"
)

func Up() {
	fmt.Println("Up")

	cnt := 0
	for {
		fmt.Printf("count is %d\n", cnt)
		cnt++
		time.Sleep(time.Second)
	}
}

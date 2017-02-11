package main

import (
	"fmt"
	"time"
	"github.com/yaman/i2c"
)

func main() {

		i2cbus := i2c.NewI2CBus(1)
	for {
		text, err := i2cbus.ReadBytes(0x60,400)

		if err != nil {
			fmt.Println("iinit errorr: ", err)
		}

		fmt.Println("how many: ", text[0])
		fmt.Println("text: ", text[0:])
		time.Sleep(100 * time.Millisecond)
	}
}


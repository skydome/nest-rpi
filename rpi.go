package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
	"github.com/yaman/i2c"
)

func main() {
	err := rpio.Open()

	if nil != err {
		fmt.Println("error occured :", err)
	}
	pin := rpio.Pin(10)

	i2cbus := i2c.NewI2CBus(1)
	for {
		for pin.Read() == rpio.High {
			text, err := i2cbus.ReadBytes(0x60, 400)

			if err != nil {
				fmt.Println("iinit errorr: ", err)
			}

			fmt.Println("how  many: ", text[0])
			fmt.Println("text: ", text[0:])
			time.Sleep(100 * time.Millisecond)
		}
	}
}

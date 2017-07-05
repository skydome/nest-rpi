package main

import (
	"fmt"
	"github.com/warthog618/gpio"
	"github.com/yaman/i2c"
	"time"
)

func main() {

	i2cbus := i2c.NewI2CBus(1)

	gpioErr := gpio.Open()

	if gpioErr != nil {
		fmt.Println("gpio init error: ", gpioErr)
	}

	defer gpio.Close()

	pin := gpio.NewPin(gpio.J8p7)
	pin.Input()

	pin.Watch(gpio.EdgeRising, func(pin *gpio.Pin) {

		text, err := i2cbus.ReadBytes(0x03, 4)

		if err != nil {
			fmt.Println("i2c read error: ", err)
		}

		fmt.Println("how many: ", text)
	})

	defer pin.Unwatch()

	fmt.Println("Watching pin BCM 4 / PHY 7")

	for {
		time.Sleep(time.Minute)
	}
}

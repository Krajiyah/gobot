// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/Krajiyah/gobot"
	"github.com/Krajiyah/gobot/drivers/gpio"
	"github.com/Krajiyah/gobot/drivers/i2c"
	"github.com/Krajiyah/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	gp := i2c.NewGrovePiDriver(r)
	button := gpio.NewButtonDriver(gp, "D3")
	led := gpio.NewLedDriver(gp, "D2")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pressed")
			led.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button released")
			led.Off()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{gp, button, led},
		work,
	)

	robot.Start()
}

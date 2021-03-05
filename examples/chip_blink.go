// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/Krajiyah/gobot"
	"github.com/Krajiyah/gobot/drivers/gpio"
	"github.com/Krajiyah/gobot/platforms/chip"
)

func main() {
	chipAdaptor := chip.NewAdaptor()
	led := gpio.NewLedDriver(chipAdaptor, "XIO-P6")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{chipAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}

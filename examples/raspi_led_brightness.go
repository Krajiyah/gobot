// +build example
//
// Do not build by default.

package main

import (
	"time"

	"github.com/Krajiyah/gobot"
	"github.com/Krajiyah/gobot/drivers/gpio"
	"github.com/Krajiyah/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "11")

	work := func() {
		brightness := uint8(0)
		fadeAmount := uint8(15)

		gobot.Every(100*time.Millisecond, func() {
			led.Brightness(brightness)
			brightness = brightness + fadeAmount
			if brightness == 0 || brightness == 255 {
				fadeAmount = -fadeAmount
			}
		})
	}

	robot := gobot.NewRobot("pwmBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}

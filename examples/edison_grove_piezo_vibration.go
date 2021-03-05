// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/Krajiyah/gobot"
	"github.com/Krajiyah/gobot/drivers/aio"
	"github.com/Krajiyah/gobot/platforms/intel-iot/edison"
)

func main() {
	board := edison.NewAdaptor()
	sensor := aio.NewGrovePiezoVibrationSensorDriver(board, "0")

	work := func() {
		sensor.On(aio.Vibration, func(data interface{}) {
			fmt.Println("got one!")
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{board},
		[]gobot.Device{sensor},
		work,
	)

	robot.Start()
}

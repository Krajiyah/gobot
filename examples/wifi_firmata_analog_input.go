// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"os"

	"github.com/Krajiyah/gobot"
	"github.com/Krajiyah/gobot/drivers/aio"
	"github.com/Krajiyah/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewTCPAdaptor(os.Args[1])
	sensor := aio.NewAnalogSensorDriver(firmataAdaptor, "A0")

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			brightness := uint8(
				gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1024), 0, 255),
			)
			fmt.Println("sensor", data)
			fmt.Println("brightness", brightness)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor},
		work,
	)

	robot.Start()
}

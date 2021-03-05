// +build example
//
// Do not build by default.

/*
 How to run
 Pass serial port to use as the first param:

	go run examples/firmata_curie_imu_step_counter.go /dev/ttyACM0
*/

package main

import (
	"log"
	"os"
	"time"

	"github.com/Krajiyah/gobot"
	"github.com/Krajiyah/gobot/drivers/gpio"
	"github.com/Krajiyah/gobot/platforms/firmata"
	"github.com/Krajiyah/gobot/platforms/intel-iot/curie"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	led := gpio.NewLedDriver(firmataAdaptor, "13")
	imu := curie.NewIMUDriver(firmataAdaptor)

	work := func() {
		imu.On("Steps", func(data interface{}) {
			log.Println("Steps", data)
		})

		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})

		imu.EnableStepCounter(true)
	}

	robot := gobot.NewRobot("curieBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{imu, led},
		work,
	)

	robot.Start()
}

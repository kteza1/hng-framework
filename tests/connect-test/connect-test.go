package main

import (
	"github.com/kteza1/homeNxtGen-framework/api"
)

/* Start mqtt broker first --> mosquitto -c /usr/local/etc/mosquitto/mosquitto.conf */
func main() {
	homeNxtGen.Connect("Hey")
}

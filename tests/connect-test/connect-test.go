package main

import "github.com/kteza1/hng-framework/api"

/* Start mqtt broker first --> mosquitto -c /usr/local/etc/mosquitto/mosquitto.conf */
func main() {
	hng.Connect("Hey")
}

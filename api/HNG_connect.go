package hng

import (
	"fmt"
	"net/rpc"

	"github.com/kteza1/hng-framework/bus"
	//"github.com/kteza1/go-ninja/rpc/json2"
)

// Connection Holds the connection to the Ninja MQTT bus, and provides all the methods needed to communicate with
// the other modules in Sphere.
type Connection struct {
	mqtt      bus.Bus
	rpc       *rpc.Client
	rpcServer *rpc.Server
}

// Connect Builds a new ninja connection to the MQTT broker, using the given client ID
func Connect(clientID string) (*Connection, error) {

	conn := Connection{}

	mqttURL := fmt.Sprintf("%s:%d", "localhost", 1883)

	conn.mqtt = bus.MustConnect(mqttURL, clientID)

	fmt.Println("MQTT Connection successful")

	//conn.rpc = rpc.NewClient(conn.mqtt, json2.NewClientCodec())
	//conn.rpcServer = rpc.NewServer(conn.mqtt, json2.NewCodec())

	return &conn, nil
}

package bus

import (
	"fmt"
	"strings"
)

type Bus interface {
	Publish(topic string, payload []byte)
	Subscribe(topic string, callback func(topic string, payload []byte)) (*Subscription, error)
	OnDisconnect(cb func())
	OnConnect(cb func())
	Connected() bool
	Destroy()
}

func MustConnect(host, id string) Bus {
	//return ConnectTinyBus(host, id)
	var bus Bus
	var err error

	bus, err = ConnectTinyBus(host, id)

	if err != nil {
		fmt.Println("Unable to connect to MQTT Bus")
	}
	return bus
}

type Subscription struct {
	topic     string
	callback  func(topic string, payload []byte)
	Cancel    func()
	cancelled bool
}

func matches(subscription string, topic string) bool {
	parts := strings.Split(topic, "/")
	subParts := strings.Split(subscription, "/")

	i := 0
	for i < len(parts) {
		// topic is longer, no match
		if i >= len(subParts) {
			return false
		}
		// matched up to here, and now the wildcard says "all others will match"
		if subParts[i] == "#" {
			return true
		}
		// text does not match, and there wasn't a + to excuse it
		if parts[i] != subParts[i] && subParts[i] != "+" {
			return false
		}
		i++
	}

	// make finance/stock/ibm/# match finance/stock/ibm
	if i == len(subParts)-1 && subParts[len(subParts)-1] == "#" {
		return true
	}

	if i == len(subParts) {
		return true
	}
	return false
}

type baseBus struct {
	destroyed          bool
	connectionStatus   bool
	disconnectHandlers []func()
	connectHandlers    []func()
}

func (b *baseBus) OnDisconnect(cb func()) {
	b.disconnectHandlers = append(b.disconnectHandlers, cb)
}

func (b *baseBus) OnConnect(cb func()) {
	b.connectHandlers = append(b.connectHandlers, cb)
}

func (b *baseBus) Connected() bool {
	if b.destroyed {
		return false
	}
	return b.connectionStatus
}

func (b *baseBus) disconnected() {
	if b.destroyed {
		return
	}
	b.connectionStatus = false
	for _, cb := range b.disconnectHandlers {
		go cb()
	}
}

func (b *baseBus) connected() {
	if b.destroyed {
		return
	}

	b.connectionStatus = true
	for _, cb := range b.connectHandlers {
		go cb()
	}
}

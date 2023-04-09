package bridgedevice

import "fmt"

// Abstraction: Device
type Device interface {
	Transmit(message string)
}

// RefinedAbstraction1: Mobile
type Mobile struct {
	communication Communication
}

func (m *Mobile) Transmit(message string) {
	m.communication.SendMessage(message)
}

// RefinedAbstraction2: Laptop
type Laptop struct {
	communication Communication
}

func (l *Laptop) Transmit(message string) {
	l.communication.SendMessage(message)
}

// Implementor: Communication
type Communication interface {
	SendMessage(message string)
}

// ConcreteImplementor1: Bluetooth
type Bluetooth struct{}

func (b *Bluetooth) SendMessage(message string) {
	fmt.Printf("Sending message via Bluetooth: %s\n", message)
}

// ConcreteImplementor2: WiFi
type WiFi struct{}

func (w *WiFi) SendMessage(message string) {
	fmt.Printf("Sending message via WiFi: %s\n", message)
}

func ExampleBridge() {
	mobileBluetooth := &Mobile{communication: &Bluetooth{}}
	mobileWiFi := &Mobile{communication: &WiFi{}}
	laptopBluetooth := &Laptop{communication: &Bluetooth{}}
	laptopWiFi := &Laptop{communication: &WiFi{}}

	mobileBluetooth.Transmit("Hello from mobile via Bluetooth!")
	mobileWiFi.Transmit("Hello from mobile via WiFi!")
	laptopBluetooth.Transmit("Hello from laptop via Bluetooth!")
	laptopWiFi.Transmit("Hello from laptop via WiFi!")
}

package observer

import "fmt"

// Observer is the common interface for all observers (subscribers)
type Observer interface {
	Update(string)
}

// ConcreteObserverA is a concrete observer class
type ConcreteObserverA struct {
	name string
}

func (o *ConcreteObserverA) Update(data string) {
	fmt.Printf("%s received: %s\n", o.name, data)
}

// ConcreteObserverB is another concrete observer class
type ConcreteObserverB struct {
	name string
}

func (o *ConcreteObserverB) Update(data string) {
	fmt.Printf("%s received: %s\n", o.name, data)
}

// Subject is the interface to add, remove, and notify observers
type Subject interface {
	AddObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers()
}

// ConcreteSubject is a concrete subject class
type ConcreteSubject struct {
	observers []Observer
	data      string
}

func (s *ConcreteSubject) AddObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}
func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}
func (s *ConcreteSubject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.data)
	}
}
func (s *ConcreteSubject) SetData(data string) {
	s.data = data
	s.NotifyObservers()
}

func Observe() {
	subject := &ConcreteSubject{}

	observerA := &ConcreteObserverA{name: "Observer A"}
	observerB := &ConcreteObserverB{name: "Observer B"}

	subject.AddObserver(observerA)
	subject.AddObserver(observerB)

	subject.SetData("New Data")
}

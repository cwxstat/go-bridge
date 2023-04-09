package decorator

import "fmt"

// Component is the common interface for the base objects and decorators
type Component interface {
	Operation() string
}

// ConcreteComponent is a concrete base object class
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// Decorator is the abstract decorator class
type Decorator struct {
	component Component
}

func (d *Decorator) SetComponent(c Component) {
	d.component = c
}

func (d *Decorator) Operation() string {
	return d.component.Operation()
}

// ConcreteDecoratorA is a concrete decorator class
type ConcreteDecoratorA struct {
	Decorator
}

func (d *ConcreteDecoratorA) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorA(%s)", d.component.Operation())
}

// ConcreteDecoratorB is another concrete decorator class
type ConcreteDecoratorB struct {
	Decorator
}

func (d *ConcreteDecoratorB) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorB(%s)", d.component.Operation())
}

func ExampleDecorator() {
	component := &ConcreteComponent{}

	decoratorA := &ConcreteDecoratorA{}
	decoratorA.SetComponent(component)

	decoratorB := &ConcreteDecoratorB{}
	decoratorB.SetComponent(decoratorA)

	fmt.Println(component.Operation())
	fmt.Println(decoratorA.Operation())
	fmt.Println(decoratorB.Operation())
}

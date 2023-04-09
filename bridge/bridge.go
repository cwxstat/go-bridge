package bridge

import "fmt"

// Implementor is the interface for the concrete implementors
type Implementor interface {
	OperationImpl() string
}

// ConcreteImplementorA is a concrete implementor class
type ConcreteImplementorA struct{}

func (c *ConcreteImplementorA) OperationImpl() string {
	return "ConcreteImplementorA"
}

// ConcreteImplementorB is another concrete implementor class
type ConcreteImplementorB struct{}

func (c *ConcreteImplementorB) OperationImpl() string {
	return "ConcreteImplementorB"
}

// Abstraction is the abstract class that represents the abstraction
type Abstraction struct {
	implementor Implementor
}

func (a *Abstraction) SetImplementor(implementor Implementor) {
	a.implementor = implementor
}

func (a *Abstraction) Operation() string {
	return a.implementor.OperationImpl()
}

// RefinedAbstraction is a refined abstraction class that extends Abstraction
type RefinedAbstraction struct {
	Abstraction
}

func (r *RefinedAbstraction) Operation() string {
	return fmt.Sprintf("RefinedAbstraction(%s)", r.implementor.OperationImpl())
}

func ExampleBridge() {
	implementorA := &ConcreteImplementorA{}
	implementorB := &ConcreteImplementorB{}

	abstraction := &Abstraction{}
	refinedAbstraction := &RefinedAbstraction{}

	// Using ConcreteImplementorA
	abstraction.SetImplementor(implementorA)
	refinedAbstraction.SetImplementor(implementorA)

	fmt.Println(abstraction.Operation())
	fmt.Println(refinedAbstraction.Operation())

	// Using ConcreteImplementorB
	abstraction.SetImplementor(implementorB)
	refinedAbstraction.SetImplementor(implementorB)

	fmt.Println(abstraction.Operation())
	fmt.Println(refinedAbstraction.Operation())
}

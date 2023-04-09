package factory

import "fmt"

// Product is the common interface for all product types
type Product interface {
	Use() string
}

// ConcreteProductA is a concrete product class
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using ConcreteProductA"
}

// ConcreteProductB is another concrete product class
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using ConcreteProductB"
}

// Factory is the interface for the factory method
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA is a concrete factory class
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteFactoryB is another concrete factory class
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func ExampleFactory() {
	var factoryA Factory = &ConcreteFactoryA{}
	var factoryB Factory = &ConcreteFactoryB{}

	productA := factoryA.CreateProduct()
	productB := factoryB.CreateProduct()

	fmt.Println(productA.Use())
	fmt.Println(productB.Use())
}

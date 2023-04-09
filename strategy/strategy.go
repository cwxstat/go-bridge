package strategy

import "fmt"

// ResourceSelector is the common interface for all resource selection strategies
type ResourceSelector interface {
	SelectResource() string
}

// Strategy1 is a concrete implementation of ResourceSelector
type Strategy1 struct{}

func (s *Strategy1) SelectResource() string {
	return "Resource A"
}

// Strategy2 is a concrete implementation of ResourceSelector
type Strategy2 struct{}

func (s *Strategy2) SelectResource() string {
	return "Resource B"
}

// Context is the class that will use the resource selection strategy
type Context struct {
	strategy ResourceSelector
}

func (c *Context) SetStrategy(strategy ResourceSelector) {
	c.strategy = strategy
}
func (c *Context) ChooseResource() string {
	return c.strategy.SelectResource()
}

func Strat() {
	context := &Context{}

	// Using Strategy1
	context.SetStrategy(&Strategy1{})
	resource := context.ChooseResource()
	fmt.Println("Selected:", resource)

	// Using Strategy2
	context.SetStrategy(&Strategy2{})
	resource = context.ChooseResource()
	fmt.Println("Selected:", resource)
}

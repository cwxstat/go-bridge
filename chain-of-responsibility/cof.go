package cof

import (
	"fmt"
)

// CarRequest represents a request to perform an action on a car
type CarRequest struct {
	Action string
	Car    string
}

// Handler interface represents a handler in the chain of responsibility
type Handler interface {
	HandleRequest(request CarRequest) bool
	SetNext(handler Handler)
}

// WashHandler handles car wash requests
type WashHandler struct {
	next Handler
}

func (w *WashHandler) HandleRequest(request CarRequest) bool {
	if request.Action == "wash" {
		fmt.Printf("Washing car: %s\n", request.Car)
		return true
	}
	if w.next != nil {
		return w.next.HandleRequest(request)
	}
	return false
}

func (w *WashHandler) SetNext(handler Handler) {
	w.next = handler
}

// RepairHandler handles car repair requests
type RepairHandler struct {
	next Handler
}

func (r *RepairHandler) HandleRequest(request CarRequest) bool {
	if request.Action == "repair" {
		fmt.Printf("Repairing car: %s\n", request.Car)
		return true
	}
	if r.next != nil {
		return r.next.HandleRequest(request)
	}
	return false
}

func (r *RepairHandler) SetNext(handler Handler) {
	r.next = handler
}

// PaintHandler handles car paint requests
type PaintHandler struct {
	next Handler
}

func (p *PaintHandler) HandleRequest(request CarRequest) bool {
	if request.Action == "paint" {
		fmt.Printf("Painting car: %s\n", request.Car)
		return true
	}
	if p.next != nil {
		return p.next.HandleRequest(request)
	}
	return false
}

func (p *PaintHandler) SetNext(handler Handler) {
	p.next = handler
}

func ExampleCoF() {
	washHandler := &WashHandler{}
	repairHandler := &RepairHandler{}
	paintHandler := &PaintHandler{}

	// Set up the chain of responsibility
	washHandler.SetNext(repairHandler)
	repairHandler.SetNext(paintHandler)

	requests := []CarRequest{
		{Action: "wash", Car: "Toyota"},
		{Action: "repair", Car: "Honda"},
		{Action: "paint", Car: "Ford"},
	}

	for _, request := range requests {
		if !washHandler.HandleRequest(request) {
			fmt.Printf("No handler for action '%s' on car: %s\n", request.Action, request.Car)
		}
	}
}

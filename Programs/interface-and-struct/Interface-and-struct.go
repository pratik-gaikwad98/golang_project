package interfaceandstruct

import (
	"fmt"
	"math"
)

/////////////////////////////////////////////////////////////////////////////////////

func InterfaceProgram() {

	//Print Rectangle and Cirecle area and paremter
	CalculateAreaAndPerimeter()
	SendNotifier()
}

// //////////////////////////////////////////////////////////////////////////////////
func CalculateAreaAndPerimeter() {
	rect := Rectangle{10, 20}
	crcl := Circle{25}
	PrintShapeDetails(rect)
	PrintShapeDetails(crcl)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Breath float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breath
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Breath + r.Length)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeDetails(s Shape) {
	fmt.Printf("\nArea of %T\n ", s)
	fmt.Printf("Area is %v\n\n", s.Area())

	fmt.Printf("\nPerimeter of %T\n ", s)
	fmt.Printf("Perimeter is %v\n\n", s.Perimeter())
}

/*
	A notification system.

	Notification Service

	1. Email
	2. SMS
	3. Push Notification


	Implements:
	1. EmailNotifier
	2. SMSNotifier
	3. PushNotifier

	Main should work like
	1. Sending Email...
	2. Sending SMS...
	3. Sending Push...
*/

type Notifier interface {
	Send() error
}

type Email struct {
	User string
}

type Mesage struct {
	User string
}

type Notification struct {
	User string
}

func (e Email) Send() error {
	fmt.Printf("Sending Email to %v\n", e.User)
	return nil
}

func (e Mesage) Send() error {
	fmt.Printf("Sending Mesage to %v\n", e.User)
	return nil
}

func (e Notification) Send() error {
	fmt.Printf("Sending Notification to %v\n", e.User)
	return nil
}

func PrintNotifier(n Notifier) {
	fmt.Printf("Notifier is %T and error is %v\n", n, n.Send())
}

func SendNotifier() {
	fmt.Printf("\n----------------------\n")

	emailUser := Email{"gpratik25@gmail.com"}
	PrintNotifier(emailUser)

	messageUser := Mesage{"+917028435555"}
	PrintNotifier(messageUser)

	notifyUser := Notification{"Pratik"}
	PrintNotifier(notifyUser)

}

package interfaceandstruct

import "fmt"

// Interface (Contract)
type PaymentGateway interface {
	Process(amount float64) error
}

// Razorpay Implementation
type Razorpay struct{}

func (r Razorpay) Process(amount float64) error {
	fmt.Printf("Processing payment of ₹%.2f using Razorpay\n", amount)
	return nil
}

// Stripe Implementation
type Stripe struct{}

func (s Stripe) Process(amount float64) error {
	fmt.Printf("Processing payment of ₹%.2f using Stripe\n", amount)
	return nil
}

// Business Logic
type PaymentService struct {
	gateway PaymentGateway
}

func (p PaymentService) Pay(amount float64) error {
	fmt.Println("Starting payment...")
	err := p.gateway.Process(amount)
	if err != nil {
		return err
	}
	fmt.Println("Payment Successful")
	return nil
}

func DoPayment() {

	// Client 1 wants Razorpay
	razorpay := Razorpay{}
	service := PaymentService{
		gateway: razorpay,
	}
	service.Pay(500)

	fmt.Println("----------------")

	// Client 2 wants Stripe
	stripe := Stripe{}
	service = PaymentService{
		gateway: stripe,
	}
	service.Pay(500)
}




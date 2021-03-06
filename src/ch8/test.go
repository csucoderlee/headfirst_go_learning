package main

import "fmt"
import "github.com/csucoderlee/magazine"

func printInfo(s *magazine.Subscriber) {
	fmt.Println(s.Rate)
	fmt.Println(s.Name)
	fmt.Println(s.Active)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var defaultSubscriber magazine.Subscriber
	defaultSubscriber.Name = name
	defaultSubscriber.Rate = 5.99
	defaultSubscriber.Active = true
	return &defaultSubscriber
}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := magazine.Subscriber{Name: "Aman Singh"}
	subscriber2.HomeAddress.Street = "123 Oak Str"

	employee := magazine.Employee{Name: "Aman Singh"}
	//匿名struct可以直接赋值
	employee.Street = "456 Oak Str"

}

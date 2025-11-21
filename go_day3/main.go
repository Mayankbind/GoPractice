package main

import (
	"fmt"
)

type Order struct {
	OrderId int
	Amount  float64
	Customer
}

type Customer struct {
	Name string
	ContactInfo
}

type ContactInfo struct {
	Email string
	Phone string
}

type Notifier interface {
	sendNotification()
}

type EmailService struct{
	SMTPhost string
}

type SmsService struct{
	Provider string
}

func (e EmailService) sendNotification(){
	fmt.Println("Email Notification: ",e.SMTPhost)
}

func (s SmsService) sendNotification(){
	fmt.Println("Sms Notification: ",s.Provider)
}

func Notification(n Notifier, o Order){
	n.sendNotification()
	fmt.Printf("OrderId: %d\nAmount: %.2f\nCustomerName: %s\nContactNo: %s\n",o.OrderId, o.Amount, o.Name, o.Phone)
}

func main() {

	o:=Order{
		OrderId: 1,
		Amount: 2000.321,
		Customer: Customer{
			Name: "Mayank",
			ContactInfo: ContactInfo{
				Phone: "1234567",
			},
		},
	}

	e:=EmailService{
		SMTPhost: "Email Notification",
	}

	s:=SmsService{
		Provider: "SMS",
	}

	Notification(e, o)
	Notification(s, o)

	


}

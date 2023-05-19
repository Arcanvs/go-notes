package main

import "fmt"

// Simulando envio de notificaciones

type INotificationFactory interface {
	sendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotification struct {
}

func (SMSNotification) sendNotification() {
	fmt.Println("Send notification via SMS")
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "CHANNEL"
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotification struct {
}

func (EmailNotification) sendNotification() {
	fmt.Println("Send notification via EMAIL")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "EMAIL" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No notification Type")
}

func sendNotification(f INotificationFactory) {
	f.sendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("EMAIL")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}

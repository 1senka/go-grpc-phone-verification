package sms

import "github.com/kavenegar/kavenegar-go"

//create an instance of client for any sms provider
type Client struct {
	SmsClient *kavenegar.Kavenegar
}

func Init(key string) Client {
	return Client{
		SmsClient: kavenegar.New(key),
	}

}

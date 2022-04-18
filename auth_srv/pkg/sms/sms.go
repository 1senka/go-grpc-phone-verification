package sms

import "github.com/kavenegar/kavenegar-go"

type Client struct {
	SmsClient *kavenegar.Kavenegar
}

func Init(key string) Client {
	return Client{
		SmsClient: kavenegar.New(key),
	}

}

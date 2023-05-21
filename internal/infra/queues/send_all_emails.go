package queues

import (
	"log"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/smtp"
	"github.com/adjust/rmq/v5"
)

func NewSendEmailQueue() rmq.Connection {

	connection, err := rmq.OpenConnection("", "tcp", "localhost:6379", 1, nil)

	if err != nil {
		log.Fatal(err)
	}

	return connection

}

type SendAllEmailsQueueInterface interface {
	Consume(dlv rmq.Delivery)
}

type SendAllEmailsQueue struct {
}

func NewSendAllEmailsQueue() SendAllEmailsQueueInterface {
	return &SendAllEmailsQueue{}
}

func (consumer *SendAllEmailsQueue) Consume(dlv rmq.Delivery) {

	var from = dlv.Payload()

	smtpClient := smtp.NewSMTPClient()

	smtpClient.SendNewProductAdded(from)

	if err := dlv.Ack(); err != nil {
		log.Fatal("aq", err)
	}

}

package queues

import (
	"context"
	"fmt"
	"log"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/configs"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/smtp"
	"github.com/go-redis/redis/v8"
	"github.com/keithwachira/go-taskq"
)

var streamName = "send_order_emails"

type SendAllEmailsQueueInterface interface {
	Process(job interface{})
}

type SendAllEmailsQueue struct {
	Redis *configs.RedisConfigs
}

func NewSendAllEmailsQueue(
	redis *configs.RedisConfigs,
) SendAllEmailsQueueInterface {
	return &SendAllEmailsQueue{
		Redis: redis,
	}
}

func (j *SendAllEmailsQueue) Process(job interface{}) {

	smtpClient := smtp.NewSMTPClient()

	if data, ok := job.(redis.XMessage); ok {
		from := data.Values["email"].(string)
		fmt.Println("Send email to: ", from)

		smtpClient.SendNewProductAdded(from)

	} else {
		log.Println("Invalid job")
	}

}

func StartProcessingEmails(cap int) {

	rdb := configs.NewRedisConfigs()

	queue := NewSendAllEmailsQueue(rdb)

	q := taskq.NewQueue(1, cap, queue.Process)

	go q.StartWorkers()

	id := "0"

	for {

		ctx := context.Background()

		data, err := rdb.Redis.XRead(ctx, &redis.XReadArgs{
			Streams: []string{streamName, id},
			Count:   4,
			Block:   0,
		}).Result()

		if err != nil {
			log.Fatal(err)
		}

		for _, result := range data {
			for _, message := range result.Messages {
				q.EnqueueJobBlocking(message)

				id = message.ID
			}
		}

	}

}

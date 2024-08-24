package delay_queue

import (
	"cms/utils"
	"context"
	"fmt"
	redisv8 "github.com/go-redis/redis/v8"
)

//var DelayQueueClient *DelayQueue

type DelayQueue struct {
	Topic            string  `json:"topic"`
	DelayTime        float64 `json:"delay_time"` // delay time: seconds
	BatchMessageNums int64   `json:"batch_message_nums"`
	SleepSeconds     int64   `json:"sleep_seconds"`

	RedisClient *redisv8.Client
	//Logger      logrus.Logger
}

type DelayQueueMessages struct {
	MessageTime string `json:"message_time"`
	Value       []byte `json:"value"` // message value
}

func InitDelayQueue(ctx context.Context, redisClient *redisv8.Client, cfg *DelayQueueConfig) *DelayQueue {
	if cfg == nil || cfg.DelayTime <= 0 || len(cfg.QueueKey) <= 0 {
		fmt.Printf("delay queue config is empty: [%s]", utils.ConvertToJsonString(cfg))
		panic("delay queue config is empty")
	}

	client := &DelayQueue{}
	if redisClient == nil {
		panic("redis client is nil")
	}
	client.RedisClient = redisClient
	client.DelayTime = cfg.DelayTime
	client.Topic = cfg.QueueKey
	client.BatchMessageNums = cfg.BatchMessageNums
	if client.BatchMessageNums <= 0 {
		client.BatchMessageNums = 1
	}
	client.SleepSeconds = cfg.SleepSeconds
	//client.Logger.Info()
	return client
}

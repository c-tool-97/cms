package delay_queue

import (
	"context"
	redisv8 "github.com/go-redis/redis/v8"
	"time"
)

// SendMessage :=> send message to queue
func (d *DelayQueue) SendMessage(ctx context.Context, msg *DelayQueueMessages) (err error) {
	err = d.RedisClient.ZAdd(ctx, d.Topic, &redisv8.Z{
		Score:  d.DelayTime + float64(time.Now().Unix()),
		Member: msg,
	}).Err()
	if err != nil {
		return
	}
	return
}

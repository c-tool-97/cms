package delay_queue

import (
	"context"
	"encoding/json"
	"fmt"
	redisv8 "github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"runtime"
	"time"
)

func (d *DelayQueue) ConsumerMessage(ctx context.Context, funcHandler func(in *DelayQueueMessages)) {
	defer func() {
		if r := recover(); r != nil {
			var buf [2 << 10]byte
			bytes := buf[:runtime.Stack(buf[:], true)]
			fmt.Printf("delay queue ConsumerMessage panic recover: [%s]", string(bytes))
		}
	}()

	result, err := d.RedisClient.ZRangeByScore(ctx, d.Topic, &redisv8.ZRangeBy{
		Min:   cast.ToString(d.DelayTime),
		Count: d.BatchMessageNums,
	}).Result()
	if err != nil && err != redisv8.Nil {
		return
	}
	for _, val := range result {
		msg := &DelayQueueMessages{}
		err = json.Unmarshal([]byte(val), msg)
		if err != nil {
			panic(fmt.Sprintf("delay queue msg unmarshal err: [%s]", err.Error()))
		}
		funcHandler(msg)
	}
}

// Start 初始化client
func (d *DelayQueue) Start(ctx context.Context,
	redisClient *redisv8.Client,
	cfg *DelayQueueConfig, funcHandler func(in *DelayQueueMessages)) {

	DelayQueueClient := InitDelayQueue(ctx, redisClient, cfg)
	for {
		DelayQueueClient.ConsumerMessage(ctx, funcHandler)
		if d.SleepSeconds > 0 {
			time.Sleep(time.Duration(d.SleepSeconds) * time.Second)
		}
	}
}

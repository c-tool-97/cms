package delay_queue

type DelayQueueConfig struct {
	QueueKey         string  `json:"queue_key"`  // redis key
	DelayTime        float64 `json:"delay_time"` // delay time: seconds
	BatchMessageNums int64   `json:"batch_message_nums"`
	SleepSeconds     int64   `json:"sleep_seconds"`
}

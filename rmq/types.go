package rmq

import "time"

type OnlineUser struct {
	UserId   int64
	DataTime time.Time
}

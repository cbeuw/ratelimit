package ratelimit

import "time"

type UnlimitedBucket struct{}

func NewUnlimitedBucket() *UnlimitedBucket { return &UnlimitedBucket{} }

func (ub *UnlimitedBucket) Available() int64                { return 9223372036854775807 }
func (ub *UnlimitedBucket) Capacity() int64                 { return 9223372036854775807 }
func (ub *UnlimitedBucket) Rate() float64                   { return 1.7976931348623157e+308 }
func (ub *UnlimitedBucket) Take(count int64) time.Duration  { return 0 }
func (ub *UnlimitedBucket) TakeAvailable(count int64) int64 { return count }
func (ub *UnlimitedBucket) TakeMaxDuration(count int64, maxWait time.Duration) (time.Duration, bool) {
	return 0, true
}
func (ub *UnlimitedBucket) Wait(count int64)                                        {}
func (ub *UnlimitedBucket) WaitMaxDuration(count int64, maxWait time.Duration) bool { return true }

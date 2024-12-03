package selector

import (
	"time"

	"golang.org/x/exp/rand"
)

type DateSelector interface {
	Before(moment time.Time, period time.Duration) time.Time
	BeforeNow(moment time.Time) time.Time
	After(moment time.Time, period time.Duration) time.Time
	Between(start, end time.Time) time.Time
}

type DateSelectorImpl struct{}

func NewDateSelector() DateSelector {
	return &DateSelectorImpl{}
}

func (d *DateSelectorImpl) Before(moment time.Time, period time.Duration) time.Time {
	return moment.Add(-period + time.Duration(rand.Int63n(int64(period))))
}

func (d *DateSelectorImpl) BeforeNow(moment time.Time) time.Time {
	return moment.Add(time.Duration(rand.Int63n(int64(time.Since(moment)))))
}

func (d *DateSelectorImpl) After(moment time.Time, period time.Duration) time.Time {
	return moment.Add(period + time.Duration(rand.Int63n(int64(period))))
}

func (d *DateSelectorImpl) Between(start, end time.Time) time.Time {
	return start.Add(time.Duration(rand.Int63n(int64(end.Sub(start)))))
}

package selector

import (
	"time"

	"golang.org/x/exp/rand"
)

type Selector interface {
	RandomSelect([]uint) uint
	ExponentialSelect([]uint) uint
}

type SelectorImpl struct{}

func NewSelector() Selector {
	return &SelectorImpl{}
}

func (s *SelectorImpl) RandomSelect(ids []uint) uint {
	return ids[rand.Intn(len(ids))]
}

func (s *SelectorImpl) ExponentialSelect(ids []uint) uint {
	rand.Seed(uint64(time.Now().UnixNano()))
	totalWeight := 0.0
	for i := range ids {
		totalWeight += 1.0 / float64((int(1) << i))
	}
	r := rand.Float64() * totalWeight
	accumulatedWeight := 0.0
	for i, id := range ids {
		accumulatedWeight += 1.0 / float64((int(1) << i))
		if r < accumulatedWeight {
			return id
		}
	}
	return ids[len(ids)-1]
}

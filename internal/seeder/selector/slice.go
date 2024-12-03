package selector

import (
	"time"

	"golang.org/x/exp/rand"
)

type SliceSelector interface {
	Random(length int) int
	Exponential(length int) int
}

type SliceSelectorImpl struct{}

func NewSliceSelector() SliceSelector {
	return &SliceSelectorImpl{}
}

func (s *SliceSelectorImpl) Random(length int) int {
	return rand.Intn(length)
}

func (s *SliceSelectorImpl) Exponential(length int) int {
	rand.Seed(uint64(time.Now().UnixNano()))
	totalWeight := 0.0
	for i := 0; i < length; i++ {
		totalWeight += 1.0 / float64((int(1) << i))
	}
	r := rand.Float64() * totalWeight
	accumulatedWeight := 0.0
	for i := 0; i < length; i++ {
		accumulatedWeight += 1.0 / float64((int(1) << i))
		if r < accumulatedWeight {
			return i
		}
	}
	return length - 1
}

package selector

import "golang.org/x/exp/rand"

type BoolSelector interface {
	Random() bool
	WithProbability(probability float64) bool
}

type BoolSelectorImpl struct{}

func NewBoolSelector() BoolSelector {
	return &BoolSelectorImpl{}
}

func (b *BoolSelectorImpl) Random() bool {
	return rand.Intn(2) == 1
}

func (b *BoolSelectorImpl) WithProbability(probability float64) bool {
	return rand.Float64() < probability
}

package utils

import (
	"math/rand"
	"time"
)

var defaultBackoff = BackOff{
	MaxDelay:  180 * time.Second,
	BaseDelay: 1.0 * time.Second,
	Factor:    1.6,
	Jitter:    0.2,
}

type BackOff struct {
	MaxDelay  time.Duration
	BaseDelay time.Duration
	Factor    float64
	Jitter    float64
	retries   int
}

func NewDefaultBackoff() *BackOff {
	return &defaultBackoff
}

func (bc *BackOff) NextBackOff() time.Duration {
	defer func() {
		bc.retries++
	}()
	retries := bc.retries
	if retries == 0 {
		return bc.BaseDelay
	}

	backoff, max := float64(bc.BaseDelay), float64(bc.MaxDelay)
	for backoff < max && retries > 0 {
		backoff *= bc.Factor
		retries--
	}

	if backoff > max {
		backoff = max
	}

	backoff *= 1 + bc.Jitter*(rand.Float64()*2-1)
	if backoff < 0 {
		return 0
	}

	return time.Duration(backoff)
}

func (bc *BackOff) Reset() {
	bc.retries = 0
}

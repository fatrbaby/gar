package registry

import (
	"golang.org/x/exp/rand"
	"sync/atomic"
)

type LoadBalancer interface {
	Take([]string) string
}

type RoundRobin struct {
	acc int64
}

func (r *RoundRobin) Take(endpoints []string) string {
	if len(endpoints) == 0 {
		return ""
	}

	n := atomic.AddInt64(&r.acc, 1)

	return endpoints[int(n%int64(len(endpoints)))]
}

type RandomSelect struct {
}

func (r *RandomSelect) Take(endpoints []string) string {
	if len(endpoints) == 0 {
		return ""
	}

	return endpoints[rand.Intn(len(endpoints))]
}

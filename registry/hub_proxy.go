package registry

import (
	"context"
	"strings"
	"sync"
	"time"
)
import (
	etcd "go.etcd.io/etcd/client/v3"
	"golang.org/x/time/rate"
)

var (
	proxy     *HubProxy
	proxyOnce sync.Once
)

type HubProxy struct {
	*Hub
	caches  sync.Map
	limiter *rate.Limiter
}

func NewHubProxy(services EtcdEndpoints, heartbeatFrequency int64, QPS int) *HubProxy {
	if proxy == nil {
		proxyOnce.Do(func() {
			proxy = &HubProxy{
				Hub:     NewHub(services, heartbeatFrequency),
				caches:  sync.Map{},
				limiter: rate.NewLimiter(rate.Every(time.Duration(1e9/QPS)*time.Nanosecond), QPS),
			}
		})
	}

	return proxy
}

func (h *HubProxy) Endpoints(service string) []string {
	if !h.limiter.Allow() {
		return nil
	}

	h.watchServices(service)

	if endpoints, has := h.caches.Load(service); has {
		return endpoints.([]string)
	}

	endpoints := h.Hub.Endpoints(service)

	if len(endpoints) > 0 {
		h.caches.Store(service, endpoints)
	}

	return endpoints
}

func (h *HubProxy) watchServices(service string) {
	if _, has := h.watches.LoadOrStore(service, true); has {
		return
	}

	ctx := context.Background()
	prefix := buildService(ServiceRoot, service, "")
	wc := h.client.Watch(ctx, prefix, etcd.WithPrefix())

	go func() {
		for response := range wc {
			for _, event := range response.Events {
				paths := strings.Split(string(event.Kv.Key), "")

				if len(paths) > 2 {
					service := paths[len(paths)-2]
					endpoints := h.Hub.Endpoints(service)

					if len(endpoints) > 0 {
						h.caches.Store(service, endpoints)
					} else {
						h.caches.Delete(service)
					}
				}
			}
		}
	}()
}

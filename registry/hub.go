package registry

import (
	"context"
	etcd "go.etcd.io/etcd/client/v3"
	"log"
	"strings"
	"sync"
	"time"
)

const ServiceRoot = "/synod/gar/index"

var (
	hub     *Hub
	hubOnce sync.Once
)

type EtcdEndpoints []string

type Hub struct {
	client             *etcd.Client
	heartbeatFrequency int64
	watches            *sync.Map
	loadBalancer       LoadBalancer
}

func NewHub(servers EtcdEndpoints, heartbeatFrequency int64) *Hub {
	if hub == nil {
		hubOnce.Do(func() {
			client, err := etcd.New(etcd.Config{
				Endpoints:   servers,
				DialTimeout: 3 * time.Second,
			})

			if err != nil {
				log.Fatalf("create etcd client error: %v\n", err)
			}

			hub = &Hub{
				client:             client,
				heartbeatFrequency: heartbeatFrequency,
				loadBalancer:       &RoundRobin{},
			}
		})
	}

	return hub
}

func (h *Hub) Register(service string, endpoint string, leaseID etcd.LeaseID) (etcd.LeaseID, error) {
	ctx := context.Background()

	if leaseID <= 0 {
		lease, err := h.client.Grant(ctx, h.heartbeatFrequency)

		if err != nil {

		}

		key := buildService(ServiceRoot, service, endpoint)

		_, err = h.client.Put(ctx, key, "", etcd.WithLease(lease.ID))

		if err != nil {
			return 0, err
		}

		return lease.ID, nil
	}

	_, err := h.client.KeepAliveOnce(ctx, leaseID)

	if err != nil {
		return leaseID, err
	}

	return leaseID, nil
}

func (h *Hub) Offline(service, endpoint string) error {
	ctx := context.Background()

	key := buildService(ServiceRoot, service, endpoint)

	if _, err := h.client.Delete(ctx, key); err != nil {
		return err
	}

	return nil
}

func (h *Hub) Endpoints(service string) []string {
	ctx := context.Background()
	prefix := buildService(ServiceRoot, service)

	response, err := h.client.Get(ctx, prefix, etcd.WithPrefix())

	if err != nil {
		return nil
	}

	endpoints := make([]string, 0, len(response.Kvs))

	for _, kv := range response.Kvs {
		paths := strings.Split(string(kv.Key), "/")
		endpoints = append(endpoints, paths[len(paths)-1])
	}

	return endpoints
}

func (h *Hub) Endpoint(service string) string {
	return h.loadBalancer.Take(h.Endpoints(service))
}

func (h *Hub) Close() error {
	return h.client.Close()
}

func buildService(parts ...string) string {
	return strings.Join(parts, "/")
}

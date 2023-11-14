package clusters

import (
	"fmt"
	"massive-screen/internal/contracts"
	"math/rand"
	"sync"
)

type Cluster struct {
	Server  string `json:"server"`
	Loading int    `json:"loading"`
}

type Service interface {
	contracts.Service

	Load(i *LoadRequest) (*Cluster, error)
	Stat(i *StatRequest) ([]*Cluster, error)
}

type service struct {
	*contracts.AService
}

func New(name string) Service {
	return &service{
		AService: contracts.NewService(name),
	}
}

type StatRequest struct {
	Servers []string
}

func (s *service) Stat(i *StatRequest) ([]*Cluster, error) {
	// validate input or return validation error

	var wg sync.WaitGroup
	resCh := make(chan *Cluster, len(i.Servers))
	for _, address := range i.Servers {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()

			cluster, err := s.Load(&LoadRequest{
				Addr: addr,
			})
			if err != nil {
				fmt.Printf("failed to get cluster loading: %v", err)
				return
			}

			resCh <- cluster
		}(address)
	}

	// close res channel when all goroutines done
	go func() {
		wg.Wait()
		close(resCh)
	}()

	var clusters []*Cluster
	for c := range resCh {
		clusters = append(clusters, c)
	}

	return clusters, nil
}

type LoadRequest struct {
	Addr string
}

func (s *service) Load(i *LoadRequest) (*Cluster, error) {
	// mock: validate input
	// mock: finding process...
	return &Cluster{
		Server:  i.Addr,
		Loading: rand.Intn(101),
	}, nil
}

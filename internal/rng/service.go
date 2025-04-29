package rng

import (
	"log"

	"piggy-bank/config"
)

type Service struct {
	client Client
}

func NewService(cfg *config.Config) (*Service, error) {
	var (
		client Client
		err    error
	)

	useMock := cfg.RNG.UseMock
	usePool := cfg.RNG.UsePool

	if useMock {
		log.Printf("UseMock: %v", useMock)
		client, err = NewMockClient(cfg)
	} else if usePool {
		log.Printf("UsePool: %v", usePool)
		client, err = NewWithPoolClient(cfg)
	} else {
		client, err = NewSimpleClient(cfg)
	}

	if err != nil {
		log.Printf("Error creating RNG client: %v, using fallback MockClient", err)
		client, err = NewMockClient(cfg)
		if err != nil {
			return nil, err
		}
	}

	return &Service{
		client: client,
	}, nil
}

func (s *Service) GetClient() Client {
	return s.client
}

func (s *Service) Rand(max uint64) (uint64, error) {
	return s.client.Rand(max)
}

func (s *Service) RandSlice(maxSlice []uint64) ([]uint64, error) {
	return s.client.RandSlice(maxSlice)
}

func (s *Service) RandFloat() (float64, error) {
	return s.client.RandFloat()
}

func (s *Service) RandFloatSlice(count int) ([]float64, error) {
	return s.client.RandFloatSlice(count)
}

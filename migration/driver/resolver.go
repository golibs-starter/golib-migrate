package driver

import "errors"

var ErrDriverNotFound = errors.New("driver not found")

type Resolver struct {
	strategiesMap map[string]Strategy
}

func NewResolver(strategies []Strategy) *Resolver {
	m := make(map[string]Strategy)
	for _, strategy := range strategies {
		m[strategy.Driver()] = strategy
	}
	return &Resolver{strategiesMap: m}
}

func (r Resolver) Resolve(driver string) (Strategy, error) {
	strategy, exists := r.strategiesMap[driver]
	if !exists {
		return nil, ErrDriverNotFound
	}
	return strategy, nil
}

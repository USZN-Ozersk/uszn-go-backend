package store

import "uszn-go-backend/internal/app/config"

// Store ...
type Store struct {
	config *config.Config
}

// New ...
func New(config *config.Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	return nil
}

// Close ...
func (s *Store) Close() {
	// ...
}

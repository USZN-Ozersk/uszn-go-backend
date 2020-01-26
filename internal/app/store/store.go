package store

import (
	"database/sql"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/config"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/logger"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config *config.Config
	logger *logger.Logger
	Db     *sql.DB
}

// New ...
func New(config *config.Config, logger *logger.Logger) *Store {
	return &Store{
		config: config,
		logger: logger,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DB_URL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.Db = db

	s.logger.Logger.Info("Database connected")

	return nil
}

// Close ...
func (s *Store) Close() {
	s.Db.Close()
	s.logger.Logger.Info("Database connected")
}

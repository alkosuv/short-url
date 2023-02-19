package service

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"

	"github.com/gen95mis/short-url/internal/database"
)

type Service struct {
	db database.Database
}

func New(db database.Database) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Get(hash string) (original string, err error) {
	return s.db.GetByHash(hash)
}

func (s *Service) Set(original string) (string, error) {
	shortened, err := s.db.GetByOriginal(original)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	if shortened != "" {
		return shortened, nil
	}

	algorithm := md5.New()
	algorithm.Write([]byte(original))
	hash := hex.EncodeToString(algorithm.Sum(nil))

	shortenedHash := string(hash[:7])

	if err := s.db.Set(original, shortenedHash); err != nil {
		return "", nil
	}

	return shortenedHash, nil
}

package service

import (
	"crypto"
	"database/sql"
	"encoding/hex"
	"errors"
	"hash"

	"github.com/gen95mis/short-url/internal/database"
)

type Service struct {
	db database.Database
	md hash.Hash
}

func New(db database.Database) *Service {
	return &Service{
		db: db,
		md: crypto.MD5.New(),
	}
}

func (s *Service) Get(hash string) (original string, err error) {
	return s.db.GetByHash(hash[1:])
}

func (s *Service) Set(original string) (string, error) {
	shortened, err := s.db.GetByOriginal(original)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	if shortened != "" {
		return shortened, nil
	}

	hash := hex.EncodeToString(s.md.Sum(nil))

	if err := s.db.Set(original, string(hash[:7])); err != nil {
		return "", nil
	}

	return shortened, nil
}

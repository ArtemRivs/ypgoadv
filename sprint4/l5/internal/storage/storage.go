package storage

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/ArtemRivs/ypgoadv/sprint4/l5/internal/storage/sqlstorage"
)

type Storage interface {
	SaveOriginLink(shortCode, longURL, userID string) (string, error)
	GetOriginLink(shortCode string) (string, error)
	GetUserOriginLinks(userID string) (map[string]string, error)
	DeleteUserLinks(ctx context.Context, userID string, shortCodes []string) error
	Close()
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIKLMNOPQRSTUVWXYZ0123456789"

func New(dataBaseDsn, fileStoragePath string) Storage {
	var ls Storage
	if dataBaseDsn != "" {
		log.Println("use sql-storage:", dataBaseDsn)
		ls = sqlstorage.New(dataBaseDsn)
		return ls
	}
	return nil
}

func GenerateCode() string {
	sRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[sRand.Intn(len(charset))]
	}
	return string(b)
}

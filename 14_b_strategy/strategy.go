package strategy

import (
	"fmt"
	"os"
)

// 存储策略
type StorageStrategy interface {
	Save(name string, data []byte) error
}

var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}
	return s, nil
}

// 保存到文件
type fileStorage struct{}

func (s *fileStorage) Save(name string, data []byte) error {
	return os.WriteFile(name, data, os.ModeAppend)
}

// 加密保存到文件
type encryptFileStorage struct{}

func (S *encryptFileStorage) Save(name string, data []byte) error {
	// 加密
	data, err := encrypt(data)
	if err != nil {
		return err
	}

	return os.WriteFile(name, data, os.ModeAppend)
}

func encrypt(data []byte) ([]byte, error) {
	return data, nil
}

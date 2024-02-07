package storage

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"log"
	"os"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/domain"
)

type StorageClient struct {
	ctx context.Context
	cfg *configs.Configs
}

const (
	imageBase = "./images/"
)

func (s *StorageClient) StoreFile(file *domain.File) domain.Error {
	img, _, err := image.Decode(bytes.NewReader(file.Data))
	if err != nil {
		return domain.NewError().SetError(err)
	}

	out, _ := os.Create(imageBase + file.Filename)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 1

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	return nil
}

func NewStorageClient(ctx context.Context, cfg *configs.Configs) (*StorageClient, error) {
	log.Println("connecting to Storage")

	return &StorageClient{ctx, cfg}, nil
}

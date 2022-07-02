package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"upload-file/internal/biz"
)

type upload-fileRepo struct {
	data *Data
	log  *log.Helper
}

// NewUpload-FileRepo .
func NewUpload-FileRepo(data *Data, logger log.Logger) biz.Upload-FileRepo {
	return &upload-fileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *upload-fileRepo) Save(ctx context.Context, b *biz.Upload-File) (*biz.Upload-File, error) {
	return b, nil
}

func (r *upload-fileRepo) Update(ctx context.Context, b *biz.Upload-File) (*biz.Upload-File, error) {
	return b, nil
}

func (r *upload-fileRepo) FindByID(context.Context, int64) (*biz.Upload-File, error) {
	return nil, nil
}

func (r *upload-fileRepo) ListByHello(context.Context, string) ([]*biz.Upload-File, error) {
	return nil, nil
}

func (r *upload-fileRepo) ListAll(context.Context) ([]*biz.Upload-File, error) {
	return nil, nil
}

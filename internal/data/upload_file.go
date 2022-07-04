package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"upload-file/internal/biz"
)

type uploadFileRepo struct {
	data *Data
	log  *log.Helper
}

// NewUploadFileRepo .
func NewUploadFileRepo(data *Data, logger log.Logger) biz.UploadFileRepo {
	return &uploadFileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *uploadFileRepo) Save(ctx context.Context, b *biz.UploadFile) (*biz.UploadFile, error) {
	return b, nil
}

func (r *uploadFileRepo) Update(ctx context.Context, b *biz.UploadFile) (*biz.UploadFile, error) {
	return b, nil
}

func (r *uploadFileRepo) FindByID(context.Context, int64) (*biz.UploadFile, error) {
	return nil, nil
}

func (r *uploadFileRepo) ListByHello(context.Context, string) ([]*biz.UploadFile, error) {
	return nil, nil
}

func (r *uploadFileRepo) ListAll(context.Context) ([]*biz.UploadFile, error) {
	return nil, nil
}

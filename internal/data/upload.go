package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"

	"upload/internal/biz"
	v1 "upload/proto/api/upload/v1"
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

func (r *uploadFileRepo) Save(ctx context.Context, b *biz.UploadFileInfo) error {
	_, err := r.data.OSS.GetClient().PutObject(ctx, b.Bucket, b.Name, *b.File, -1,
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Errorf("oss upload failed. err=%v ctx=%v", err, ctx)
		return v1.ErrorInternalServer("oss upload failed")
	}
	return nil
}

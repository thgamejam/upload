package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Upload-File is a Upload-File model.
type Upload-File struct {
	Hello string
}

// Upload-FileUseCase is a Upload-File use case.
type Upload-FileUseCase struct {
	repo Upload-FileRepo
	log  *log.Helper
}

// NewUpload-FileUseCase new a Upload-File use case.
func NewUpload-FileUseCase(repo Upload-FileRepo, logger log.Logger) *Upload-FileUseCase {
	return &Upload-FileUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUpload-File creates a Upload-File, and returns the new Upload-File.
func (uc *Upload-FileUseCase) CreateUpload-File(ctx context.Context, g *Upload-File) (*Upload-File, error) {
	uc.log.WithContext(ctx).Infof("CreateUpload-File: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

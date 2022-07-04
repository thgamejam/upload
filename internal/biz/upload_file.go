package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// UploadFile is a UploadFile model.
type UploadFile struct {
	Hello string
}

// UploadFileUseCase is a UploadFile use case.
type UploadFileUseCase struct {
	repo UploadFileRepo
	log  *log.Helper
}

// NewUploadFileUseCase new a UploadFile use case.
func NewUploadFileUseCase(repo UploadFileRepo, logger log.Logger) *UploadFileUseCase {
	return &UploadFileUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUploadFile creates a UploadFile, and returns the new UploadFile.
func (uc *UploadFileUseCase) CreateUploadFile(ctx context.Context, g *UploadFile) (*UploadFile, error) {
	uc.log.WithContext(ctx).Infof("CreateUploadFile: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

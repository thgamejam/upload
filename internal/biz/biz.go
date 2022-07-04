package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUploadFileUseCase)

// UploadFileRepo is a UploadFile repo.
type UploadFileRepo interface {
	Save(context.Context, *UploadFile) (*UploadFile, error)
	Update(context.Context, *UploadFile) (*UploadFile, error)
	FindByID(context.Context, int64) (*UploadFile, error)
	ListByHello(context.Context, string) ([]*UploadFile, error)
	ListAll(context.Context) ([]*UploadFile, error)
}

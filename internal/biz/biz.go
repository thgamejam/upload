package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUploadFileUseCase)

// UploadFileRepo is a UploadFileInfo repo.
type UploadFileRepo interface {
	Save(context.Context, *UploadFileInfo) error
}

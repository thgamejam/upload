package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUpload-FileUseCase)

// Upload-FileRepo is a Upload-File repo.
type Upload-FileRepo interface {
	Save(context.Context, *Upload-File) (*Upload-File, error)
	Update(context.Context, *Upload-File) (*Upload-File, error)
	FindByID(context.Context, int64) (*Upload-File, error)
	ListByHello(context.Context, string) ([]*Upload-File, error)
	ListAll(context.Context) ([]*Upload-File, error)
}

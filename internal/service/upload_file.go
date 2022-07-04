package service

import (
	"upload-file/internal/biz"
	v1 "upload-file/proto/api/upload-file/v1"
)

// UploadFileService is a upload-file service.
type UploadFileService struct {
	v1.UnimplementedUploadFileServer

	uc *biz.UploadFileUseCase
}

// NewUploadFileService new a upload-file service.
func NewUploadFileService(uc *biz.UploadFileUseCase) *UploadFileService {
	return &UploadFileService{uc: uc}
}

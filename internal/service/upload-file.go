package service

import (
	"context"
	
	"upload-file/internal/biz"
	v1 "upload-file/proto/api/upload-file/v1"
)

// Upload-FileService is a upload-file service.
type Upload-FileService struct {
	v1.UnimplementedUpload-FileServer

	uc *biz.Upload-FileUseCase
}

// NewUpload-FileService new a upload-file service.
func NewUpload-FileService(uc *biz.Upload-FileUseCase) *Upload-FileService {
	return &Upload-FileService{uc: uc}
}

// SayHello implements helloworld.Upload-FileServer.
func (s *Upload-FileService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateUpload-File(ctx, &biz.Upload-File{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

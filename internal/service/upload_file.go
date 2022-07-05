package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"upload-file/internal/middleware"

	"upload-file/internal/biz"
	pb "upload-file/proto/api/upload-file/v1"
)

type UploadFileService struct {
	pb.UnimplementedUploadFileServer

	uc  *biz.UploadFileUseCase
	log *log.Helper
}

func NewUploadFileService(uc *biz.UploadFileUseCase, logger log.Logger) *UploadFileService {
	return &UploadFileService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UploadFileService) UploadFile(ctx context.Context, req *pb.UploadFileReq) (*pb.UploadFileReply, error) {
	file, ok := middleware.FromUploadFileContext(ctx)
	log.Debugf("ok=%v, file=%v\n", ok, file)
	log.Debugf("req=%v\n", req)
	return &pb.UploadFileReply{}, nil
}

func (s *UploadFileService) UploadSliceFile(ctx context.Context, req *pb.UploadSliceFileReq) (*pb.UploadFileReply, error) {
	return &pb.UploadFileReply{}, nil
}

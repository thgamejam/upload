package service

import (
	"context"
	"github.com/thgamejam/pkg/uuid"
	"upload/internal/conf"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/thgamejam/pkg/authentication"
	"upload/internal/biz"
	"upload/internal/middleware"
	pb "upload/proto/api/upload/v1"
)

type UploadFileService struct {
	pb.UnimplementedUploadFileServer

	uc  *biz.UploadFileUseCase
	log *log.Helper

	SecretKey string // 校验上传链接使用的密钥
}

func NewUploadFileService(c *conf.UploadFile, uc *biz.UploadFileUseCase, logger log.Logger) *UploadFileService {
	return &UploadFileService{
		uc:  uc,
		log: log.NewHelper(logger),

		SecretKey: c.SecretKey,
	}
}

func (s *UploadFileService) UploadFile(ctx context.Context, req *pb.UploadFileReq) (*pb.UploadFileReply, error) {
	// 校验连接是否正确
	fileUUID, err := uuid.Parse(req.Uuid)
	if err != nil {
		// 错误的uuid
		return nil, pb.ErrorWrongUuid("")
	}
	var claims authentication.UploadFileClaims
	claims = authentication.UploadFileClaims{
		Bucket:    req.Bucket,
		Name:      req.Name,
		UUID:      fileUUID,
		ExpiresAt: req.Exp,
		CRC:       req.Crc,
		SHA1:      req.Sha1,
	}
	ok, err := authentication.ValidateUploadInfo(&claims, &s.SecretKey, &req.Hash)
	if err != nil {
		// 不正常的url校验退出
		return nil, pb.ErrorUrlValidationException("")
	}
	if !ok {
		// 未经授权的上传
		return nil, pb.ErrorUnauthorizedUpload("", claims, s.SecretKey, req.Hash)
	}

	// 校验会话是否可用
	ok, err = s.uc.ValidationSession(fileUUID.String())
	if err != nil {
		return nil, err
	}
	if !ok {
		// 错误的会话
		return nil, pb.ErrorWrongSession("session_uuid=%s", req.Uuid)
	}

	// 停用链接
	defer func(s *UploadFileService, uuid string) {
		err := s.uc.CloseUploadSession(uuid)
		if err != nil {
			s.log.Errorf("Unable to disable upload url. err=%v\n", err)
		}
	}(s, req.Uuid)

	file, ok := middleware.FromUploadFileContext(ctx)
	// 无法获取文件
	if !ok || file == nil {
		// 上传文件丢失
		return nil, pb.ErrorUploadFileMissing("")
	}
	var info biz.UploadFileInfo
	info = biz.UploadFileInfo{
		Bucket: req.Bucket,
		Name:   req.Name,
		UUID:   fileUUID,
		CRC:    req.Crc,
		SHA1:   req.Sha1,
		File:   file,
	}
	err = s.uc.UploadFile(ctx, &info)
	if err != nil {
		return nil, err
	}

	return &pb.UploadFileReply{}, nil
}

func (s *UploadFileService) UploadSliceFile(ctx context.Context, req *pb.UploadSliceFileReq) (*pb.UploadFileReply, error) {
	return &pb.UploadFileReply{}, nil
}

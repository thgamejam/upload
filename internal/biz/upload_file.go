package biz

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"hash/crc32"
	"io"
	"mime/multipart"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/thgamejam/pkg/uuid"
	v1 "upload-file/proto/api/upload-file/v1"
)

var validationCRC32 = func(crc32hash string, file *multipart.File) error {
	hash := crc32.NewIEEE()
	_, err := io.Copy(hash, *file)
	if err != nil {
		// 损坏的文件
		return v1.ErrorDamagedFile("")
	}

	value := hash.Sum(nil)
	if crc32hash != hex.EncodeToString(value) {
		// 错误的CRC校验
		return v1.ErrorWrongCrcCheck("")
	}

	// 还原文件光标
	_, err = (*file).Seek(0, 0)
	if err != nil {
		// 不正常的校验退出
		return v1.ErrorFileValidationException("")
	}

	return nil
}

var validationSHA1 = func(sha1hash string, file *multipart.File) error {
	hash := sha1.New()
	_, err := io.Copy(hash, *file)
	if err != nil {
		// 损坏的文件
		return v1.ErrorDamagedFile("")
	}
	value := hash.Sum(nil)
	if sha1hash != hex.EncodeToString(value) {
		// 错误的CRC校验
		return v1.ErrorWrongSha1Check("")
	}
	// 还原文件光标
	_, err = (*file).Seek(0, 0)
	if err != nil {
		// 不正常的文件校验退出
		return v1.ErrorFileValidationException("")
	}

	return nil
}

// UploadFileInfo 上传文件的信息
type UploadFileInfo struct {
	Bucket string          // 需要上传文件所在的oss桶
	Name   string          // oss中设置的文件名
	UUID   uuid.UUID       // 链接的唯一id 实现幂等性
	CRC    string          // 上传文件的crc-32-hash值
	SHA1   string          // 上传文件的sha1-hash值
	File   *multipart.File // 文件
}

// UploadFileUseCase is a UploadFileInfo use case.
type UploadFileUseCase struct {
	repo UploadFileRepo
	log  *log.Helper
}

// NewUploadFileUseCase new a UploadFileInfo use case.
func NewUploadFileUseCase(repo UploadFileRepo, logger log.Logger) *UploadFileUseCase {
	return &UploadFileUseCase{repo: repo, log: log.NewHelper(logger)}
}

// UploadFile 上传文件
func (uc *UploadFileUseCase) UploadFile(ctx context.Context, g *UploadFileInfo) error {

	// 计算crc32值
	err := validationCRC32(g.CRC, g.File)
	if err != nil {
		return err
	}

	// 计算sha1值
	err = validationSHA1(g.SHA1, g.File)
	if err != nil {
		return err
	}

	return uc.repo.Save(ctx, g)
}

// CloseUploadSession 关闭上传链接
// 保证幂等性，是的每个uuid对于的上传链接只能使用一次
func (uc *UploadFileUseCase) CloseUploadSession(uuid string) error {
	return nil
}

// ValidationSession 验证会话是否可用
func (uc *UploadFileUseCase) ValidationSession(uuid string) (bool, error) {
	return true, nil
}

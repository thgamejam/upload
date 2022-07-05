package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	kratosHTTP "github.com/go-kratos/kratos/v2/transport/http"
	"mime/multipart"
)

type (
	// uploadFileKey context中上传文件的键
	uploadFileKey struct{}
)

// UploadFileMiddleware 获取http中上传文件的中间件
func UploadFileMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			kratosContext, ok := ctx.(kratosHTTP.Context)
			_ = kratosContext
			if !ok {
				return handler(ctx, req)
			}

			request := kratosContext.Request()

			// 从http请求中获取文件
			file, _, err := request.FormFile("file")
			if err != nil {
				return handler(ctx, req)
			}
			defer file.Close()
			// 将文件添加到上下文
			ctx = context.WithValue(ctx, uploadFileKey{}, &file)

			return handler(ctx, req)
		}
	}
}

// FromUploadFileContext 返回存储在context中的上传文件，如果有
func FromUploadFileContext(ctx context.Context) (file *multipart.File, ok bool) {
	file, ok = ctx.Value(uploadFileKey{}).(*multipart.File)
	return
}

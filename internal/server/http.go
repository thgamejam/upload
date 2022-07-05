package server

import (
	"io/ioutil"
	"strings"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"

	"upload-file/internal/conf"
	"upload-file/internal/middleware"
	"upload-file/internal/service"
	v1 "upload-file/proto/api/upload-file/v1"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, service *service.UploadFileService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			middleware.UploadFileMiddleware(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	// 更改默认的数据解析方法
	opts = append(opts, http.RequestDecoder(DecodeRequest))
	srv := http.NewServer(opts...)
	v1.RegisterUploadFileHTTPServer(srv, service)
	return srv
}

func DecodeRequest(r *http.Request, v interface{}) error {
	// 从Request Header的Content-Type中提取出对应的解码器
	codec, ok := codecForRequest(r, "Content-Type")
	// 如果找不到对应的解码器则认为Content-Type="multipart/form-data"
	if !ok {
		return nil
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.BadRequest("CODEC", err.Error())
	}
	if err = codec.Unmarshal(data, v); err != nil {
		return errors.BadRequest("CODEC", err.Error())
	}
	return nil
}

// codecForRequest in https://github.com/go-kratos/kratos/blob/main/transport/http/codec.go
func codecForRequest(r *http.Request, name string) (encoding.Codec, bool) {
	for _, accept := range r.Header[name] {
		codec := encoding.GetCodec(contentSubtype(accept))
		if codec != nil {
			return codec, true
		}
	}
	return encoding.GetCodec("json"), false
}

// contentSubtype in https://github.com/go-kratos/kratos/blob/main/internal/httputil/http.go
func contentSubtype(contentType string) string {
	left := strings.Index(contentType, "/")
	if left == -1 {
		return ""
	}
	right := strings.Index(contentType, ";")
	if right == -1 {
		right = len(contentType)
	}
	if right < left {
		return ""
	}
	return contentType[left+1 : right]
}

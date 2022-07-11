package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/thgamejam/pkg/cache"
	"github.com/thgamejam/pkg/object_storage"
	"upload/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUploadFileRepo,
)

// Data .
type Data struct {
	// 封装的数据库客户端
	Cache *cache.Cache
	OSS   *object_storage.ObjectStorage
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	redis, err := cache.NewCache(c.Redis)
	if err != nil {
		return nil, nil, err
	}
	oss, err := object_storage.NewObjectStorage(c.ObjectStorage)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		l := log.NewHelper(logger)
		err := redis.Close()
		if err != nil {
			l.Errorf("close cache err=%v", err)
		}
		l.Info("closing the data resources.")
	}

	return &Data{
		// 装填数据库客户端
		Cache: redis,
		OSS:   oss,
	}, cleanup, nil
}

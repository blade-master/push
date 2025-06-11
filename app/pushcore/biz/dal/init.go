package dal

import (
	"github.com/blade-master/push/app/pushcore/biz/dal/mysql"
	"github.com/blade-master/push/app/pushcore/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

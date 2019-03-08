package main

import (
	"github.com/dy-dayan/common-srv-atomicid/idl/dayan/common/srv-atomicid"
	"github.com/dy-dayan/common-srv-atomicid/dal/db"
	"github.com/dy-dayan/common-srv-atomicid/handler"
	"github.com/dy-dayan/common-srv-atomicid/util/config"
	"github.com/dy-gopkg/kit/micro"
	"github.com/sirupsen/logrus"
)

func main() {
	micro.Init()

	// 初始化配置
	uconfig.Init()

	// 初始化数据库
	db.Init()

	//TODO 初始化缓存
	//cache.CacheInit()

	err := dayan_common_srv_atomicid.RegisterAtomicIDHandler(micro.Server(), &handler.Handler{})
	if err != nil {
		logrus.Fatalf("RegisterPassportHandler error:%v", err)
	}

	micro.Run()
}
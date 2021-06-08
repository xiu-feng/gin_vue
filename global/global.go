package global

import (
	"github.com/spf13/viper"
	"myself_rep/gin_vue/conf"
	"github.com/jinzhu/gorm"
)

//定义全局变量
var (
	GVA_DB *gorm.DB	//全局数据库连接
	GVA_CONFIG *conf.Server //全局配置
	GVA_VP *viper.Viper //全局viper
	LOGINNULL=10000
	LOGINLEN =10001
	LOGINUS =10002
	SUCCESS = 200
)

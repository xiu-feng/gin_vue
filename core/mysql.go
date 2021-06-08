package core

//连接mysql
import (
"fmt"
	"myself_rep/gin_vue/global"
	"time"

_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
)

func init(){
	dirverName := global.GVA_CONFIG.Database.DBType
	host := global.GVA_CONFIG.Database.Host
	database := global.GVA_CONFIG.Database.DBName
	username := global.GVA_CONFIG.Database.UserName
	password := global.GVA_CONFIG.Database.Password
	charset := global.GVA_CONFIG.Database.Charset
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		database,
		charset)

	db, err := gorm.Open(dirverName, args)
	if err != nil {
		panic("filed to connect database, err:" + err.Error())
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Hour)
	global.GVA_DB = db
}

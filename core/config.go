package core


//加载配置文件
import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"myself_rep/gin_vue/global"
)

// 配置文件路径 (包名 + 配置文件名 )
const defaultConfigFile = "conf/config.yaml"

// 初始化配置文件
func init(){
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)

	// 读取配置文件中的配置信息，并将信息保存 到 v中
	err := v.ReadInConfig()
	if err !=nil {
		panic(fmt.Errorf("Fatal error config file: #{err}\n"))
	}
	// 监控配置文件
	v.WatchConfig()

	// 配置文件改变，则将 v中的配置信息，刷新到 global.GVA_CONFIG
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:",e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG);err !=nil {
			fmt.Println(err)
		}
	})

	// 将 v 中的配置信息 反序列化成 结构体 (将v 中配置信息 刷新到 global.GVA_CONFIG)
	if err := v.Unmarshal(&global.GVA_CONFIG);err !=nil {
		fmt.Println(err)
	}
	fmt.Println(global.GVA_CONFIG)

	// 保存 viper 实例 v
	global.GVA_VP = v
}

package main
import (
	"bytes"
	"github.com/spf13/viper"
	"fmt"
	"log"
)

/*
 * 读取配置文件
 * 	Get 自动判断类型
 *  GetTime 获取时间类型
 *  GetString 获取字符串类型
 *  GetBool  获取bool类型
 *	GetInt	 获取bool类型
 *	GetStringMap  获取map类型
 *
 *  指定配置文件的全路径
 *  	SetConfigFile("文件")
 *	
 *  Re
 * 	
 */
func main()  {
	viper.SetConfigType("toml")
	tomlConfig := []byte(`
  app_name = "awesome web"
  
  # possible values: DEBUG, INFO, WARNING, ERROR, FATAL
  log_level = "DEBUG"
  
  [mysql]
  ip = "127.0.0.1"
  port = 3306
  user = "dj"
  password = 123456
  database = "awesome"
  
  [redis]
  ip = "127.0.0.1"
  port = 7381
  `)
	err := viper.ReadConfig(bytes.NewBuffer(tomlConfig))
	if err != nil {
	  log.Fatal("read config failed: %v", err)
	}
  
	fmt.Println(viper.GetString("log_level"))
	fmt.Println("redis port: ", viper.GetInt("redis.port"))
	fmt.Println("mysql database: ", viper.GetString("mysql.database") )
}

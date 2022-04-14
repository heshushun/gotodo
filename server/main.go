package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
	"todos/config"
	"todos/model"
	"todos/router"
)

// * 读取命令行配置（config文件）
var cfg = pflag.StringP("config", "c", "", "api server config file path.")

// @BasePath /v1
func main() {

	// * 初始化配置
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	log.Printf("The config inited.")

	// * 初始化数据库
	model.DB.Init()
	defer model.DB.Close()
	log.Printf("The database inited.")

	// * 设置 run mode
	gin.SetMode(viper.GetString("runmode"))
	log.Printf("Gin run mode set to: %s", viper.GetString("runmode"))

	// * gin engine.
	g := gin.New()
	router.Load(g)
	log.Printf("The gin engine started, and the router loaded.")

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Println("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening on: %s", viper.GetString("addr"))
	g.Run(viper.GetString("addr")).Error()
}

// * api 状态自检
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}

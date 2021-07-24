package bootstrap

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"
	"go-prow/app"
	"go-prow/app/api"
	"go-prow/app/consumer"
	"go-prow/app/models"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

var (
	prowConfig = config.ReadFromJson("./config/config.json")
)

func init() {
	models.Setup()
	consumer.Setup()
}

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := app.InitRouter()
	template.AddComp(chartjs.NewChart())
	eng := engine.Default()
	//adminPlugin := admin.NewAdmin(datamodel.Generators)
	//adminPlugin.AddGenerator("user", datamodel.GetUserTable)
	//login.Init(login.Config{
	//	Theme: "theme2",
	//	//CaptchaDigits: 5, // 使用图片验证码，这里代表多少个验证码数字
	//	// 使用腾讯验证码，需提供appID与appSecret
	//	// TencentWaterProofWallData: login.TencentWaterProofWallData{
	//	//    AppID:"",
	//	//    AppSecret: "",
	//	// }
	//})
	if err := eng.AddConfig(&prowConfig).
		AddGenerators(models.Generators).
		//AddPlugins(adminPlugin).
		Use(r); err != nil {
		panic(err)
	}

	//adminPlugin.SetCaptcha(map[string]string{"driver": login.CaptchaDriverKeyDefault})
	r.Static("/uploads", "./uploads")
	eng.HTML("GET", "/", api.GetDashBoard)
	eng.HTML("GET", "/prow", api.GetDashBoard)
	//eng.HTMLFile("GET", "/goprow/hello", "./html/hello.tmpl", map[string]interface{}{
	//	//	"msg": "Hello world",
	//	//})

	_ = r.Run(":9033")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}

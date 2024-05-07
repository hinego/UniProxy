package handle

import (
	"UniProxy/proxy"
	"UniProxy/v2b"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type initParamsRequest struct {
	MixedPort int    `json:"mixed_port"`
	AppName   string `json:"app_name"`
	Url       string `json:"url"`
	Token     string `json:"token"`
	License   string `json:"license"`
	UserPath  string `json:"user_path"`
}

var inited bool

func InitParamsManual() {

}
func InitParams2() {
	p := initParamsRequest{
		MixedPort: 33210,
		AppName:   "测试加速",
		Url:       "https://api.nodebackapis.top",
		Token:     "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6Mywic2Vzc2lvbiI6IjVhMzBjYTM1OTIxNGMzNTE5NTYxY2ZlYzViNGRiYjc0In0.2SmJXQ_Rd0LEuEEAMbbrm34tspyjwdECJv270_i_3JI",
		License:   "1e0ad0b950993ce3382f9ee4ca61f572d56edec2",
		UserPath:  "D:\\opt\\workplace\\golang\\UniProxy\\cmd\\uniproxy",
	}
	f, err := os.OpenFile(path.Join(p.UserPath, "uniproxy.log"), os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	log.SetOutput(f)
	v2b.Init(p.Url, p.Token)
	proxy.InPort = p.MixedPort
	proxy.DataPath = p.UserPath
	inited = true
}

func InitParams(c *gin.Context) {
	// p := initParamsRequest{}
	// err := c.ShouldBindJSON(&p)
	// if err != nil {
	// 	c.JSON(400, &Rsp{Success: false, Message: err.Error()})
	// 	return
	// }
	// if encrypt.Sha([]byte(encrypt.Sha([]byte(p.Url))+"1145141919")) != p.License {
	// 	c.JSON(400, &Rsp{Success: false})
	// 	return
	// }
	// f, err := os.OpenFile(path.Join(p.UserPath, "uniproxy.log"), os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0755)
	// if err != nil {
	// 	c.JSON(400, &Rsp{Success: false, Message: err.Error()})
	// 	return
	// }
	// log.SetOutput(f)
	// v2b.Init(p.Url, p.Token)
	// proxy.InPort = p.MixedPort
	// proxy.DataPath = p.UserPath
	// inited = true
	c.JSON(200, &Rsp{Success: true})
}

func GetStatus(c *gin.Context) {
	c.JSON(200, &Rsp{
		Success: true,
		Data: StatusData{
			Inited:      inited,
			Running:     proxy.Running,
			GlobalMode:  proxy.GlobalMode,
			SystemProxy: proxy.SystemProxy,
		},
	})
}

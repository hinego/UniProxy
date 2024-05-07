package handle

import (
	"UniProxy/proxy"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/encoding/gjson"
	log "github.com/sirupsen/logrus"
)

type StartUniProxyRequest struct {
	Tag        string `json:"tag"`
	Uuid       string `json:"uuid"`
	GlobalMode bool   `json:"global_mode"`
}

func StartUniProxy2() {
	p := StartUniProxyRequest{
		Tag:        "shadowsocks_1",
		Uuid:       "72493186-abeb-479d-982c-b7dd7a0afc6d",
		GlobalMode: true,
	}
	proxy.GlobalMode = p.GlobalMode
	log.Println(gjson.MustEncodeString(p))
	log.Println(servers)
	err := proxy.StartProxy(p.Tag, p.Uuid, servers[p.Tag])
	if err != nil {
		log.Error("start proxy error: ", err)
		return
	}
}
func StartUniProxy(c *gin.Context) {
	p := StartUniProxyRequest{
		Tag:        "shadowsocks_1",
		Uuid:       "72493186-abeb-479d-982c-b7dd7a0afc6d",
		GlobalMode: true,
	}
	proxy.GlobalMode = p.GlobalMode
	log.Println(gjson.MustEncodeString(p))
	log.Println(servers)
	err := proxy.StartProxy(p.Tag, p.Uuid, servers[p.Tag])
	if err != nil {
		log.Error("start proxy error: ", err)
		c.JSON(400, Rsp{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, Rsp{
		Success: true,
		Message: "ok",
		Data: StatusData{
			Inited:      inited,
			Running:     proxy.Running,
			GlobalMode:  proxy.GlobalMode,
			SystemProxy: proxy.SystemProxy,
		},
	})
}

func StopUniProxy(c *gin.Context) {
	if proxy.Running {
		proxy.StopProxy()
	}
	c.JSON(200, Rsp{
		Success: true,
		Message: "ok",
	})
}

func SetSystemProxy(c *gin.Context) {
	c.JSON(200, Rsp{
		Success: true,
		Message: "ok",
	})
}

func ClearSystemProxy(c *gin.Context) {
	err := proxy.ClearSystemProxy()
	if err != nil {
		c.JSON(200, Rsp{
			Success: false,
			Message: err.Error(),
		})
	}
	c.JSON(200, Rsp{
		Success: true,
		Message: "ok",
	})
}

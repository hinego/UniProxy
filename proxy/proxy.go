package proxy

import (
	"UniProxy/common/sysproxy"
	"UniProxy/v2b"
	"log"
	"net/netip"
	"time"

	box "github.com/sagernet/sing-box"
)

var (
	Running     bool
	SystemProxy bool
	GlobalMode  bool
	TunMode     bool = true
	InPort      int
	DataPath    string
	ResUrl      string
)

var client *box.Box

func StartProxy(tag string, uuid string, server *v2b.ServerInfo) error {
	if !Running {
		StopProxy()
	}
	SystemProxy = true
	c, err := GetSingBoxConfig(uuid, server)
	if err != nil {
		return err
	}
	client, err = box.New(box.Options{Options: c})
	if err != nil {
		return err
	}
	err = client.Start()
	if err != nil {
		return err
	}
	go func() {
		for {
			log.Println(client.Router().InterfaceMonitor().DefaultInterface(netip.MustParseAddr("1.1.1.1")))
			log.Println(client.Router().InterfaceMonitor().DefaultInterfaceName(netip.MustParseAddr("1.1.1.1")))
			log.Println(client.Router().InterfaceMonitor().DefaultInterfaceIndex(netip.MustParseAddr("1.1.1.1")))
			time.Sleep(2 * time.Second)
		}
	}()
	Running = true
	return nil
}

func StopProxy() {
	if Running {
		client.Close()
		Running = false
	}
}

func ClearSystemProxy() error {
	if Running {
		client.Close()
		Running = false
		return nil
	}
	sysproxy.ClearSystemProxy()
	return nil
}

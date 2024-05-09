package proxy

import (
	"UniProxy/common/file"
	"UniProxy/geo"
	"UniProxy/v2b"
	"errors"
	"fmt"
	"log"
	"net/netip"
	"os"
	"path"

	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/option"
	dns "github.com/sagernet/sing-dns"
)

func GetSingBoxConfig(uuid string, server *v2b.ServerInfo) (option.Options, error) {
	in := option.Inbound{
		Type: "tun",
		TunOptions: option.TunInboundOptions{
			Inet4Address: option.Listable[netip.Prefix]{
				netip.MustParsePrefix("172.19.0.1/24"),
			},
			InterfaceName: "KeyLink",
			MTU:           9000,
			AutoRoute:     true,
			// StrictRoute: true,
			Inet4RouteAddress: option.Listable[netip.Prefix]{
				netip.MustParsePrefix("0.0.0.0/1"),
				netip.MustParsePrefix("128.0.0.0/1"),
			},
			Stack: "gvisor",
			InboundOptions: option.InboundOptions{
				SniffEnabled:             true,
				SniffOverrideDestination: true,
				DomainStrategy:           option.DomainStrategy(dns.DomainStrategyPreferIPv4),
			},
		},
	}
	so := option.ServerOptions{
		Server:     server.Host,
		ServerPort: uint16(server.Port),
	}
	var out option.Outbound
	switch server.Type {
	case "shadowsocks":
		out = option.Outbound{
			Type: "shadowsocks",
			Tag:  "p",
			ShadowsocksOptions: option.ShadowsocksOutboundOptions{
				ServerOptions: so,
				Password:      uuid,
				Method:        server.Cipher,
			},
		}
	default:
		return option.Options{}, errors.New("server type is unknown")
	}
	log.Println(out)
	out = option.Outbound{
		Type: "shadowsocks",
		Tag:  "p",
		ShadowsocksOptions: option.ShadowsocksOutboundOptions{
			ServerOptions: option.ServerOptions{
				Server:     "120.232.240.44",
				ServerPort: 13124,
			},
			Password: "123456",
			Method:   "chacha20-ietf-poly1305",
			UDPOverTCP: &option.UDPOverTCPOptions{
				Enabled: false,
			},
		},
	}
	r, err := getRules(GlobalMode)
	if err != nil {
		return option.Options{}, fmt.Errorf("get rules error: %s", err)
	}
	return option.Options{
		DNS: &option.DNSOptions{
			Servers: []option.DNSServerOptions{
				{
					Tag:      "dns_proxy",
					Address:  "1.1.1.1",
					Strategy: option.DomainStrategy(dns.DomainStrategyPreferIPv4),
					Detour:   "p",
				},
			},
			Rules: []option.DNSRule{
				{
					DefaultOptions: option.DefaultDNSRule{
						Outbound: []string{
							"any",
						},
						Server: "dns_proxy",
					},
				},
			},
		},
		Log: &option.LogOptions{
			Output: path.Join(DataPath, "proxy.log"),
		},
		Inbounds: []option.Inbound{
			in,
		},
		Outbounds: []option.Outbound{
			out,
			{
				Tag:  "dns",
				Type: "dns",
			},
		},
		Route: r,
	}, nil
}

func getRules(global bool) (*option.RouteOptions, error) {
	var r = option.RouteOptions{
		Rules: []option.Rule{
			{
				Type: C.RuleTypeDefault,
				DefaultOptions: option.DefaultRule{
					Outbound: "dns",
					Protocol: []string{
						"dns",
					},
				},
			},
		},
		Final:               "p",
		AutoDetectInterface: true,
	}
	return &r, nil
}

func checkRes(p string) error {
	if !file.IsExist(path.Join(p, "geoip.dat")) {
		f, err := os.OpenFile(path.Join(p, "geoip.dat"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write(geo.Ip)
		if err != nil {
			return err
		}
	}
	if !file.IsExist(path.Join(p, "geosite.dat")) {
		f, err := os.OpenFile(path.Join(p, "geosite.dat"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write(geo.Site)
		if err != nil {
			return err
		}
	}
	return nil
}

package api

type (
	Node struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
	Region struct {
		Sort uint64  `json:"sort"`
		Name string  `json:"name"`
		Data []*Node ``
	}
	Version struct { // 每次打开App时，如果发现有新版本，就弹出提示框
		Current string // decimal 类型 例如： 3.66 只有一个小数点
		Latest  string // 同上
		Name    string // 新版本名称
		Desc    string // 新版说明
		Must    bool   // 是否强制更新 true:强制更新 false:可选更新
	}
	StatusRes struct {
		Connected bool      `json:"connected"` // 是否已链接
		Mode      string    `json:"mode"`      // global | rule
		NodeID    uint64    `json:"node_id"`   // 当前选择节点
		Data      []*Region `json:"node_data"` // 当前节点数据
	}
	NodeReq struct { // 更新选择的节点
		ID uint64
	}
	ModeReq struct { // 更新选择的模式
		Mode string
	}
	ConnectReq struct { // 创建链接
	}
	DisconnectReq struct { // 断开链接
	}
	RegisterReq struct { // 注册
		Username string
		Password string
	}
	LoginReq struct { // 登录
		Username string
		Password string
	}
)

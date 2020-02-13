package model

type PersonaInfo struct {
	BaseTagModel
	Introduction string `json:"introduction"`
	Facebook     string `json:"facebook"`
	Twitter      string `json:"twitter"`
	Weibo        string `json:"weibo"`
	Zhihu        string `json:"zhihu"`
	Wechat       string `json:"wechat"`
	Beian        string `json:"beian"`
}

func (p *PersonaInfo) TableName() string {
	return DBTableNamePersonalInfo
}

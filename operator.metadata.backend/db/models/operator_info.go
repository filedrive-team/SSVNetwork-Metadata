package models

type OperatorSignMsg struct {
	// base info
	Id           uint64 `gorm:"id;primaryKey" json:"operator_id"`
	Name         string `gorm:"name;type:varchar(30)" json:"name"`
	OwnerAddress string `gorm:"owner_address;type:varchar(42)" json:"owner_address"`

	// extern info
	Location        string `gorm:"location" json:"location"`
	CloudProvider   string `gorm:"cloud_provider" json:"cloud_provider"`
	Eth1NodeClient  string `gorm:"eth1_node_client" json:"eth1_node_client"`
	Eth2NodeClient  string `gorm:"eth2_node_client" json:"eth2_node_client"`
	Description     string `gorm:"description" json:"description"`
	WebsiteUrl      string `gorm:"website_url" json:"website_url"`
	TwitterUrl      string `gorm:"twitter_url" json:"twitter_url"`
	LinkedinUrl     string `gorm:"linkedin_url" json:"linkedin_url"`
	DiscordUrl      string `gorm:"discord_url" json:"discord_url"`
	TelegramUrl     string `gorm:"telegram_url" json:"telegram_url"`
	RelaysSupported string `gorm:"column:relays_supported" json:"relays_supported"`
	MevBostEnabled  bool   `gorm:"column:mev_bost_enabled" json:"mev_bost_enabled"`
	Logo            string `gorm:"logo" json:"logo"`
	Timestamp       int64  `gorm:"timestamp" json:"timestamp"`
}

type OperatorInfo struct {
	OperatorSignMsg

	Cid       string `gorm:"cid" json:"cid"`
	Signature string `gorm:"signature" json:"signature"`
} //@name OperatorInfo

type OperatorCid struct {
	Id  uint   `gorm:"id" json:"operator_id"`
	Cid string `gorm:"cid" json:"cid"`
}

func init() {
	RegisterModel(&OperatorInfo{})
}

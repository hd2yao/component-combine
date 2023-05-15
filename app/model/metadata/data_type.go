package metadata

type Attribute struct {
	TraitType string `bson:"trait_type"` // 组件部位
	Value     string `bson:"value"`      // 组件名称
}

// MetaData 结构体
type MetaData struct {
	Name       string      `bson:"name"`
	Image      string      `bson:"image"`
	Attributes []Attribute `bson:"attributes"`
	Rarity     string      `bson:"rarity"` // 不一定有
	//Description   string      `bson:"description"`
	//DescriptionEn string      `bson:"description_en"`
}

// CharacterDetail 结构体
type CharacterDetail struct {
	CreatedAt   string   `bson:"created_at"`   // 创建时间（时间戳）
	CharacterID string   `bson:"character_id"` // 对应人物的编号
	Metadata    MetaData `bson:"metadata"`
	//TokenId         string   `bson:"token_id"`
	//ContractAddress string   `bson:"contract_address"`
	//UpdatedAt       string   `bson:"updated_at"`
}

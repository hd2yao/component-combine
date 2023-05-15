package metadata

import (
	"component-combine/app/global/variable"
	"go.uber.org/zap"
	"strings"
	"time"

	"component-combine/app/model/generate_component_info"
)

// JavascriptISOString 定义 MongoDB 中的格式，将 golang 中的时间转成这种格式
var JavascriptISOString = "2006-01-02T15:04:05.999Z07:00"
var CreatedAt, UpdateAt time.Time

func createISOTime(createdAt time.Time) string {
	if h, err := time.ParseDuration("1h"); err != nil {
		variable.ZapLog.Error("createISOTime 出错:", zap.Error(err))
		return ""
	} else {
		createdAt = time.Now().Add(h * 8)
		createdAtStr := createdAt.UTC().Format(JavascriptISOString)
		return createdAtStr
	}
}

func CreateCharacterDetailFactory() *CharacterDetail {
	return &CharacterDetail{}
}

func getComponentInfo(code string) generate_component_info.ComponentsInfo {
	componentsInfo := generate_component_info.GetComponents(strings.TrimSuffix(code, ".png"))
	return componentsInfo
}

// CreateAttributes 构造一组 attributes
func createAttributes(code string) ([]Attribute, string, string) {
	compInfo := getComponentInfo(code)
	components := compInfo.GetComponents()
	//"background", "body", "face", "eye", "posthair", "ear", "clothes", "mouth", "head", "weapon"
	var traitType = []string{"background", "body", "face", "eye", "posthair", "ear", "clothes", "mouth", "head", "weapon"}
	attributes := make([]Attribute, 10)
	for i := 0; i < len(traitType); i++ {
		attributes[i].TraitType = traitType[i]
		attributes[i].Value = components[i]
	}
	return attributes, code, compInfo.Rarity
}

func (m *MetaData) CreateMetaData(code string, image string) {
	attributes, name, rarity := createAttributes(code)
	m.Name = name
	m.Image = image
	m.Attributes = attributes
	if rarity == "0" {
		m.Rarity = "N"
	} else if rarity == "1" {
		m.Rarity = "R"
	}
}

func (c *CharacterDetail) CreateCharacter(code string, imageURL string, characterId string) {
	var metadata MetaData
	metadata.CreateMetaData(code, imageURL)
	c.CreatedAt = createISOTime(CreatedAt)
	c.Metadata = metadata
	c.CharacterID = characterId
}

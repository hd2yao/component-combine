package generate_component_info

// 稀有度编号
const (
	N = "0"
	R = "1"
)

// 阵营编号
const (
	Neutral = "00" // 中立阵营
	Shu     = "01" // 蜀国阵营
	Wei     = "10" // 魏国阵营
	Wu      = "11" // 吴国阵营
)

// 蜀国组件
var (
	BgNameShuN       = map[string]string{"01": "s-background"}
	BodyNameShuN     = map[string]string{"01": "s-body1", "02": "s-body2", "03": "s-body3", "04": "s-body4"}
	FaceNameShuN     = map[string]string{"00": "s-null", "01": "s-heart", "02": "s-net", "03": "s-patch", "04": "s-scar1"}
	PostHairNameShuN = map[string]string{"01": "s-curls1", "02": "s-curls2", "03": "s-curls3", "04": "s-curls4", "05": "s-curls5",
		"06": "s-curls6", "07": "s-curls7", "08": "s-curls8", "09": "s-curls9", "10": "s-curls10", "11": "s-curls11", "12": "s-curls12",
		"13": "s-darkblue1", "14": "s-darkblue2", "15": "s-darkblue3", "16": "s-darkblue4", "17": "s-darkblue5", "18": "s-darkblue6", "19": "s-darkblue7",
		"20": "s-hairpin1", "21": "s-hairpin2", "22": "s-ponytail",
		"23": "s-mushroom1", "24": "s-mushroom2", "25": "s-mushroom3", "26": "s-mushroom4",
		"27": "s-princesscut1", "28": "s-princesscut2", "29": "s-princesscut3", "30": "s-princesscut4", "31": "s-princesscut5",
		"32": "s-princesscut6", "33": "s-princesscut7", "34": "s-princesscut8", "35": "s-princesscut9", "36": "s-princesscut10",
		"37": "s-princesscut11", "38": "s-princesscut12", "39": "s-princesscut13", "40": "s-princesscut14", "41": "s-princesscut15", "42": "s-princesscut16",
		"43": "s-unilateralhair1", "44": "s-unilateralhair2", "45": "s-unilateralhair3", "46": "s-unilateralhair4",
		"47": "s-unilateralhair5", "48": "s-unilateralhair6", "49": "s-unilateralhair7", "50": "s-unilateralhair8",
		"51": "s-whitelonghair1", "52": "s-whitelonghair2", "53": "s-whitelonghair3", "54": "s-whitelonghair4"}
	EarNameShuN = map[string]string{"00": "s-null", "01": "s-earrings", "02": "s-dangler", "03": "s-earphone", "04": "s-earring1",
		"05": "s-earring2", "06": "s-earstuds"}
	EyeNameShuN     = map[string]string{"00": "s-null", "01": "s-eyemask2", "02": "s-glasses1", "03": "s-glasses2", "04": "s-laser"}
	ClothesNameShuN = map[string]string{"01": "s-bandage", "02": "s-cheongsam", "03": "s-combat", "04": "s-eveninggown",
		"05": "s-jacket", "06": "s-mechanicalarmor", "07": "s-mechanicaljacket", "08": "s-medical", "09": "s-powersupply",
		"10": "s-scoutsuit", "11": "s-strapless", "12": "s-suspenders", "13": "s-vest"}
	MouthNameShuN      = map[string]string{"00": "s-null", "01": "s-ironjaw", "02": "s-medicalmask", "03": "s-transparentmask", "04": "s-veil"}
	FormerhairNameShuN = map[string]string{"00": "s-null", "01": "s-darkblue1234567",
		"02": "s-princesscut1", "03": "s-princesscut5", "04": "s-princesscut8", "05": "s-princesscut11", "06": "s-princesscut14",
		"07": "s-princesscut67", "08": "s-princesscut234", "09": "s-princesscut910", "10": "s-princesscut1213", "11": "s-princesscut1516",
		"12": "s-whitelonghair1234"}
	HeadNameShuN = map[string]string{"00": "s-null", "01": "s-hat",
		"02": "s-beret+curls1234", "03": "s-beret+princesscut1581114", "04": "s-beret+unilateralhair12", "05": "s-beret+whitelonghair1",
		"06": "s-beret+curls+unilateralhair", "07": "s-beret+princesscut", "08": "s-beret+whitelonghair"}
	WeaponNameShuN = map[string]string{"00": "w-null"}
)

var (
	BgNameShuR       = map[string]string{"01": "s-background"}
	BodyNameShuR     = map[string]string{"01": "s-body1", "02": "s-body2", "03": "s-body4"}
	FaceNameShuR     = map[string]string{"00": "s-null", "01": "s-heart", "02": "s-net", "03": "s-patch", "04": "s-scar1"}
	PostHairNameShuR = map[string]string{"01": "s-curls1", "02": "s-curls2", "03": "s-curls3", "04": "s-curls4", "05": "s-curls5",
		"06": "s-curls6", "07": "s-curls7", "09": "s-curls9", "10": "s-curls10", "11": "s-curls11",
		"13": "s-darkblue1", "14": "s-darkblue2", "16": "s-darkblue4", "17": "s-darkblue5", "19": "s-darkblue7",
		"21": "s-hairpin2",
		"23": "s-mushroom1", "24": "s-mushroom2", "25": "s-mushroom3", "26": "s-mushroom4",
		"27": "s-princesscut1", "28": "s-princesscut2", "30": "s-princesscut4", "31": "s-princesscut5",
		"32": "s-princesscut6", "35": "s-princesscut9", "36": "s-princesscut10",
		"37": "s-princesscut11", "39": "s-princesscut13", "40": "s-princesscut14", "41": "s-princesscut15", "42": "s-princesscut16",
		"43": "s-unilateralhair1", "44": "s-unilateralhair2", "48": "s-unilateralhair6", "49": "s-unilateralhair7", "50": "s-unilateralhair8",
		"51": "s-whitelonghair1", "53": "s-whitelonghair3", "54": "s-whitelonghair4"}
	EarNameShuR = map[string]string{"00": "s-null", "01": "s-earrings", "02": "s-dangler", "03": "s-earphone", "04": "s-earring1",
		"05": "s-earring2", "06": "s-earstuds"}
	EyeNameShuR = map[string]string{"00": "s-null", "01": "s-eyemask2", "02": "s-glasses1", "03": "s-glasses2",
		"04": "s-laser", "05": "s-mask"}
	ClothesNameShuR = map[string]string{"01": "s-combat", "02": "s-jacket", "03": "s-mechanicalarmor", "04": "s-mechanicaljacket",
		"05": "s-medical", "06": "s-powersupply", "07": "s-scoutsuit", "08": "s-strapless", "09": "s-heavyarmor"}
	MouthNameShuR = map[string]string{"00": "s-null", "01": "s-ironjaw", "02": "s-medicalmask", "03": "s-transparentmask",
		"04": "s-veil", "05": "s-mask"}
	FormerhairNameShuR = map[string]string{"00": "s-null", "01": "s-darkblue1234567",
		"02": "s-princesscut1", "03": "s-princesscut5", "04": "s-princesscut8", "05": "s-princesscut11", "06": "s-princesscut14",
		"07": "s-princesscut67", "08": "s-princesscut234", "09": "s-princesscut910", "10": "s-princesscut1213", "11": "s-princesscut1516",
		"12": "s-whitelonghair1234"}
	HeadNameShuR = map[string]string{"00": "s-null", "01": "s-hat",
		"02": "s-beret+curls1234", "03": "s-beret+princesscut1581114", "04": "s-beret+unilateralhair12", "05": "s-beret+whitelonghair1",
		"06": "s-beret+curls+unilateralhair", "07": "s-beret+princesscut", "08": "s-beret+whitelonghair"}
	WeaponNameShuR = map[string]string{"01": "s-bow", "02": "s-doublegun", "03": "s-doublestick", "04": "s-gun",
		"05": "s-lute", "06": "s-mechanicalball", "07": "s-spellstaff1", "08": "s-spellstaff2", "09": "s-spellstaff3",
		"10": "s-sword", "11": "s-fan"}
)

// 魏国组件
var (
	BgNameWeiN   = map[string]string{"01": "w-background"}
	BodyNameWeiN = map[string]string{"01": "w-body1", "02": "w-body2", "03": "w-body3", "04": "w-body4", "05": "w-body5",
		"06": "w-body6", "07": "w-body7"}
	FaceNameWeiN = map[string]string{"00": "w-null", "01": "w-patch2", "02": "w-scar1", "03": "w-scar2"}
	EyeNameWeiN  = map[string]string{"00": "w-null", "01": "w-singlemechanical", "02": "w-eyemask1", "03": "w-eyemask2",
		"04": "w-glasses1", "05": "w-glasses2"}
	PostHairNameWeiN = map[string]string{"01": "w-bluelonghair1", "02": "w-bluelonghair2", "03": "w-bluelonghair3",
		"04": "w-bluelonghair4", "05": "w-bluelonghair5", "06": "w-bluelonghair6",
		"07": "w-blueponytail1", "08": "w-blueponytail2", "09": "w-blueponytail3",
		"15": "w-blueshort2", "16": "w-blueshort3", "17": "w-blueshort4",
		"18": "w-claret1", "19": "w-claret2", "21": "w-claret4",
		"22": "w-greylonghair1", "23": "w-greylonghair2", "24": "w-greylonghair3", "25": "w-greylonghair4", "26": "w-greylonghair5",
		"27": "w-mushroom1", "28": "w-mushroom2", "29": "w-mushroom3", "30": "w-mushroom4", "31": "w-mushroom5", "32": "w-mushroom6",
		"33": "w-mushroom7", "34": "w-mushroom8", "35": "w-mushroom9", "36": "w-mushroom10", "37": "w-mushroom11",
		"38": "w-princesscut1", "39": "w-princesscut2", "40": "w-princesscut3", "41": "w-princesscut4", "42": "w-princesscut5",
		"43": "w-princesscut6", "44": "w-princesscut7", "45": "w-princesscut8", "47": "w-princesscut10",
		"49": "w-princesscut12",
		"50": "w-whitelonghair1", "51": "w-whitelonghair2", "52": "w-whitelonghair3"}
	EarNameWeiN = map[string]string{"00": "w-null", "01": "w-earrings", "02": "w-dangler", "03": "w-earphone", "04": "w-earring1",
		"05": "w-earring2"}
	ClothesNameWeiN = map[string]string{"01": "w-bandage", "02": "w-biochemical", "03": "w-cheongsam", "04": "w-fighting",
		"05": "w-strapless", "06": "w-jacket", "07": "w-medical", "08": "w-reconnaissance", "09": "w-skullcoat", "10": "w-suspenders",
		"11": "w-vest", "12": "w-warframe"}
	MouthNameWeiN      = map[string]string{"00": "w-null", "01": "w-mask", "03": "w-transparentmask"}
	FormerhairNameWeiN = map[string]string{"00": "w-null", "01": "w-+bluelonghair", "02": "w-+bluelonghair23456",
		"03": "w-+greylonghair12345", "04": "w-+princesscut3", "05": "w-+princesscut6", "06": "w-+princesscut4789",
		"07": "w-+princesscut125101112", "08": "w-+whitelonghair123"}
	HeadNameWeiN   = map[string]string{"00": "w-null", "01": "w-hat"}
	WeaponNameWeiN = map[string]string{"00": "w-null"}
)
var (
	BgNameWeiR       = map[string]string{"01": "w-background"}
	BodyNameWeiR     = map[string]string{"01": "w-body1", "02": "w-body2", "03": "w-body3", "04": "w-zombie"}
	FaceNameWeiR     = map[string]string{"00": "w-null", "01": "w-patch2", "02": "w-net", "03": "w-sticker"}
	PostHairNameWeiR = map[string]string{"01": "w-bluelonghair1", "02": "w-bluelonghair6", "03": "w-blueponytail1",
		"04": "w-blueponytail2", "05": "w-greylonghair1", "06": "w-greylonghair2", "07": "w-greylonghair4",
		"08": "w-mushroom2", "09": "w-mushroom4", "10": "w-mushroom5", "11": "w-mushroom7", "12": "w-mushroom10",
		"13": "w-princesscut1", "14": "w-princesscut2", "15": "w-princesscut3", "16": "w-princesscut4", "17": "w-princesscut6",
		"18": "w-princesscut7", "19": "w-princesscut8", "20": "w-princesscut10", "21": "w-princesscut12",
		"22": "w-whitelonghair1", "23": "w-whitelonghair3", "24": "w-butterfly", "25": "w-curls1", "26": "w-curls2",
		"27": "w-unilateralhair", "28": "w-water", "29": "w-zombie"}
	EarNameWeiR = map[string]string{"00": "w-null", "01": "w-earrings", "02": "w-dangler", "03": "w-earphone", "04": "w-earring1",
		"05": "w-earring2", "06": "w-earstuds"}
	EyeNameWeiR = map[string]string{"00": "w-null", "01": "w-singlemechanical", "02": "w-glasses2", "03": "w-laser",
		"04": "w-mask", "05": "w-pupil"}
	ClothesNameWeiR = map[string]string{"01": "w-biochemical", "02": "w-fighting", "03": "w-strapless", "04": "w-jacket",
		"05": "w-medical", "06": "w-reconnaissance", "07": "w-skullcoat", "08": "w-warframe", "09": "w-heavyarmor",
		"10": "w-powersupply"}
	MouthNameWeiR      = map[string]string{"00": "w-null", "01": "w-mask", "02": "w-mechanicalmask", "03": "w-ironjaw"}
	FormerhairNameWeiR = map[string]string{"00": "w-null", "01": "w-+bluelonghair", "02": "w-+bluelonghair23456",
		"03": "w-+greylonghair12345", "04": "w-+princesscut3", "05": "w-+princesscut6", "06": "w-+princesscut4789",
		"07": "w-+princesscut125101112", "08": "w-+whitelonghair123", "09": "w-butterfly"}
	HeadNameWeiR   = map[string]string{"00": "w-null", "01": "w-hat"}
	WeaponNameWeiR = map[string]string{"01": "w-bow", "02": "w-broadsword", "03": "w-doublegun", "04": "w-gun",
		"05": "w-lute", "06": "w-magicsword", "07": "w-mechanicalball", "08": "w-mechanicalball2", "09": "w-spellstaff1",
		"10": "w-spellstaff2", "11": "w-spellstaff3", "12": "w-sword", "13": "w-claws", "14": "w-fan"}
)

// 吴国组件
var (
	BgNameWuN       = map[string]string{"01": "u-background"}
	BodyNameWuN     = map[string]string{"01": "u-body1", "02": "u-body2", "03": "u-body3", "04": "u-body4"}
	FaceNameWuN     = map[string]string{"00": "u-null", "01": "u-heart", "02": "u-net", "03": "u-patch", "04": "u-patch2", "05": "u-scar"}
	EyeNameWuN      = map[string]string{"00": "u-null", "01": "u-eyemask2", "02": "u-glasses1", "03": "u-laser"}
	PostHairNameWuN = map[string]string{"01": "u-bluelonghair1", "02": "u-bluelonghair2", "03": "u-bluelonghair3",
		"04": "u-blueshot1", "05": "u-blueshot2", "06": "u-blueshot3",
		"07": "u-brownlonghair1", "08": "u-brownlonghair2", "09": "u-brownlonghair3", "10": "u-brownlonghair4", "11": "u-brownlonghair5",
		"12": "u-claret1", "13": "u-claret2", "14": "u-claret3",
		"15": "u-curls0", "16": "u-curls1", "17": "u-curls2", "18": "u-curls3", "19": "u-curls4", "20": "u-curls5", "21": "u-curls6",
		"22": "u-curls7", "23": "u-curls8", "24": "u-curls9", "25": "u-curls10", "26": "u-curls11", "27": "u-curls12",
		"28": "u-curls13", "29": "u-curls14", "30": "u-curls15",
		"31": "u-darkblue1", "32": "u-darkblue2", "33": "u-darkblue3", "34": "u-darkblue4", "35": "u-darkblue5", "36": "u-darkblue6",
		"37": "u-greylonghair1", "38": "u-greylonghair2", "39": "u-greylonghair3", "40": "u-greylonghair4",
		"41": "u-hairpin1", "42": "u-hairpin2",
		"43": "u-mushroom1", "44": "u-mushroom2", "45": "u-mushroom3", "46": "u-mushroom4", "47": "u-mushroom5", "48": "u-mushroom6",
		"49": "u-mushroom7",
		"50": "u-octopus", "51": "u-ponytail",
		"52": "u-princesscut1", "53": "u-princesscut2", "54": "u-princesscut3", "55": "u-princesscut4", "56": "u-princesscut5",
		"57": "u-princesscut6", "58": "u-princesscut7", "59": "u-princesscut8", "60": "u-princesscut9", "61": "u-princesscut10",
		"62": "u-princesscut11", "63": "u-princesscut12", "64": "u-princesscut13", "65": "u-princesscut14", "66": "u-princesscut15",
		"67": "u-princesscut16", "68": "u-princesscut17",
		"69": "u-redponytail1", "70": "u-redponytail2",
		"71": "u-unilateralhair1", "72": "u-unilateralhair2", "73": "u-unilateralhair3", "74": "u-unilateralhair4",
		"75": "u-unilateralhair5", "76": "u-unilateralhair6", "77": "u-unilateralhair7", "78": "u-unilateralhair8",
		"79": "u-whitelonghair1", "80": "u-whitelonghair2", "81": "u-whitelonghair3", "82": "u-whitelonghair4", "83": "u-whitelonghair5",
		"84": "u-whitelonghair6"}
	EarNameWuN = map[string]string{"00": "u-null", "01": "u-earrings", "02": "u-dangler", "03": "u-earring1", "04": "u-earring2",
		"05": "u-earstuds"}
	ClothesNameWuN = map[string]string{"01": "u-bandage", "02": "u-biochemical", "03": "u-cheongsam", "04": "u-combat",
		"05": "u-eveninggown", "06": "u-mechanicalarmor", "07": "u-mechanicaljacket", "08": "u-medical", "09": "u-powersupply",
		"10": "u-scoutsuit", "11": "u-strapless1", "12": "u-strapless2", "13": "u-suspenders", "14": "u-vest"}
	MouthNameWuN = map[string]string{"00": "u-null", "01": "u-ironjaw", "02": "u-ironjaw2", "03": "u-medicalmask",
		"04": "u-transparentmask", "05": "u-veil"}
	FormerhairNameWuN = map[string]string{"00": "u-null", "01": "u-bluelonghair123", "02": "u-brownlonghair12345",
		"03": "u-darkblue123456", "04": "u-greylonghair1234", "05": "u-princesscut1", "06": "u-princesscut234", "07": "u-princesscut5",
		"08": "u-princesscut67", "09": "u-princesscut8", "10": "u-princesscut910", "11": "u-princesscut11", "12": "u-princesscut1213",
		"13": "u-princesscut1416", "14": "u-whitelonghair123456"}
	HeadNameWuN = map[string]string{"00": "u-null", "01": "u-beret+curls01215", "02": "u-beret+darkblue1",
		"03": "u-beret+princesscut15811", "04": "u-beret+whitelonghair1"}
	WeaponNameWuN = map[string]string{"00": "u-null"}
)

var (
	BgNameWuR       = map[string]string{"01": "u-background"}
	BodyNameWuR     = map[string]string{"01": "u-body1", "02": "u-body2", "03": "u-body3"}
	FaceNameWuR     = map[string]string{"00": "u-null", "01": "u-heart", "02": "u-net", "03": "u-patch", "04": "u-patch2"}
	PostHairNameWuR = map[string]string{"01": "u-bluelonghair1",
		"04": "u-blueshot1",
		"07": "u-brownlonghair1", "08": "u-brownlonghair2", "09": "u-brownlonghair3",
		"12": "u-claret1",
		"16": "u-curls1", "17": "u-curls2", "18": "u-curls3", "20": "u-curls5", "26": "u-curls11", "30": "u-curls15",
		"31": "u-darkblue1", "32": "u-darkblue2", "35": "u-darkblue5",
		"37": "u-greylonghair1", "38": "u-greylonghair2", "39": "u-greylonghair3", "40": "u-greylonghair4",
		"42": "u-hairpin2",
		"43": "u-mushroom1", "46": "u-mushroom4", "47": "u-mushroom5", "48": "u-mushroom6",
		"50": "u-octopus",
		"52": "u-princesscut1", "53": "u-princesscut2", "54": "u-princesscut3", "55": "u-princesscut4", "56": "u-princesscut5",
		"58": "u-princesscut7", "60": "u-princesscut9", "61": "u-princesscut10",
		"62": "u-princesscut11", "63": "u-princesscut12", "64": "u-princesscut13", "66": "u-princesscut15",
		"67": "u-princesscut16", "68": "u-princesscut17",
		"69": "u-redponytail1",
		"71": "u-unilateralhair1", "74": "u-unilateralhair4",
		"75": "u-unilateralhair5", "76": "u-unilateralhair6",
		"79": "u-whitelonghair1", "80": "u-whitelonghair2", "82": "u-whitelonghair4"}
	EarNameWuR = map[string]string{"00": "u-null", "01": "u-earrings", "02": "u-dangler", "03": "u-earring1", "04": "u-earring2",
		"05": "u-earstuds"}
	EyeNameWuR     = map[string]string{"00": "u-null", "01": "u-eyemask2", "02": "u-glasses1", "03": "u-laser", "04": "u-mask"}
	ClothesNameWuR = map[string]string{"01": "u-biochemical", "02": "u-cheongsam", "03": "u-combat",
		"04": "u-mechanicalarmor", "05": "u-mechanicaljacket", "06": "u-medical", "07": "u-powersupply",
		"08": "u-scoutsuit", "09": "u-strapless2", "10": "u-heavyarmor"}
	MouthNameWuR      = map[string]string{"00": "u-null", "01": "u-ironjaw2", "02": "u-medicalmask", "03": "u-veil"}
	FormerhairNameWuR = map[string]string{"00": "u-null", "01": "u-bluelonghair123", "02": "u-brownlonghair12345",
		"03": "u-darkblue123456", "04": "u-greylonghair1234", "05": "u-princesscut1", "06": "u-princesscut234", "07": "u-princesscut5",
		"08": "u-princesscut67", "09": "u-princesscut8", "10": "u-princesscut910", "11": "u-princesscut11", "12": "u-princesscut1213",
		"13": "u-princesscut1416", "14": "u-whitelonghair123456"}
	HeadNameWuR = map[string]string{"00": "u-null", "01": "u-beret+curls01215", "02": "u-beret+darkblue1",
		"03": "u-beret+princesscut15811", "04": "u-beret+whitelonghair1"}
	WeaponNameWuR = map[string]string{"01": "u-bow", "02": "u-doublegun", "03": "u-doublestick", "04": "u-gun",
		"05": "u-mechanicalball", "06": "u-spellstaff1", "07": "u-spellstaff2", "08": "u-fan"}
)

// 中立组件
var (
	BgNameNeutralN      = map[string]string{}
	BodyNameNeutralN    = map[string]string{}
	IneyeNameNeutralN   = map[string]string{}
	ClothesNameNeutralN = map[string]string{}
	HeadNameNeutralN    = map[string]string{}
	EarNameNeutralN     = map[string]string{}
	MouthNameNeutralN   = map[string]string{}
	OuteyeNameNeutralN  = map[string]string{}

	BgNameNeutralR      = map[string]string{}
	BodyNameNeutralR    = map[string]string{}
	IneyeNameNeutralR   = map[string]string{}
	ClothesNameNeutralR = map[string]string{}
	HeadNameNeutralR    = map[string]string{}
	EarNameNeutralR     = map[string]string{}
	MouthNameNeutralR   = map[string]string{}
	OuteyeNameNeutralR  = map[string]string{}
)

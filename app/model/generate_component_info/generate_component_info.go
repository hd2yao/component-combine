package generate_component_info

import (
	"strings"

	"component-combine/app/global/consts"
)

type ComponentsInfo struct {
	Bg         []string
	Body       []string
	Face       []string
	Eye        []string
	Posthair   []string
	Ear        []string
	Clothes    []string
	Mouth      []string
	Formerhair []string
	Head       []string
	Weapon     []string
	Rarity     string
	Group      string
}

func CreateComponentsInfoFactory(componentsCombination []string, rarity, group string) *ComponentsInfo {
	return &ComponentsInfo{}
}

// 组件生成
func (comp *ComponentsInfo) backgroundGeneration() {
	comp.Bg = append(comp.Bg, "01")
	if comp.Rarity == N {
		if comp.Group == Shu {
			comp.Bg = append(comp.Bg, consts.ShuNBgFile+BgNameShuN[comp.Bg[0]]+".png")
		} else if comp.Group == Wei {
			comp.Bg = append(comp.Bg, consts.WeiNBgFile+BgNameWeiN[comp.Bg[0]]+".png")
		} else if comp.Group == Wu {
			comp.Bg = append(comp.Bg, consts.WuNBgFile+BgNameWuN[comp.Bg[0]]+".png")
		} else {
			comp.Bg = append(comp.Bg, consts.NeutralNBgFile+BgNameNeutralN[comp.Bg[0]]+".png")
		}
	} else {
		if comp.Group == Shu {
			comp.Bg = append(comp.Bg, consts.ShuRBgFile+BgNameShuR[comp.Bg[0]]+".png")
		} else if comp.Group == Wei {
			comp.Bg = append(comp.Bg, consts.WeiRBgFile+BgNameWeiR[comp.Bg[0]]+".png")
		} else if comp.Group == Wu {
			comp.Bg = append(comp.Bg, consts.WuRBgFile+BgNameWuR[comp.Bg[0]]+".png")
		} else {
			comp.Bg = append(comp.Bg, consts.NeutralRBgFile+BgNameNeutralR[comp.Bg[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) bodyGeneration(number string) {
	//no := []string{"01", "02", "03", "04"}
	//comp.Body = append(comp.Body, no[rand.Intn(len(no))])
	comp.Body = append(comp.Body, number)
	if comp.Rarity == N {
		if comp.Group == Shu {
			comp.Body = append(comp.Body, consts.ShuNBodyFile+BodyNameShuN[comp.Body[0]]+".png")
		} else if comp.Group == Wei {
			comp.Body = append(comp.Body, consts.WeiNBodyFile+BodyNameWeiN[comp.Body[0]]+".png")
		} else if comp.Group == Wu {
			comp.Body = append(comp.Body, consts.WuNBodyFile+BodyNameWuN["02"]+".png")
		} else {
			comp.Body = append(comp.Body, consts.NeutralNBodyFile+BodyNameNeutralN[comp.Body[0]]+".png")
		}
	} else {
		if comp.Group == Shu {
			comp.Body = append(comp.Body, consts.ShuRBodyFile+BodyNameShuR[comp.Body[0]]+".png")
		} else if comp.Group == Wei {
			comp.Body = append(comp.Body, consts.WeiRBodyFile+BodyNameWeiR[comp.Body[0]]+".png")
		} else if comp.Group == Wu {
			comp.Body = append(comp.Body, consts.WuRBodyFile+BodyNameWuR[comp.Body[0]]+".png")
		} else {
			comp.Body = append(comp.Body, consts.NeutralRBodyFile+BodyNameNeutralR[comp.Body[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) faceGeneration(number string) {
	//no := []string{"00", "01", "02", "03", "04"}
	//comp.Face = append(comp.Face, no[rand.Intn(len(no))])
	comp.Face = append(comp.Face, number)
	switch comp.Rarity {
	case N:
		switch comp.Group {
		case Shu:
			comp.Face = append(comp.Face, consts.ShuNFaceFile+FaceNameShuN[comp.Face[0]]+".png")
		case Wei:
			comp.Face = append(comp.Face, consts.WeiNFaceFile+FaceNameWeiN[comp.Face[0]]+".png")
		case Wu:
			comp.Face = append(comp.Face, consts.WuNFaceFile+FaceNameWuN[comp.Face[0]]+".png")
		}
	case R:
		switch comp.Group {
		case Shu:
			comp.Face = append(comp.Face, consts.ShuRFaceFile+FaceNameShuR[comp.Face[0]]+".png")
		case Wei:
			comp.Face = append(comp.Face, consts.WeiRFaceFile+FaceNameWeiR[comp.Face[0]]+".png")
		case Wu:
			comp.Face = append(comp.Face, consts.WuRFaceFile+FaceNameWuR[comp.Face[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) eyeGeneration(number string) {
	comp.Eye = append(comp.Eye, number)
	switch comp.Rarity {
	case N:
		switch comp.Group {
		case Shu:
			comp.Eye = append(comp.Eye, consts.ShuNEyeFile+EyeNameShuN[comp.Eye[0]]+".png")
		case Wei:
			comp.Eye = append(comp.Eye, consts.WeiNEyeFile+EyeNameWeiN[comp.Eye[0]]+".png")
		case Wu:
			comp.Eye = append(comp.Eye, consts.WuNEyeFile+EyeNameWuN[comp.Eye[0]]+".png")
		}
	case R:
		switch comp.Group {
		case Shu:
			comp.Eye = append(comp.Eye, consts.ShuREyeFile+EyeNameShuR[comp.Eye[0]]+".png")
		case Wei:
			comp.Eye = append(comp.Eye, consts.WeiREyeFile+EyeNameWeiR[comp.Eye[0]]+".png")
		case Wu:
			comp.Eye = append(comp.Eye, consts.WuREyeFile+EyeNameWuR[comp.Eye[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) earGeneration(number string) {
	comp.Ear = append(comp.Ear, number)
	if comp.Rarity == N {
		if comp.Group == Shu {
			comp.Ear = append(comp.Ear, consts.ShuNEarFile+EarNameShuN[comp.Ear[0]]+".png")
		} else if comp.Group == Wei {
			comp.Ear = append(comp.Ear, consts.WeiNEarFile+EarNameWeiN[comp.Ear[0]]+".png")
		} else if comp.Group == Wu {
			comp.Ear = append(comp.Ear, consts.WuNEarFile+EarNameWuN[comp.Ear[0]]+".png")
		} else {
			comp.Ear = append(comp.Ear, consts.NeutralNEarFile+EarNameNeutralN[comp.Ear[0]]+".png")
		}
	} else {
		if comp.Group == Shu {
			comp.Ear = append(comp.Ear, consts.ShuREarFile+EarNameShuR[comp.Ear[0]]+".png")
		} else if comp.Group == Wei {
			comp.Ear = append(comp.Ear, consts.WeiREarFile+EarNameWeiR[comp.Ear[0]]+".png")
		} else if comp.Group == Wu {
			comp.Ear = append(comp.Ear, consts.WuREarFile+EarNameWuR[comp.Ear[0]]+".png")
		} else {
			comp.Ear = append(comp.Ear, consts.NeutralREarFile+EarNameNeutralR[comp.Ear[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) posthairGeneration(number string) {
	comp.Posthair = append(comp.Posthair, number)
	switch comp.Rarity {
	case N:
		switch comp.Group {
		case Shu:
			comp.Posthair = append(comp.Posthair, consts.ShuNPosthairFile+PostHairNameShuN[comp.Posthair[0]]+".png")
		case Wei:
			comp.Posthair = append(comp.Posthair, consts.WeiNPosthairFile+PostHairNameWeiN[comp.Posthair[0]]+".png")
		case Wu:
			comp.Posthair = append(comp.Posthair, consts.WuNPosthairFile+PostHairNameWuN[comp.Posthair[0]]+".png")
		}
	case R:
		switch comp.Group {
		case Shu:
			comp.Posthair = append(comp.Posthair, consts.ShuNPosthairFile+PostHairNameShuR[comp.Posthair[0]]+".png")
		case Wei:
			comp.Posthair = append(comp.Posthair, consts.WeiRPosthairFile+PostHairNameWeiR[comp.Posthair[0]]+".png")
		case Wu:
			comp.Posthair = append(comp.Posthair, consts.WuRPosthairFile+PostHairNameWuR[comp.Posthair[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) clothesGeneration(number string) {
	comp.Clothes = append(comp.Clothes, number)
	if comp.Rarity == N {
		if comp.Group == Shu {
			comp.Clothes = append(comp.Clothes, consts.ShuNClothesFile+ClothesNameShuN[comp.Clothes[0]]+".png")
		} else if comp.Group == Wei {
			comp.Clothes = append(comp.Clothes, consts.WeiNClothesFile+ClothesNameWeiN[comp.Clothes[0]]+".png")
		} else if comp.Group == Wu {
			comp.Clothes = append(comp.Clothes, consts.WuNClothesFile+ClothesNameWuN[comp.Clothes[0]]+".png")
		} else {
			comp.Clothes = append(comp.Clothes, consts.NeutralNClothesFile+ClothesNameNeutralN[comp.Clothes[0]]+".png")
		}
	} else {
		if comp.Group == Shu {
			comp.Clothes = append(comp.Clothes, consts.ShuRClothesFile+ClothesNameShuR[comp.Clothes[0]]+".png")
		} else if comp.Group == Wei {
			comp.Clothes = append(comp.Clothes, consts.WeiRClothesFile+ClothesNameWeiR[comp.Clothes[0]]+".png")
		} else if comp.Group == Wu {
			comp.Clothes = append(comp.Clothes, consts.WuRClothesFile+ClothesNameWuR[comp.Clothes[0]]+".png")
		} else {
			comp.Clothes = append(comp.Clothes, consts.NeutralRClothesFile+ClothesNameNeutralR[comp.Clothes[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) mouthGeneration(number string) {
	comp.Mouth = append(comp.Mouth, number)
	if comp.Rarity == N {
		if comp.Group == Shu {
			comp.Mouth = append(comp.Mouth, consts.ShuNMouthFile+MouthNameShuN[comp.Mouth[0]]+".png")
		} else if comp.Group == Wei {
			comp.Mouth = append(comp.Mouth, consts.WeiNMouthFile+MouthNameWeiN[comp.Mouth[0]]+".png")
		} else if comp.Group == Wu {
			comp.Mouth = append(comp.Mouth, consts.WuNMouthFile+MouthNameWuN[comp.Mouth[0]]+".png")
		} else {
			comp.Mouth = append(comp.Mouth, consts.NeutralNMouthFile+MouthNameShuN[comp.Mouth[0]]+".png")
		}
	} else {
		if comp.Group == Shu {
			comp.Mouth = append(comp.Mouth, consts.ShuRMouthFile+MouthNameShuR[comp.Mouth[0]]+".png")
		} else if comp.Group == Wei {
			comp.Mouth = append(comp.Mouth, consts.WeiRMouthFile+MouthNameWeiR[comp.Mouth[0]]+".png")
		} else if comp.Group == Wu {
			comp.Mouth = append(comp.Mouth, consts.WuRMouthFile+MouthNameWuR[comp.Mouth[0]]+".png")
		} else {
			comp.Mouth = append(comp.Mouth, consts.NeutralRMouthFile+MouthNameShuR[comp.Mouth[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) headGeneration() {
	comp.Head = append(comp.Head, "00")
	if comp.Rarity == N {
		if comp.Group == Shu {
			comp.Head = append(comp.Head, consts.ShuNHeadFile+HeadNameShuN[comp.Head[0]]+".png")
		} else if comp.Group == Wei {
			comp.Head = append(comp.Head, consts.WeiNHeadFile+HeadNameWeiN[comp.Head[0]]+".png")
		} else if comp.Group == Wu {
			comp.Head = append(comp.Head, consts.WuNHeadFile+HeadNameWuN[comp.Head[0]]+".png")
		} else {
			comp.Head = append(comp.Head, consts.NeutralNHeadFile+HeadNameNeutralN[comp.Head[0]]+".png")
		}
	} else {
		if comp.Group == Shu {
			comp.Head = append(comp.Head, consts.ShuRHeadFile+HeadNameShuR[comp.Head[0]]+".png")
		} else if comp.Group == Wei {
			comp.Head = append(comp.Head, consts.WeiRHeadFile+HeadNameWeiR[comp.Head[0]]+".png")
		} else if comp.Group == Wu {
			comp.Head = append(comp.Head, consts.WuRHeadFile+HeadNameWuR[comp.Head[0]]+".png")
		} else {
			comp.Head = append(comp.Head, consts.NeutralRHeadFile+HeadNameNeutralR[comp.Head[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) formerhairGeneration() {
	comp.Formerhair = append(comp.Formerhair, "00")
	if comp.Rarity == N {
		if comp.Group == Shu {
			comp.Formerhair = append(comp.Formerhair, consts.ShuNFormerhairFile+FormerhairNameShuN[comp.Formerhair[0]]+".png")
		} else if comp.Group == Wei {
			comp.Formerhair = append(comp.Formerhair, consts.WeiNFormerhairFile+FormerhairNameWeiN[comp.Formerhair[0]]+".png")
		} else if comp.Group == Wu {
			comp.Formerhair = append(comp.Formerhair, consts.WuNFormerhairFile+FormerhairNameWuN[comp.Formerhair[0]]+".png")
		} else {
			comp.Formerhair = append(comp.Formerhair, consts.NeutralNHeadFile+HeadNameNeutralN[comp.Formerhair[0]]+".png")
		}
	} else {
		if comp.Group == Shu {
			comp.Formerhair = append(comp.Formerhair, consts.ShuRFormerhairFile+FormerhairNameShuR[comp.Formerhair[0]]+".png")
		} else if comp.Group == Wei {
			comp.Formerhair = append(comp.Formerhair, consts.WeiRFormerhairFile+FormerhairNameWeiR[comp.Formerhair[0]]+".png")
		} else if comp.Group == Wu {
			comp.Formerhair = append(comp.Formerhair, consts.WuRFormerhairFile+FormerhairNameWuR[comp.Formerhair[0]]+".png")
		} else {
			comp.Formerhair = append(comp.Formerhair, consts.NeutralRHeadFile+HeadNameNeutralR[comp.Formerhair[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) weaponGeneration(number string) {
	comp.Weapon = append(comp.Weapon, number)
	if comp.Rarity == N {
		if comp.Group == Shu {
			comp.Weapon = append(comp.Weapon, consts.ShuNWeaponFile+WeaponNameShuN[comp.Weapon[0]]+".png")
		} else if comp.Group == Wei {
			comp.Weapon = append(comp.Weapon, consts.WeiNWeaponFile+WeaponNameWeiN[comp.Weapon[0]]+".png")
		} else if comp.Group == Wu {
			comp.Weapon = append(comp.Weapon, consts.WuNWeaponFile+WeaponNameWuN[comp.Weapon[0]]+".png")
		} else {
			comp.Weapon = append(comp.Weapon, consts.NeutralNMouthFile+WeaponNameShuN[comp.Weapon[0]]+".png")
		}
	} else {
		if comp.Group == Shu {
			switch number {
			case "11":
				comp.Weapon = append(comp.Weapon, consts.ShuRBehWeaponFile+WeaponNameShuR[comp.Weapon[0]]+".png")
			default:
				comp.Weapon = append(comp.Weapon, consts.ShuRFroWeaponFile+WeaponNameShuR[comp.Weapon[0]]+".png")
			}
		} else if comp.Group == Wei {
			switch number {
			case "13", "14":
				comp.Weapon = append(comp.Weapon, consts.WeiRBehWeaponFile+WeaponNameWeiR[comp.Weapon[0]]+".png")
			default:
				comp.Weapon = append(comp.Weapon, consts.WeiRFroWeaponFile+WeaponNameWeiR[comp.Weapon[0]]+".png")
			}
		} else if comp.Group == Wu {
			switch number {
			case "08":
				comp.Weapon = append(comp.Weapon, consts.WuRBehWeaponFile+WeaponNameWuR[comp.Weapon[0]]+".png")
			default:
				comp.Weapon = append(comp.Weapon, consts.WuRFroWeaponFile+WeaponNameWuR[comp.Weapon[0]]+".png")
			}
		} else {
			comp.Weapon = append(comp.Weapon, consts.NeutralNMouthFile+WeaponNameShuN[comp.Weapon[0]]+".png")
		}
	}
}

func (comp *ComponentsInfo) GetComponents() []string {
	var components []string
	//"background", "body", "face", "eye", "posthair", "ear", "clothes", "mouth", "head", "weapon"
	components = append(components, comp.Bg[1], comp.Body[1], comp.Face[1], comp.Eye[1], comp.Posthair[1], comp.Ear[1],
		comp.Clothes[1], comp.Mouth[1], comp.Head[1], comp.Weapon[1])
	return components
}

func (comp *ComponentsInfo) add() {
	var componentsName []map[string]string
	switch comp.Rarity {
	case N:
		switch comp.Group {
		case Shu:
			componentsName = append(componentsName, BgNameShuN, BodyNameShuN, FaceNameShuN, EyeNameShuN, PostHairNameShuN,
				EarNameShuN, ClothesNameShuN, MouthNameShuN, HeadNameShuN, WeaponNameShuN)
		case Wei:
			componentsName = append(componentsName, BgNameWeiN, BodyNameWeiN, FaceNameWeiN, EyeNameWeiN, PostHairNameWeiN,
				EarNameWeiN, ClothesNameWeiN, MouthNameWeiN, HeadNameWeiN, WeaponNameWeiN)
		case Wu:
			componentsName = append(componentsName, BgNameWuN, BodyNameWuN, FaceNameWuN, EyeNameWuN, PostHairNameWuN,
				EarNameWuN, ClothesNameWuN, MouthNameWuN, HeadNameWuN, WeaponNameWuN)
		}
	case R:
		switch comp.Group {
		case Shu:
			componentsName = append(componentsName, BgNameShuR, BodyNameShuR, FaceNameShuR, EyeNameShuR, PostHairNameShuR,
				EarNameShuR, ClothesNameShuR, MouthNameShuR, HeadNameShuR, WeaponNameShuR)
		case Wei:
			componentsName = append(componentsName, BgNameWeiR, BodyNameWeiR, FaceNameWeiR, EyeNameWeiR, PostHairNameWeiR,
				EarNameWeiR, ClothesNameWeiR, MouthNameWeiR, HeadNameWeiR, WeaponNameWeiR)
		case Wu:
			componentsName = append(componentsName, BgNameWuR, BodyNameWuR, FaceNameWuR, EyeNameWuR, PostHairNameWuR,
				EarNameWuR, ClothesNameWuR, MouthNameWuR, HeadNameWuR, WeaponNameWuR)
		}
	}
	// 去掉阵容前缀，背景除外
	comp.Bg = append(comp.Bg, componentsName[0][comp.Bg[0]])
	comp.Body = append(comp.Body, componentsName[1][comp.Body[0]][2:])
	comp.Face = append(comp.Face, componentsName[2][comp.Face[0]][2:])
	comp.Eye = append(comp.Eye, componentsName[3][comp.Eye[0]][2:])
	comp.Posthair = append(comp.Posthair, componentsName[4][comp.Posthair[0]][2:])
	comp.Ear = append(comp.Ear, componentsName[5][comp.Ear[0]][2:])
	comp.Clothes = append(comp.Clothes, componentsName[6][comp.Clothes[0]][2:])
	comp.Mouth = append(comp.Mouth, componentsName[7][comp.Mouth[0]][2:])
	if strings.Contains(componentsName[8][comp.Head[0]], "beret") {
		comp.Head = append(comp.Head, "beret")
	} else {
		comp.Head = append(comp.Head, componentsName[8][comp.Head[0]][2:])
	}
	comp.Weapon = append(comp.Weapon, componentsName[9][comp.Weapon[0]][2:])
}

// GetComponents 获取图片编码，并对应配件名称
func GetComponents(code string) ComponentsInfo {
	var componentsInfo ComponentsInfo

	componentsInfo.Rarity = string(code[20])
	componentsInfo.Group = code[21:]
	componentsInfo.Bg = append(componentsInfo.Bg, code[:2])
	componentsInfo.Body = append(componentsInfo.Body, code[2:4])
	componentsInfo.Face = append(componentsInfo.Face, code[4:6])
	componentsInfo.Eye = append(componentsInfo.Eye, code[6:8])
	componentsInfo.Posthair = append(componentsInfo.Posthair, code[8:10])
	componentsInfo.Ear = append(componentsInfo.Ear, code[10:12])
	componentsInfo.Clothes = append(componentsInfo.Clothes, code[12:14])
	componentsInfo.Mouth = append(componentsInfo.Mouth, code[14:16])
	componentsInfo.Head = append(componentsInfo.Head, code[16:18])
	componentsInfo.Weapon = append(componentsInfo.Weapon, code[18:20])

	componentsInfo.add()
	return componentsInfo
}

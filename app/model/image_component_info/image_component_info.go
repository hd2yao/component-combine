package image_component_info

type ImageComponentsInfo struct {
    Bg        string `bson:"background"`
    Body      string `bson:"body"`
    Eye       string `bson:"eye"`
    Clothes   string `bson:"clothes"`
    Head      string `bson:"head"`
    Ear       string `bson:"ear"`
    Mouth     string `bson:"mouth"`
    Weapon    string `bson:"weapon"`
    ImageName string `bson:"image_name"`
    ImageURL  string `bson:"image_url"`
}

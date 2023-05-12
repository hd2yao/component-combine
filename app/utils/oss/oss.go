package oss

import (
    "strings"
    "time"

    "ComponentsCombineFemale/env"

    "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func ConnectOSS(envInfo *env.EnvConfig) (*oss.Client, *oss.Bucket) {
    //envInfo := &env.EnvConfig{}
    envInfo.GetOssEnvConfig()

    // 创建OSSClient实例
    client, err := oss.New(envInfo.Oss.Endpoint, envInfo.Oss.AccessKeyId, envInfo.Oss.AccessKeySecret)
    if err != nil {
        //handleError(err)
        panic(err)
    }

    // 获取存储空间
    bucket, err := client.Bucket(envInfo.Oss.BucketName)
    if err != nil {
        //handleError(err)
        panic(err)
    }

    // https://bycoin.oss-cn-shanghai.aliyuncs.com/prod-t-nft/2020-08-10/character-nft/xxx
    // objectName := "prod-t-nft/2022-08-10/characterNFT/CryptoSanguo0077A.jpg"
    // localFileName := "D:\\GOLANG\\GoProject\\src\\Character-NFT\\resource\\CryptoSanguo0077A.jpg"
    // downloadedFileName := "C:\\Users\\10542\\Desktop\\characterNFT.jpg"

    return client, bucket
}

// FormatFileName 生成 objectName
func FormatFileName(jpgName string) (objectName string) {
    var fileName []string
    fileTime := time.Now().Format("2006-01-02")
    //jpgName = jpgName + ".jpg"
    fileName = append(fileName, "prod-t-nft", fileTime, "SanguoMaleN", jpgName)
    objectName = strings.Join(fileName, "/")
    return objectName // prod-t-nft/2022-11-04/SanguoMaleN/CryptoSanguo0077A.jpg
}

// UploadFromBucket 上传文件
func UploadFromBucket(bucket *oss.Bucket, objectName, localFileName string) {
    err := bucket.PutObjectFromFile(objectName, localFileName)
    if err != nil {
        //handleError(err)
        panic(err)
    }
}

// DownloadFromBucket 下载文件
func DownloadFromBucket(bucket *oss.Bucket, objectName, downloadedFileName string) {
    err := bucket.GetObjectToFile(objectName, downloadedFileName)
    if err != nil {
        //handleError(err)
        panic(err)
    }
}

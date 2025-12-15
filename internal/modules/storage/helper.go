package storage

import (
	"bytes"
	"context"
	"fmt"
	"image/png"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
)

func GenerateQr(client *minio.Client, qrData string) (string, error) {
	qrCode, _ := qr.Encode(qrData, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)
	buf := new(bytes.Buffer)
	png.Encode(buf, qrCode)
	var imageData = bytes.NewReader(buf.Bytes())

	uploadInfo, errUploadInfo := client.PutObject(context.Background(), "ebarimt-qr-images", time.Now().Format("20060102150405")+".png", imageData, imageData.Size(), minio.PutObjectOptions{ContentType: "image/png"})
	fmt.Println("uploadInfo :", uploadInfo, errUploadInfo)

	if errUploadInfo != nil {
		return "", errUploadInfo
	}

	return "http://" + viper.GetString("STORAGE_HOST") + "/" + uploadInfo.Bucket + "/" + uploadInfo.Key, nil
}

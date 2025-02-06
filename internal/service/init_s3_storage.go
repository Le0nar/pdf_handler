package service

import (
	"fmt"
	"log"

	"github.com/minio/minio-go"
)

const bucketName = "my-bucket"
const objectName = "my-object"

func initS3Storage() *minio.Client {
	// Параметры подключения
	endpoint := "localhost:9000"       // Адрес MinIO
	accessKeyID := "youraccesskey"     //  Access Key
	secretAccessKey := "yoursecretkey" //  Secret Key
	useSSL := false                    // SSL отключен (по умолчанию)

	// Создание клиента MinIO
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Проверяем, существует ли бакет
	found, err := client.BucketExists(bucketName)
	if err != nil {
		log.Fatalln(err)
	}

	// Если бакет не существует, создаем его
	if !found {
		err = client.MakeBucket(bucketName, "us-east-1") // Регион можно выбрать любой
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Bucket created successfully")
	}

	// Проверка подключения (получение информации о хранилище)
	buckets, err := client.ListBuckets()
	if err != nil {
		log.Fatalln(err)
	}

	// Вывод всех бакетов
	fmt.Println("Buckets:")
	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}

	return client
}

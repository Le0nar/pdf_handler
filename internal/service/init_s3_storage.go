package service

import (
	"fmt"
	"log"

	"github.com/minio/minio-go"
)

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

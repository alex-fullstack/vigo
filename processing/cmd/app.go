package main

import (
	"processing/pkg/infrastructure"
	"processing/pkg/services"
)

func main() {
	s3StorageCfg := infrastructure.NewS3StorageConfig(
		"vigo-storage",
		64*1024*1024,
		"ru-central1",
		"yc",
		"https://storage.yandexcloud.net",
		"YL",
		"Y6",
	)
	fileService := services.NewFileService(infrastructure.NewS3Storage(s3StorageCfg))

	natsBrokerCfg := infrastructure.NewNatsBrokerConfig(
		"nats://bus:4222",
		"EVENTS",
		256,
		"replay",
	)

	broker := infrastructure.NewNatsBroker(natsBrokerCfg, fileService)
	broker.Start()
}

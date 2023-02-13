package infrastructure

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3StorageConfig struct {
	Bucket           string
	DownloadPartSize int64
	Region           string
	PartitionId      string
	Url              string
	Key              string
	Secret           string
}

type S3Storage struct {
	bucket           string
	downloadPartSize int64
	endpoint         aws.Endpoint
	credentials      aws.Credentials
}

func NewS3Storage(cfg S3StorageConfig) *S3Storage {
	return &S3Storage{
		bucket:           cfg.Bucket,
		downloadPartSize: cfg.DownloadPartSize,
		endpoint: aws.Endpoint{
			PartitionID:   cfg.PartitionId,
			URL:           cfg.Url,
			SigningRegion: cfg.Region,
		},
		credentials: aws.Credentials{
			AccessKeyID:     cfg.Key,
			SecretAccessKey: cfg.Secret,
		},
	}
}

func NewS3StorageConfig(bucket string, downloadPartSize int64, region string, partitionId string, url string, key string, secret string) S3StorageConfig {
	return S3StorageConfig{
		Bucket:           bucket,
		DownloadPartSize: downloadPartSize,
		PartitionId:      partitionId,
		Url:              url,
		Region:           region,
		Key:              key,
		Secret:           secret,
	}
}

func (fs *S3Storage) Download(fileId string) (fileData []byte, err error) {
	client, err := fs.getClient()
	if err != nil {
		return
	}
	downloader := manager.NewDownloader(client, func(d *manager.Downloader) {
		d.PartSize = fs.downloadPartSize
	})
	w := manager.NewWriteAtBuffer([]byte{})
	_, err = downloader.Download(context.TODO(), w, &s3.GetObjectInput{
		Bucket: aws.String(fs.bucket),
		Key:    aws.String(fileId),
	})
	if err != nil {
		return
	}
	return w.Bytes(), nil
}

func (fs *S3Storage) getEndpointResolver() aws.EndpointResolverWithOptionsFunc {
	return func(service, region string, options ...interface{}) (endpoint aws.Endpoint, err error) {
		if service == s3.ServiceID && region == fs.endpoint.SigningRegion {
			endpoint = fs.endpoint
			return
		}
		err = fmt.Errorf("unknown endpoint requested")
		return
	}
}

func (fs *S3Storage) getClient() (client *s3.Client, err error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(fs.endpoint.SigningRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			fs.credentials.AccessKeyID,
			fs.credentials.SecretAccessKey,
			"",
		)),
		config.WithEndpointResolverWithOptions(fs.getEndpointResolver()),
	)
	if err != nil {
		return
	}

	client = s3.NewFromConfig(cfg)
	return
}

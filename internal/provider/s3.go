package provider

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	appconfig "github.com/regalangcom/go-shop-api/internal/config"
)

type S3Provider struct {
	client   *s3.Client
	Uploader *transfermanager.Client
	bucket   string
	endpoint string
}

func NewS3Provider(cfg *appconfig.Config) (*S3Provider, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.AWS.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AWS.AccessKeyID, cfg.AWS.SecretAccessKey, "")),
	)

	if err != nil {
		return nil, err
	}

	// configure the localstack endpoint if needed

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		if cfg.AWS.S3Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.AWS.S3Endpoint)
			o.UsePathStyle = true
		}
	})

	tm := transfermanager.New(client, func(o *transfermanager.Options) {
		o.PartSizeBytes = 64 * 1024 * 1024
		o.Concurrency = 3
	})

	return &S3Provider{
		Uploader: tm,
		client:   client,
		bucket:   cfg.AWS.S3Bucket,
		endpoint: cfg.AWS.S3Endpoint,
	}, nil

}

func (p *S3Provider) UploadFile(file *multipart.FileHeader, path string) (string, error) {

	log.Printf("Uploading file to S3: bucket=%s, key=%s, content-type=%s", p.bucket, path, file.Header.Get("Content-Type"))

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	// defer src.Close()
	defer func() {
		if cerr := src.Close(); cerr != nil {
			fmt.Println("failed to close source file:", cerr)
		}
	}()

	_, err = p.Uploader.UploadObject(context.TODO(), &transfermanager.UploadObjectInput{
		Bucket:      aws.String(p.bucket),
		Key:         aws.String(path),
		Body:        src,
		ContentType: aws.String(file.Header.Get("Content-Type")),
	})
	if err != nil {
		return "", err
	}

	return p.buildURL(path), nil
}

func (p *S3Provider) buildURL(path string) string {
	if p.endpoint != "" {
		return fmt.Sprintf("%s/%s/%s", p.endpoint, p.bucket, path)
	}
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", p.bucket, path)
}

func (p *S3Provider) DeleteFile(path string) error {
	_, err := p.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return err
	}

	return nil
}

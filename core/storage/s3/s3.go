package s3

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
	"sync"
)

type S3 struct {
	ep           string
	accessKeyId  string
	secretAccess string
	useSSL       bool

	Client *minio.Client
}

var (
	instance *S3 = nil
	once         = sync.Once{}
)

func New() *S3 {

	var err error = nil

	once.Do(func() {
		inst := S3{
			ep:           os.Getenv("S3_EP"),
			accessKeyId:  os.Getenv("S3_ACCESS_KEY_ID"),
			secretAccess: os.Getenv("S3_SECRET_ACCESS"),
			useSSL:       false,
			Client:       nil,
		}

		client, minio_err := minio.New(inst.ep, &minio.Options{
			Creds:  credentials.NewStaticV4(inst.accessKeyId, inst.secretAccess, ""),
			Secure: inst.useSSL,
		})
		err = minio_err
		inst.Client = client

		instance = &inst
	})

	if err != nil {
		log.Fatal("Cannot initialize S3 instance")
	}

	return instance
}

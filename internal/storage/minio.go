package storage

import (
	"github.com/minio/minio-go"
	"github.com/rs/zerolog/log"
	"io"
	"net/url"
	"path"
	"time"
)

const (
	MINIO_DEFAULT_EXPIRE_PRESIGNED = time.Hour * 24 * 7 //7 days
)

type FileStorage interface {
	GetPublicUrl(bucketName string, objectName string) string
	SignUrl(bucketName string, objectName string) (string, error)
	DownloadFile(bucketName string, objectName string, base string) error
	UploadData(bucketName string, objectName string, reader io.Reader, dataSize int64, acl string) error
	Upload(bucketName string, fileName string, objectName string, acl string) error
}
type Storage struct {
	minioClient    *minio.Client
	endPoint       string
	publicEndPoint string
}

var _ FileStorage = (*Storage)(nil)

func New(minioClient *minio.Client, endPoint string, publicEndPoint string) Storage {
	return Storage{minioClient: minioClient, endPoint: endPoint, publicEndPoint: publicEndPoint}
}

func (s Storage) GetPublicUrl(bucketName string, objectName string) string {
	publicURL, _ := url.Parse(s.publicEndPoint)
	publicURL.Path = path.Join(publicURL.Path, bucketName, objectName)
	return publicURL.String()
}

// SignUrl Generates a presigned URL for HTTP GET operations.
func (s Storage) SignUrl(bucketName string, objectName string) (string, error) {
	reqParams := make(url.Values)
	fileurl, err := s.minioClient.PresignedGetObject(bucketName, objectName, MINIO_DEFAULT_EXPIRE_PRESIGNED, reqParams)

	if err != nil {
		log.Printf("ERROR: Didn't manage to sign url for %s/%s", bucketName, objectName)
		return "", err
	}

	return fileurl.Host + fileurl.Path, nil
}

// DownloadFile Downloads and saves the object as a file in the local filesystem.
func (s Storage) DownloadFile(bucketName string, objectName string, base string) error {

	var err = s.minioClient.FGetObject(bucketName, objectName, base+"/"+objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("ERROR: Didn't manage to download file %s/%s", bucketName, objectName)
		return err
	}
	return err

}

func (s Storage) checkBuchet(bucketName string) error {
	exists, err := s.minioClient.BucketExists(bucketName)
	if err != nil {
		log.Printf("ERROR: Didn't manage to check the bucket %s", bucketName)
		return err
	}

	if !exists {
		if err = s.minioClient.MakeBucket(bucketName, ""); err != nil {
			log.Printf("ERROR: Didn't manage to create the bucket %s", bucketName)
			return err
		}
	}
	return nil
}

// UploadData Uploads contents from a reader to object
func (s Storage) UploadData(bucketName string, objectName string, reader io.Reader, dataSize int64, acl string) error {
	if err := s.checkBuchet(bucketName); err != nil {
		return err
	}

	putOptions := minio.PutObjectOptions{}
	if acl != "" {
		userMetaData := map[string]string{"x-amz-acl": acl}
		putOptions.UserMetadata = userMetaData
	}

	_, err := s.minioClient.PutObject(bucketName, objectName, reader, dataSize, putOptions)
	if err != nil {
		log.Printf("ERROR: Didn't manage to upload data %s/%s", bucketName, objectName)
	}
	return err
}

// Upload Uploads contents from a file to object
func (s Storage) Upload(bucketName string, fileName string, objectName string, acl string) error {
	if err := s.checkBuchet(bucketName); err != nil {
		return err
	}

	putOptions := minio.PutObjectOptions{}
	if acl != "" {
		userMetaData := map[string]string{"x-amz-acl": acl}
		putOptions.UserMetadata = userMetaData
	}
	_, err := s.minioClient.FPutObject(bucketName, objectName, fileName, putOptions)
	if err != nil {
		log.Printf("ERROR: Didn't manage to upload file %s/%s - %s", bucketName, objectName, err.Error())
		return err
	}
	return err
}

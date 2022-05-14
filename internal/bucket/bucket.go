package bucket

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/vimcolorschemes/search/internal/dotenv"
)

func Get(bucket string, key string) (string, error) {
	accessKeyId, exists := dotenv.Get("AWS_S3_ACCESS_KEY_ID")
	if !exists {
		return "", errors.New("AWS_S3_ACCESS_KEY_ID is missing")
	}

	secretAccessKey, exists := dotenv.Get("AWS_S3_SECRET_ACCESS_KEY")
	if !exists {
		return "", errors.New("AWS_S3_SECRET_ACCESS_KEY is missing")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(accessKeyId, secretAccessKey, ""),
	})
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	requestInput := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := svc.GetObject(requestInput)
	if err != nil {
		return "", err
	}
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", body), nil
}

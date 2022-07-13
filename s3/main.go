package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIAVLGKSUQHKQAGPLKL", "rv0b7cTKoUxfqEeFBvM0pVQqH4B6FYszPDsqTMjg", ""),
	})
	if err != nil {
		log.Fatal("NewSession(): ", err)
	}

	fmt.Println(sess)

	// /** download */
	// downloader := s3manager.NewDownloader(sess)

	// f, err := os.Create("from.tmp")
	// if err != nil {
	// 	log.Fatal("Create(): ", err)
	// }

	// downloader.Download(f, &s3.GetObjectInput{
	// 	Bucket: aws.String("leooooo-mpc"),
	// 	Key:    aws.String("1529888140439425/small-2020-01-01.txt"),
	// })

	// /** uploader */
	// uploader := s3manager.NewUploader(sess)

	// f, err := os.Open("from.tmp")
	// if err != nil {
	// 	log.Fatal("Open(): ", err)
	// }

	// uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String("leooooo-mpc"),
	// 	Key:    aws.String("1529888140439425/uploader.txt"),
	// 	Body:   f,
	// })
}

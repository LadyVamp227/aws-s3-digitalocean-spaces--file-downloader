package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func main() {

	// Defining the flags to access their values later in the program
	key := flag.String("key", "", "Access key")
	secret := flag.String("secret", "", "Your secret")
	endpoint := flag.String("endpoint", "", "The name of bucket endpoint with https")
	bucket := flag.String("bucket", "/", "The name of the bucket")
	region := flag.String("region", "", "Region of the bucket")
	fileName := flag.String("file", "", "The name of the file that you want to download")
	fileDestination := flag.String("destination", "", "The path and name of the file that you want to save the output")
	flag.Parse()

	// Configuring aws s3
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(*key, *secret, ""),
		Endpoint:    aws.String(*endpoint),
		Region:      aws.String(*region),
	}

	// The session validates your request and directs it to your specified endpoint using the AWS SDK.
	newSession, _ := session.NewSession(s3Config)

	// Configuring the downloader
	downloader := s3manager.NewDownloader(newSession, func(d *s3manager.Downloader) {
		d.PartSize = 64 * 1024 * 1024 // 64MB per part
		d.Concurrency = 6
	})

	// Creation of the destination file
	file, err := os.Create(*fileDestination)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Downloading the file from the bucket
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(*bucket),
			Key:    aws.String(*fileName),
		})
	if err != nil {
		fmt.Printf("Error in downloading from file: %v \n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Download completed", file.Name(), numBytes, "bytes")
}

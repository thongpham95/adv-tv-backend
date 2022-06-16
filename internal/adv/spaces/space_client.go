package spaces

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// NewSpaceClient return a new s3 client instance
func NewSpaceClient() *s3.S3 {
	// key := os.Getenv("ADV_SPACES_KEY")
	// secret := os.Getenv("ADV_SPACES_SECRET")
	// advSpaceEndpoint := os.Getenv("ADV_SPACE_ENDPOINT")
	// advSpaceRegion := os.Getenv("ADV_SPACE_REGION")

	// [ Dev enviroment only
	key := "2TDO5XPJ7DQP6GZ4JD3E"
	secret := "nAxNFEgocuFOVIvH8SIJE4SY3KaSHC/s6i/UiW4sFnA"
	advSpaceEndpoint := "https://sgp1.digitaloceanspaces.com"
	advSpaceRegion := "sgp1"
	// Dev enviroment only ]

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(advSpaceEndpoint),
		Region:      aws.String(advSpaceRegion),
	}
	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)
	fmt.Println("S3 client spinned up!")
	return s3Client
}

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
	key := "YERUV3CASNPOG7FQ2QPS"
	secret := "2xQ+uI7ha94qehWlW+UC7WpghL/lJWnmWHEpHkCA8OU"
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

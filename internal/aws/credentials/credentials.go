package credentials

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func NewFromKeys(accessKey, secretKey string, sessionToken string) *aws.CredentialsCache {
	return aws.NewCredentialsCache(
		credentials.NewStaticCredentialsProvider(
			accessKey,
			secretKey,
			sessionToken,
		),
	)
}

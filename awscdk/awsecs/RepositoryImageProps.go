package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// The properties for an image hosted in a public or private repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//
//   repositoryImageProps := &RepositoryImageProps{
//   	Credentials: secret,
//   }
//
type RepositoryImageProps struct {
	// The secret to expose to the container that contains the credentials for the image repository.
	//
	// The supported value is the full ARN of an AWS Secrets Manager secret.
	Credentials awssecretsmanager.ISecret `field:"optional" json:"credentials" yaml:"credentials"`
}


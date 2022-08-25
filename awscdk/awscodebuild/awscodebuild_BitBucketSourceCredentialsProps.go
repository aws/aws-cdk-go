package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties of {@link BitBucketSourceCredentials}.
//
// Example:
//   codebuild.NewBitBucketSourceCredentials(this, jsii.String("CodeBuildBitBucketCreds"), &bitBucketSourceCredentialsProps{
//   	username: awscdk.SecretValue.secretsManager(jsii.String("my-bitbucket-creds"), &secretsManagerSecretOptions{
//   		jsonField: jsii.String("username"),
//   	}),
//   	password: awscdk.SecretValue.secretsManager(jsii.String("my-bitbucket-creds"), &secretsManagerSecretOptions{
//   		jsonField: jsii.String("password"),
//   	}),
//   })
//
type BitBucketSourceCredentialsProps struct {
	// Your BitBucket application password.
	Password awscdk.SecretValue `field:"required" json:"password" yaml:"password"`
	// Your BitBucket username.
	Username awscdk.SecretValue `field:"required" json:"username" yaml:"username"`
}


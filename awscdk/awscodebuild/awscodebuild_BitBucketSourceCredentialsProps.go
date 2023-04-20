package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Construction properties of {@link BitBucketSourceCredentials}.
//
// Example:
//   codebuild.NewBitBucketSourceCredentials(this, jsii.String("CodeBuildBitBucketCreds"), &BitBucketSourceCredentialsProps{
//   	Username: awscdk.SecretValue_SecretsManager(jsii.String("my-bitbucket-creds"), &SecretsManagerSecretOptions{
//   		JsonField: jsii.String("username"),
//   	}),
//   	Password: awscdk.SecretValue_*SecretsManager(jsii.String("my-bitbucket-creds"), &SecretsManagerSecretOptions{
//   		JsonField: jsii.String("password"),
//   	}),
//   })
//
// Experimental.
type BitBucketSourceCredentialsProps struct {
	// Your BitBucket application password.
	// Experimental.
	Password awscdk.SecretValue `field:"required" json:"password" yaml:"password"`
	// Your BitBucket username.
	// Experimental.
	Username awscdk.SecretValue `field:"required" json:"username" yaml:"username"`
}


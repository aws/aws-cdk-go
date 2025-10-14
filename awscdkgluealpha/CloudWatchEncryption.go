package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// CloudWatch Logs encryption configuration.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &SecurityConfigurationProps{
//   	CloudWatchEncryption: &CloudWatchEncryption{
//   		Mode: glue.CloudWatchEncryptionMode_KMS,
//   	},
//   	JobBookmarksEncryption: &JobBookmarksEncryption{
//   		Mode: glue.JobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	S3Encryption: &S3Encryption{
//   		Mode: glue.S3EncryptionMode_KMS,
//   	},
//   })
//
// Experimental.
type CloudWatchEncryption struct {
	// Encryption mode.
	// Experimental.
	Mode CloudWatchEncryptionMode `field:"required" json:"mode" yaml:"mode"`
	// The KMS key to be used to encrypt the data.
	// Default: A key will be created if one is not provided.
	//
	// Experimental.
	KmsKey awskms.IKeyRef `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}


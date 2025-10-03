package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// S3 encryption configuration.
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
type S3Encryption struct {
	// Encryption mode.
	// Experimental.
	Mode S3EncryptionMode `field:"required" json:"mode" yaml:"mode"`
	// The KMS key to be used to encrypt the data.
	// Default: no kms key if mode = S3_MANAGED. A key will be created if one is not provided and mode = KMS.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}


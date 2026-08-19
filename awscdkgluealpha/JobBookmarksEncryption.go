package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Job bookmarks encryption configuration for a `SecurityConfiguration`.
//
// Job bookmarks support only client-side encryption with a KMS key.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &SecurityConfigurationProps{
//   	CloudWatchEncryption: glue.CloudWatchEncryption_Kms(),
//   	JobBookmarksEncryption: glue.JobBookmarksEncryption_ClientSideKms(),
//   	S3Encryption: glue.S3Encryption_Kms(),
//   })
//
// Experimental.
type JobBookmarksEncryption interface {
}

// The jsii proxy struct for JobBookmarksEncryption
type jsiiProxy_JobBookmarksEncryption struct {
	_ byte // padding
}

// Client-side encryption (CSE) with an AWS KMS key managed by the account owner.
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
//
// Experimental.
func JobBookmarksEncryption_ClientSideKms(kmsKey interfacesawskms.IKeyRef) JobBookmarksEncryption {
	_init_.Initialize()

	var returns JobBookmarksEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobBookmarksEncryption",
		"clientSideKms",
		[]interface{}{kmsKey},
		&returns,
	)

	return returns
}


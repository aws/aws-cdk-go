package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// S3 encryption configuration for a `SecurityConfiguration`.
//
// Use {@link S3Encryption.s3Managed} for SSE-S3 or {@link S3Encryption.kms} for
// SSE-KMS. Because these are separate factories, a KMS key can never be paired
// with S3-managed encryption.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &SecurityConfigurationProps{
//   	CloudWatchEncryption: glue.CloudWatchEncryption_Kms(),
//   	JobBookmarksEncryption: glue.JobBookmarksEncryption_ClientSideKms(),
//   	S3Encryption: glue.S3Encryption_Kms(),
//   })
//
// Experimental.
type S3Encryption interface {
}

// The jsii proxy struct for S3Encryption
type jsiiProxy_S3Encryption struct {
	_ byte // padding
}

// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html
//
// Experimental.
func S3Encryption_Kms(kmsKey interfacesawskms.IKeyRef) S3Encryption {
	_init_.Initialize()

	var returns S3Encryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3Encryption",
		"kms",
		[]interface{}{kmsKey},
		&returns,
	)

	return returns
}

// Server-side encryption (SSE) with an Amazon S3-managed key.
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html
//
// Experimental.
func S3Encryption_S3Managed() S3Encryption {
	_init_.Initialize()

	var returns S3Encryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3Encryption",
		"s3Managed",
		nil, // no parameters
		&returns,
	)

	return returns
}


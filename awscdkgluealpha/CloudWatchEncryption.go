package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// CloudWatch Logs encryption configuration for a `SecurityConfiguration`.
//
// CloudWatch Logs support only server-side encryption with a KMS key.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &SecurityConfigurationProps{
//   	CloudWatchEncryption: glue.CloudWatchEncryption_Kms(),
//   	JobBookmarksEncryption: glue.JobBookmarksEncryption_ClientSideKms(),
//   	S3Encryption: glue.S3Encryption_Kms(),
//   })
//
// Experimental.
type CloudWatchEncryption interface {
}

// The jsii proxy struct for CloudWatchEncryption
type jsiiProxy_CloudWatchEncryption struct {
	_ byte // padding
}

// Server-side encryption (SSE) with an AWS KMS key managed by the account owner.
// Experimental.
func CloudWatchEncryption_Kms(kmsKey interfacesawskms.IKeyRef) CloudWatchEncryption {
	_init_.Initialize()

	var returns CloudWatchEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.CloudWatchEncryption",
		"kms",
		[]interface{}{kmsKey},
		&returns,
	)

	return returns
}


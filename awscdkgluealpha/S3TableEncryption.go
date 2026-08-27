package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Server-side encryption for the S3 bucket that a managed `S3Table` creates.
//
// Applies only when the table manages its own bucket (via
// `S3TableStorage.managedBucket`). An existing bucket keeps whatever encryption
// it was created with.
//
// Example:
//   var myDatabase Database
//
//   glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
//   	Storage: glue.S3TableStorage_ManagedBucket(glue.S3TableEncryption_S3Managed()),
//   	// ...
//   	Database: myDatabase,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("col1"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   })
//
// Experimental.
type S3TableEncryption interface {
}

// The jsii proxy struct for S3TableEncryption
type jsiiProxy_S3TableEncryption struct {
	_ byte // padding
}

// Server-side encryption (SSE-KMS) with an AWS KMS key managed by the account owner.
// Experimental.
func S3TableEncryption_Kms(key awskms.IKey) S3TableEncryption {
	_init_.Initialize()

	var returns S3TableEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3TableEncryption",
		"kms",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Server-side encryption (SSE-KMS) with an AWS KMS key managed by the KMS service.
// Experimental.
func S3TableEncryption_KmsManaged() S3TableEncryption {
	_init_.Initialize()

	var returns S3TableEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3TableEncryption",
		"kmsManaged",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Server-side encryption (SSE-S3) with an Amazon S3-managed key.
// Experimental.
func S3TableEncryption_S3Managed() S3TableEncryption {
	_init_.Initialize()

	var returns S3TableEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3TableEncryption",
		"s3Managed",
		nil, // no parameters
		&returns,
	)

	return returns
}


package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Where an `S3Table` stores its data.
//
// The two paths are mutually exclusive: a managed bucket may specify its
// server-side encryption, while an existing bucket keeps its own encryption — so
// an encryption choice can never be paired with a bring-your-own bucket.
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
type S3TableStorage interface {
}

// The jsii proxy struct for S3TableStorage
type jsiiProxy_S3TableStorage struct {
	_ byte // padding
}

// Store the table's data in an existing bucket. CDK does not manage the bucket's encryption.
//
// The bucket can be one you don't own, imported with
// `Bucket.fromBucketArn()` or `Bucket.fromBucketAttributes()`. If that bucket
// is KMS-encrypted, import it with `Bucket.fromBucketAttributes()` and supply
// the `encryptionKey` attribute. Otherwise, CDK has no reference to the key,
// which means that `S3Table.grantRead()`/`grantWrite()` will correctly grant
// S3 access but silently skip the KMS permissions on the key. As a consequence,
// at runtime, reads and writes will fail with access denied on the key.
// Experimental.
func S3TableStorage_FromBucket(bucket awss3.IBucket) S3TableStorage {
	_init_.Initialize()

	if err := validateS3TableStorage_FromBucketParameters(bucket); err != nil {
		panic(err)
	}
	var returns S3TableStorage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3TableStorage",
		"fromBucket",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}

// Store the table's data in a bucket created and managed by the table.
// Default: - S3-managed (SSE-S3) encryption.
//
// Experimental.
func S3TableStorage_ManagedBucket(encryption S3TableEncryption) S3TableStorage {
	_init_.Initialize()

	var returns S3TableStorage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.S3TableStorage",
		"managedBucket",
		[]interface{}{encryption},
		&returns,
	)

	return returns
}


package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// An import source from an S3 object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3ImportSource := awscdk.Aws_cloudfront.NewS3ImportSource(bucket, jsii.String("key"))
//
type S3ImportSource interface {
	ImportSource
	// the S3 bucket that contains the data.
	Bucket() awss3.IBucket
	// the key within the S3 bucket that contains the data.
	Key() *string
}

// The jsii proxy struct for S3ImportSource
type jsiiProxy_S3ImportSource struct {
	jsiiProxy_ImportSource
}

func (j *jsiiProxy_S3ImportSource) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ImportSource) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}


func NewS3ImportSource(bucket awss3.IBucket, key *string) S3ImportSource {
	_init_.Initialize()

	if err := validateNewS3ImportSourceParameters(bucket, key); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ImportSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.S3ImportSource",
		[]interface{}{bucket, key},
		&j,
	)

	return &j
}

func NewS3ImportSource_Override(s S3ImportSource, bucket awss3.IBucket, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.S3ImportSource",
		[]interface{}{bucket, key},
		s,
	)
}

// An import source that exists as a local file.
func S3ImportSource_FromAsset(path *string, options *awss3assets.AssetOptions) ImportSource {
	_init_.Initialize()

	if err := validateS3ImportSource_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.S3ImportSource",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// An import source that exists as an object in an S3 bucket.
func S3ImportSource_FromBucket(bucket awss3.IBucket, key *string) ImportSource {
	_init_.Initialize()

	if err := validateS3ImportSource_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.S3ImportSource",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}


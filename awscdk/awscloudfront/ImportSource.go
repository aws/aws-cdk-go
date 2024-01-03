package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// The data to be imported to the key value store.
//
// Example:
//   store := cloudfront.NewKeyValueStore(this, jsii.String("KeyValueStore"), &KeyValueStoreProps{
//   	KeyValueStoreName: jsii.String("KeyValueStore"),
//   	Source: cloudfront.ImportSource_FromAsset(jsii.String("path-to-data.json")),
//   })
//
type ImportSource interface {
}

// The jsii proxy struct for ImportSource
type jsiiProxy_ImportSource struct {
	_ byte // padding
}

func NewImportSource_Override(i ImportSource) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.ImportSource",
		nil, // no parameters
		i,
	)
}

// An import source that exists as a local file.
func ImportSource_FromAsset(path *string, options *awss3assets.AssetOptions) ImportSource {
	_init_.Initialize()

	if err := validateImportSource_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ImportSource",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// An import source that exists as an object in an S3 bucket.
func ImportSource_FromBucket(bucket awss3.IBucket, key *string) ImportSource {
	_init_.Initialize()

	if err := validateImportSource_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns ImportSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ImportSource",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}


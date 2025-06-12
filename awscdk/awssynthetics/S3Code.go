package awssynthetics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// S3 bucket path to the code zip file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3Code := awscdk.Aws_synthetics.NewS3Code(bucket, jsii.String("key"), jsii.String("objectVersion"))
//
type S3Code interface {
	Code
	// Called when the canary is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(_scope constructs.Construct, _handler *string, _family RuntimeFamily, _runtimeName *string) *CodeConfig
}

// The jsii proxy struct for S3Code
type jsiiProxy_S3Code struct {
	jsiiProxy_Code
}

func NewS3Code(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateNewS3CodeParameters(bucket, key); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Code{}

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.S3Code",
		[]interface{}{bucket, key, objectVersion},
		&j,
	)

	return &j
}

func NewS3Code_Override(s S3Code, bucket awss3.IBucket, key *string, objectVersion *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.S3Code",
		[]interface{}{bucket, key, objectVersion},
		s,
	)
}

// Specify code from a local path.
//
// Path must include the folder structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `AssetCode` associated with the specified path.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
func S3Code_FromAsset(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateS3Code_FromAssetParameters(assetPath, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.S3Code",
		"fromAsset",
		[]interface{}{assetPath, options},
		&returns,
	)

	return returns
}

// Specify code from an s3 bucket.
//
// The object in the s3 bucket must be a .zip file that contains
// the structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `S3Code` associated with the specified S3 object.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
func S3Code_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateS3Code_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.S3Code",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Specify code inline.
//
// Returns: `InlineCode` with inline code.
func S3Code_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateS3Code_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.S3Code",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Code) Bind(_scope constructs.Construct, _handler *string, _family RuntimeFamily, _runtimeName *string) *CodeConfig {
	if err := s.validateBindParameters(_scope, _handler, _family); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _handler, _family, _runtimeName},
		&returns,
	)

	return returns
}


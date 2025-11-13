package awssynthetics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Canary code from an inline string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inlineCode := awscdk.Aws_synthetics.NewInlineCode(jsii.String("code"))
//
type InlineCode interface {
	Code
	// Called when the canary is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct, handler *string, family RuntimeFamily, runtimeName *string) *CodeConfig
}

// The jsii proxy struct for InlineCode
type jsiiProxy_InlineCode struct {
	jsiiProxy_Code
}

func NewInlineCode(code *string) InlineCode {
	_init_.Initialize()

	if err := validateNewInlineCodeParameters(code); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineCode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.InlineCode",
		[]interface{}{code},
		&j,
	)

	return &j
}

func NewInlineCode_Override(i InlineCode, code *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.InlineCode",
		[]interface{}{code},
		i,
	)
}

// Specify code from a local path.
//
// Path must include the folder structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `AssetCode` associated with the specified path.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
func InlineCode_FromAsset(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateInlineCode_FromAssetParameters(assetPath, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.InlineCode",
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
func InlineCode_FromBucket(bucket interfacesawss3.IBucketRef, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateInlineCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.InlineCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Specify code inline.
//
// Returns: `InlineCode` with inline code.
func InlineCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateInlineCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.InlineCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineCode) Bind(scope constructs.Construct, handler *string, family RuntimeFamily, runtimeName *string) *CodeConfig {
	if err := i.validateBindParameters(scope, handler, family); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, handler, family, runtimeName},
		&returns,
	)

	return returns
}


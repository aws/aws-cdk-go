package awssynthetics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// The code the canary should execute.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   canary := synthetics.NewCanary(this, jsii.String("MyCanary"), &CanaryProps{
//   	Schedule: synthetics.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(5))),
//   	Test: synthetics.Test_Custom(&CustomTestOptions{
//   		Code: synthetics.Code_FromAsset(path.join(__dirname, jsii.String("canary"))),
//   		Handler: jsii.String("index.handler"),
//   	}),
//   	Runtime: synthetics.Runtime_SYNTHETICS_NODEJS_PUPPETEER_6_2(),
//   	Memory: cdk.Size_Mebibytes(jsii.Number(1024)),
//   })
//
type Code interface {
	// Called when the canary is initialized to allow this object to bind to the stack, add resources and have fun.
	//
	// Returns: a bound `CodeConfig`.
	Bind(scope constructs.Construct, handler *string, family RuntimeFamily) *CodeConfig
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_synthetics.Code",
		nil, // no parameters
		c,
	)
}

// Specify code from a local path.
//
// Path must include the folder structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `AssetCode` associated with the specified path.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
func Code_FromAsset(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateCode_FromAssetParameters(assetPath, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.Code",
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
func Code_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.Code",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Specify code inline.
//
// Returns: `InlineCode` with inline code.
func Code_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.Code",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Code) Bind(scope constructs.Construct, handler *string, family RuntimeFamily) *CodeConfig {
	if err := c.validateBindParameters(scope, handler, family); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, handler, family},
		&returns,
	)

	return returns
}


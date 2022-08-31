package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

// Represents a Docker image in ECR that can be bound as Lambda Code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var repository repository
//
//   ecrImageCode := awscdk.Aws_lambda.NewEcrImageCode(repository, &ecrImageCodeProps{
//   	cmd: []*string{
//   		jsii.String("cmd"),
//   	},
//   	entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	tag: jsii.String("tag"),
//   	tagOrDigest: jsii.String("tagOrDigest"),
//   	workingDirectory: jsii.String("workingDirectory"),
//   })
//
// Experimental.
type EcrImageCode interface {
	Code
	// Determines whether this Code is inline code or not.
	// Experimental.
	IsInline() *bool
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	// Experimental.
	Bind(_arg awscdk.Construct) *CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	// Experimental.
	BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions)
}

// The jsii proxy struct for EcrImageCode
type jsiiProxy_EcrImageCode struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_EcrImageCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


// Experimental.
func NewEcrImageCode(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	j := jsiiProxy_EcrImageCode{}

	_jsii_.Create(
		"monocdk.aws_lambda.EcrImageCode",
		[]interface{}{repository, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcrImageCode_Override(e EcrImageCode, repository awsecr.IRepository, props *EcrImageCodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda.EcrImageCode",
		[]interface{}{repository, props},
		e,
	)
}

// DEPRECATED.
// Deprecated: use `fromAsset`.
func EcrImageCode_Asset(path *string) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"asset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `fromBucket`.
func EcrImageCode_Bucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"bucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `fromCfnParameters`.
func EcrImageCode_CfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"cfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from a local disk path.
// Experimental.
func EcrImageCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
// Experimental.
func EcrImageCode_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
// Experimental.
func EcrImageCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`.
// Experimental.
func EcrImageCode_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
// Experimental.
func EcrImageCode_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
// Experimental.
func EcrImageCode_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
// Experimental.
func EcrImageCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `fromInline`.
func EcrImageCode_Inline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.EcrImageCode",
		"inline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrImageCode) Bind(_arg awscdk.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrImageCode) BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions) {
	_jsii_.InvokeVoid(
		e,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}


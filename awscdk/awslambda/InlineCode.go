package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Lambda code from an inline string.
//
// Example:
//   layer := lambda.NewLayerVersion(stack, jsii.String("MyLayer"), &LayerVersionProps{
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("layer-code"))),
//   	CompatibleRuntimes: []runtime{
//   		lambda.*runtime_NODEJS_LATEST(),
//   	},
//   	License: jsii.String("Apache-2.0"),
//   	Description: jsii.String("A layer to test the L2 construct"),
//   })
//
//   // To grant usage by other AWS accounts
//   layer.addPermission(jsii.String("remote-account-grant"), &LayerVersionPermission{
//   	AccountId: awsAccountId,
//   })
//
//   // To grant usage to all accounts in some AWS Ogranization
//   // layer.grantUsage({ accountId: '*', organizationId });
//
//   // To grant usage to all accounts in some AWS Ogranization
//   // layer.grantUsage({ accountId: '*', organizationId });
//   lambda.NewFunction(stack, jsii.String("MyLayeredLambda"), &FunctionProps{
//   	Code: lambda.NewInlineCode(jsii.String("foo")),
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.*runtime_NODEJS_LATEST(),
//   	Layers: []iLayerVersion{
//   		layer,
//   	},
//   })
//
type InlineCode interface {
	Code
	// Determines whether this Code is inline code or not.
	IsInline() *bool
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(_scope constructs.Construct) *CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions)
}

// The jsii proxy struct for InlineCode
type jsiiProxy_InlineCode struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_InlineCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


func NewInlineCode(code *string) InlineCode {
	_init_.Initialize()

	if err := validateNewInlineCodeParameters(code); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineCode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.InlineCode",
		[]interface{}{code},
		&j,
	)

	return &j
}

func NewInlineCode_Override(i InlineCode, code *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.InlineCode",
		[]interface{}{code},
		i,
	)
}

// Loads the function code from a local disk path.
func InlineCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateInlineCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func InlineCode_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	if err := validateInlineCode_FromAssetImageParameters(directory, props); err != nil {
		panic(err)
	}
	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func InlineCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateInlineCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`.
func InlineCode_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	if err := validateInlineCode_FromCfnParametersParameters(props); err != nil {
		panic(err)
	}
	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
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
func InlineCode_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateInlineCode_FromDockerBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func InlineCode_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	if err := validateInlineCode_FromEcrImageParameters(repository, props); err != nil {
		panic(err)
	}
	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func InlineCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateInlineCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineCode) Bind(_scope constructs.Construct) *CodeConfig {
	if err := i.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineCode) BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions) {
	if err := i.validateBindToResourceParameters(_resource, _options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}


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

// Represents the Lambda Handler Code.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   signingProfile := signer.NewSigningProfile(this, jsii.String("SigningProfile"), &SigningProfileProps{
//   	Platform: signer.Platform_AWS_LAMBDA_SHA384_ECDSA(),
//   })
//
//   codeSigningConfig := lambda.NewCodeSigningConfig(this, jsii.String("CodeSigningConfig"), &CodeSigningConfigProps{
//   	SigningProfiles: []iSigningProfile{
//   		signingProfile,
//   	},
//   })
//
//   lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	CodeSigningConfig: CodeSigningConfig,
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   })
//
type Code interface {
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions)
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Code",
		nil, // no parameters
		c,
	)
}

// Loads the function code from a local disk path.
func Code_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func Code_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	if err := validateCode_FromAssetImageParameters(directory, props); err != nil {
		panic(err)
	}
	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func Code_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`.
func Code_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	if err := validateCode_FromCfnParametersParameters(props); err != nil {
		panic(err)
	}
	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
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
func Code_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateCode_FromDockerBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func Code_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	if err := validateCode_FromEcrImageParameters(repository, props); err != nil {
		panic(err)
	}
	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func Code_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Code) Bind(scope constructs.Construct) *CodeConfig {
	if err := c.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Code) BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions) {
	if err := c.validateBindToResourceParameters(_resource, _options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}


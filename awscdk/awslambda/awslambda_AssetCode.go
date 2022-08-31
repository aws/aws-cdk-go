package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

// Lambda code from a local directory.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   /*
//    * Stack verification steps:
//    * * `curl -s -o /dev/null -w "%{http_code}" <url>` should return 401
//    * * `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: deny' <url>` should return 403
//    * * `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: allow' <url>` should return 200
//    */
//
//   app := awscdk.NewApp()
//   stack := awscdk.NewStack(app, jsii.String("TokenAuthorizerInteg"))
//
//   authorizerFn := lambda.NewFunction(stack, jsii.String("MyAuthorizerFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.assetCode.fromAsset(path.join(__dirname, jsii.String("integ.token-authorizer.handler"))),
//   })
//
//   restapi := awscdk.NewRestApi(stack, jsii.String("MyRestApi"))
//
//   authorizer := awscdk.NewTokenAuthorizer(stack, jsii.String("MyAuthorizer"), &tokenAuthorizerProps{
//   	handler: authorizerFn,
//   })
//
//   restapi.root.addMethod(jsii.String("ANY"), awscdk.NewMockIntegration(&integrationOptions{
//   	integrationResponses: []integrationResponse{
//   		&integrationResponse{
//   			statusCode: jsii.String("200"),
//   		},
//   	},
//   	passthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   	requestTemplates: map[string]*string{
//   		"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   	},
//   }), &methodOptions{
//   	methodResponses: []methodResponse{
//   		&methodResponse{
//   			statusCode: jsii.String("200"),
//   		},
//   	},
//   	authorizer: authorizer,
//   })
//
// Experimental.
type AssetCode interface {
	Code
	// Determines whether this Code is inline code or not.
	// Experimental.
	IsInline() *bool
	// The path to the asset file or directory.
	// Experimental.
	Path() *string
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	// Experimental.
	Bind(scope awscdk.Construct) *CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	// Experimental.
	BindToResource(resource awscdk.CfnResource, options *ResourceBindOptions)
}

// The jsii proxy struct for AssetCode
type jsiiProxy_AssetCode struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_AssetCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetCode) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// Experimental.
func NewAssetCode(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	j := jsiiProxy_AssetCode{}

	_jsii_.Create(
		"monocdk.aws_lambda.AssetCode",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetCode_Override(a AssetCode, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda.AssetCode",
		[]interface{}{path, options},
		a,
	)
}

// DEPRECATED.
// Deprecated: use `fromAsset`.
func AssetCode_Asset(path *string) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
		"asset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `fromBucket`.
func AssetCode_Bucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
		"bucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `fromCfnParameters`.
func AssetCode_CfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
		"cfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from a local disk path.
// Experimental.
func AssetCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
// Experimental.
func AssetCode_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
// Experimental.
func AssetCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
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
func AssetCode_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
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
func AssetCode_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
// Experimental.
func AssetCode_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
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
func AssetCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// DEPRECATED.
// Deprecated: use `fromInline`.
func AssetCode_Inline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.AssetCode",
		"inline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetCode) Bind(scope awscdk.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetCode) BindToResource(resource awscdk.CfnResource, options *ResourceBindOptions) {
	_jsii_.InvokeVoid(
		a,
		"bindToResource",
		[]interface{}{resource, options},
	)
}


package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Origin for a Lambda Function URL.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn Function
//
//   fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
//   	AuthType: lambda.FunctionUrlAuthType_NONE,
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewFunctionUrlOrigin(fnUrl, &FunctionUrlOriginProps{
//   			ReadTimeout: awscdk.Duration_Seconds(jsii.Number(30)),
//   			ResponseCompletionTimeout: awscdk.Duration_*Seconds(jsii.Number(90)),
//   			KeepaliveTimeout: awscdk.Duration_*Seconds(jsii.Number(45)),
//   		}),
//   	},
//   })
//
type FunctionUrlOrigin interface {
	awscloudfront.OriginBase
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
	RenderVpcOriginConfig() *awscloudfront.CfnDistribution_VpcOriginConfigProperty
	// Validates that responseCompletionTimeout is greater than or equal to readTimeout when both are specified.
	//
	// This method should be called by subclasses that support readTimeout.
	ValidateResponseCompletionTimeoutWithReadTimeout(responseCompletionTimeout awscdk.Duration, readTimeout awscdk.Duration)
}

// The jsii proxy struct for FunctionUrlOrigin
type jsiiProxy_FunctionUrlOrigin struct {
	internal.Type__awscloudfrontOriginBase
}

func NewFunctionUrlOrigin(lambdaFunctionUrl awslambda.IFunctionUrl, props *FunctionUrlOriginProps) FunctionUrlOrigin {
	_init_.Initialize()

	if err := validateNewFunctionUrlOriginParameters(lambdaFunctionUrl, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionUrlOrigin{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.FunctionUrlOrigin",
		[]interface{}{lambdaFunctionUrl, props},
		&j,
	)

	return &j
}

func NewFunctionUrlOrigin_Override(f FunctionUrlOrigin, lambdaFunctionUrl awslambda.IFunctionUrl, props *FunctionUrlOriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.FunctionUrlOrigin",
		[]interface{}{lambdaFunctionUrl, props},
		f,
	)
}

// Create a Lambda Function URL Origin with Origin Access Control (OAC) configured.
func FunctionUrlOrigin_WithOriginAccessControl(lambdaFunctionUrl awslambda.IFunctionUrl, props *FunctionUrlOriginWithOACProps) awscloudfront.IOrigin {
	_init_.Initialize()

	if err := validateFunctionUrlOrigin_WithOriginAccessControlParameters(lambdaFunctionUrl, props); err != nil {
		panic(err)
	}
	var returns awscloudfront.IOrigin

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront_origins.FunctionUrlOrigin",
		"withOriginAccessControl",
		[]interface{}{lambdaFunctionUrl, props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionUrlOrigin) Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := f.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionUrlOrigin) RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		f,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionUrlOrigin) RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		f,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionUrlOrigin) RenderVpcOriginConfig() *awscloudfront.CfnDistribution_VpcOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_VpcOriginConfigProperty

	_jsii_.Invoke(
		f,
		"renderVpcOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionUrlOrigin) ValidateResponseCompletionTimeoutWithReadTimeout(responseCompletionTimeout awscdk.Duration, readTimeout awscdk.Duration) {
	_jsii_.InvokeVoid(
		f,
		"validateResponseCompletionTimeoutWithReadTimeout",
		[]interface{}{responseCompletionTimeout, readTimeout},
	)
}


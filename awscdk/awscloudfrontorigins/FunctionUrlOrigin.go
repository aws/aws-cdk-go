package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

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
//   var fn function
//
//   fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
//   	AuthType: lambda.FunctionUrlAuthType_NONE,
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewFunctionUrlOrigin(fnUrl),
//   	},
//   })
//
type FunctionUrlOrigin interface {
	awscloudfront.OriginBase
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
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

func (f *jsiiProxy_FunctionUrlOrigin) Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := f.validateBindParameters(_scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{_scope, options},
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


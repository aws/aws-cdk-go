package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Origin for an API Gateway REST API.
//
// Example:
//   var api restApi
//
//   cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewRestApiOrigin(api),
//   	},
//   })
//
type RestApiOrigin interface {
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

// The jsii proxy struct for RestApiOrigin
type jsiiProxy_RestApiOrigin struct {
	internal.Type__awscloudfrontOriginBase
}

func NewRestApiOrigin(restApi awsapigateway.RestApiBase, props *RestApiOriginProps) RestApiOrigin {
	_init_.Initialize()

	if err := validateNewRestApiOriginParameters(restApi, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RestApiOrigin{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.RestApiOrigin",
		[]interface{}{restApi, props},
		&j,
	)

	return &j
}

func NewRestApiOrigin_Override(r RestApiOrigin, restApi awsapigateway.RestApiBase, props *RestApiOriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.RestApiOrigin",
		[]interface{}{restApi, props},
		r,
	)
}

func (r *jsiiProxy_RestApiOrigin) Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := r.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiOrigin) RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		r,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiOrigin) RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		r,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiOrigin) RenderVpcOriginConfig() *awscloudfront.CfnDistribution_VpcOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_VpcOriginConfigProperty

	_jsii_.Invoke(
		r,
		"renderVpcOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApiOrigin) ValidateResponseCompletionTimeoutWithReadTimeout(responseCompletionTimeout awscdk.Duration, readTimeout awscdk.Duration) {
	_jsii_.InvokeVoid(
		r,
		"validateResponseCompletionTimeoutWithReadTimeout",
		[]interface{}{responseCompletionTimeout, readTimeout},
	)
}


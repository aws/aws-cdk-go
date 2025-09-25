package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Origin for a v2 load balancer.
//
// Example:
//   // Creates a distribution from an ELBv2 load balancer
//   var vpc vpc
//
//   // Create an application load balancer in a VPC. 'internetFacing' must be 'true'
//   // for CloudFront to access the load balancer and use it as an origin.
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewLoadBalancerV2Origin(lb),
//   	},
//   })
//
type LoadBalancerV2Origin interface {
	HttpOrigin
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

// The jsii proxy struct for LoadBalancerV2Origin
type jsiiProxy_LoadBalancerV2Origin struct {
	jsiiProxy_HttpOrigin
}

func NewLoadBalancerV2Origin(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, props *LoadBalancerV2OriginProps) LoadBalancerV2Origin {
	_init_.Initialize()

	if err := validateNewLoadBalancerV2OriginParameters(loadBalancer, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancerV2Origin{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.LoadBalancerV2Origin",
		[]interface{}{loadBalancer, props},
		&j,
	)

	return &j
}

func NewLoadBalancerV2Origin_Override(l LoadBalancerV2Origin, loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, props *LoadBalancerV2OriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.LoadBalancerV2Origin",
		[]interface{}{loadBalancer, props},
		l,
	)
}

func (l *jsiiProxy_LoadBalancerV2Origin) Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := l.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerV2Origin) RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		l,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerV2Origin) RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		l,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerV2Origin) RenderVpcOriginConfig() *awscloudfront.CfnDistribution_VpcOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_VpcOriginConfigProperty

	_jsii_.Invoke(
		l,
		"renderVpcOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerV2Origin) ValidateResponseCompletionTimeoutWithReadTimeout(responseCompletionTimeout awscdk.Duration, readTimeout awscdk.Duration) {
	_jsii_.InvokeVoid(
		l,
		"validateResponseCompletionTimeoutWithReadTimeout",
		[]interface{}{responseCompletionTimeout, readTimeout},
	)
}


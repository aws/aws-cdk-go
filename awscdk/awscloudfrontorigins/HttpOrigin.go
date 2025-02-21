package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Origin for an HTTP server or S3 bucket configured for website hosting.
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewOriginGroup(&OriginGroupProps{
//   			PrimaryOrigin: origins.S3BucketOrigin_WithOriginAccessControl(myBucket),
//   			FallbackOrigin: origins.NewHttpOrigin(jsii.String("www.example.com")),
//   			// optional, defaults to: 500, 502, 503 and 504
//   			FallbackStatusCodes: []*f64{
//   				jsii.Number(404),
//   			},
//   		}),
//   	},
//   })
//
type HttpOrigin interface {
	awscloudfront.OriginBase
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
	RenderVpcOriginConfig() *awscloudfront.CfnDistribution_VpcOriginConfigProperty
}

// The jsii proxy struct for HttpOrigin
type jsiiProxy_HttpOrigin struct {
	internal.Type__awscloudfrontOriginBase
}

func NewHttpOrigin(domainName *string, props *HttpOriginProps) HttpOrigin {
	_init_.Initialize()

	if err := validateNewHttpOriginParameters(domainName, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpOrigin{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.HttpOrigin",
		[]interface{}{domainName, props},
		&j,
	)

	return &j
}

func NewHttpOrigin_Override(h HttpOrigin, domainName *string, props *HttpOriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.HttpOrigin",
		[]interface{}{domainName, props},
		h,
	)
}

func (h *jsiiProxy_HttpOrigin) Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := h.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpOrigin) RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		h,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpOrigin) RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		h,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpOrigin) RenderVpcOriginConfig() *awscloudfront.CfnDistribution_VpcOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_VpcOriginConfigProperty

	_jsii_.Invoke(
		h,
		"renderVpcOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}


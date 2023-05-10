package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfrontorigins/internal"
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
// Experimental.
type RestApiOrigin interface {
	awscloudfront.OriginBase
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	// Experimental.
	Bind(_scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	// Experimental.
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	// Experimental.
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
}

// The jsii proxy struct for RestApiOrigin
type jsiiProxy_RestApiOrigin struct {
	internal.Type__awscloudfrontOriginBase
}

// Experimental.
func NewRestApiOrigin(restApi awsapigateway.RestApi, props *RestApiOriginProps) RestApiOrigin {
	_init_.Initialize()

	if err := validateNewRestApiOriginParameters(restApi, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RestApiOrigin{}

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.RestApiOrigin",
		[]interface{}{restApi, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRestApiOrigin_Override(r RestApiOrigin, restApi awsapigateway.RestApi, props *RestApiOriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.RestApiOrigin",
		[]interface{}{restApi, props},
		r,
	)
}

func (r *jsiiProxy_RestApiOrigin) Bind(_scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := r.validateBindParameters(_scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_scope, options},
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


package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

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
//   cloudfront.NewDistribution(this, jsii.String("Distribution"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewRestApiOrigin(api),
//   	},
//   })
//
type RestApiOrigin interface {
	awscloudfront.OriginBase
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
}

// The jsii proxy struct for RestApiOrigin
type jsiiProxy_RestApiOrigin struct {
	internal.Type__awscloudfrontOriginBase
}

func NewRestApiOrigin(restApi awsapigateway.RestApi, props *RestApiOriginProps) RestApiOrigin {
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

func NewRestApiOrigin_Override(r RestApiOrigin, restApi awsapigateway.RestApi, props *RestApiOriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.RestApiOrigin",
		[]interface{}{restApi, props},
		r,
	)
}

func (r *jsiiProxy_RestApiOrigin) Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
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


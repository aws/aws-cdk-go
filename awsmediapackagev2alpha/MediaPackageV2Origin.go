package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFront Origin for AWS Elemental MediaPackage V2 endpoints.
//
// Automatically creates an OAC and wires the origin endpoint policy
// to grant the CloudFront distribution access.
//
// Uses `addToResourcePolicy()` on the origin endpoint, which is compatible
// with other policy statements added to the same endpoint. Do not use this
// alongside a manually created `OriginEndpointPolicy` construct for the same endpoint.
//
// Example:
//   var endpoint OriginEndpoint
//   var group ChannelGroup
//
//
//   cloudfront.NewDistribution(this, jsii.String("Dist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: awsmediapackagev2alpha.NewMediaPackageV2Origin(endpoint, &MediaPackageV2OriginProps{
//   			ChannelGroup: group,
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-mediapackage.html
//
// Experimental.
type MediaPackageV2Origin interface {
	awscloudfront.OriginBase
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	// Experimental.
	Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	// Experimental.
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	// Experimental.
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
	// Experimental.
	RenderVpcOriginConfig() *awscloudfront.CfnDistribution_VpcOriginConfigProperty
	// Validates that responseCompletionTimeout is greater than or equal to readTimeout when both are specified.
	//
	// This method should be called by subclasses that support readTimeout.
	// Experimental.
	ValidateResponseCompletionTimeoutWithReadTimeout(responseCompletionTimeout awscdk.Duration, readTimeout awscdk.Duration)
}

// The jsii proxy struct for MediaPackageV2Origin
type jsiiProxy_MediaPackageV2Origin struct {
	internal.Type__awscloudfrontOriginBase
}

// Experimental.
func NewMediaPackageV2Origin(endpoint IOriginEndpoint, props *MediaPackageV2OriginProps) MediaPackageV2Origin {
	_init_.Initialize()

	if err := validateNewMediaPackageV2OriginParameters(endpoint, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaPackageV2Origin{}

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.MediaPackageV2Origin",
		[]interface{}{endpoint, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMediaPackageV2Origin_Override(m MediaPackageV2Origin, endpoint IOriginEndpoint, props *MediaPackageV2OriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.MediaPackageV2Origin",
		[]interface{}{endpoint, props},
		m,
	)
}

func (m *jsiiProxy_MediaPackageV2Origin) Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := m.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaPackageV2Origin) RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		m,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaPackageV2Origin) RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		m,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaPackageV2Origin) RenderVpcOriginConfig() *awscloudfront.CfnDistribution_VpcOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_VpcOriginConfigProperty

	_jsii_.Invoke(
		m,
		"renderVpcOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaPackageV2Origin) ValidateResponseCompletionTimeoutWithReadTimeout(responseCompletionTimeout awscdk.Duration, readTimeout awscdk.Duration) {
	_jsii_.InvokeVoid(
		m,
		"validateResponseCompletionTimeoutWithReadTimeout",
		[]interface{}{responseCompletionTimeout, readTimeout},
	)
}


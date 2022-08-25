package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a distribution origin, that describes the Amazon S3 bucket, HTTP server (for example, a web server), Amazon MediaStore, or other server from which CloudFront gets your files.
type OriginBase interface {
	IOrigin
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	Bind(_scope constructs.Construct, options *OriginBindOptions) *OriginBindConfig
	RenderCustomOriginConfig() *CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *CfnDistribution_S3OriginConfigProperty
}

// The jsii proxy struct for OriginBase
type jsiiProxy_OriginBase struct {
	jsiiProxy_IOrigin
}

func NewOriginBase_Override(o OriginBase, domainName *string, props *OriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.OriginBase",
		[]interface{}{domainName, props},
		o,
	)
}

func (o *jsiiProxy_OriginBase) Bind(_scope constructs.Construct, options *OriginBindOptions) *OriginBindConfig {
	var returns *OriginBindConfig

	_jsii_.Invoke(
		o,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginBase) RenderCustomOriginConfig() *CfnDistribution_CustomOriginConfigProperty {
	var returns *CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		o,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginBase) RenderS3OriginConfig() *CfnDistribution_S3OriginConfigProperty {
	var returns *CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		o,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}


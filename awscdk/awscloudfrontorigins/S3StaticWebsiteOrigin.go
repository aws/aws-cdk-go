package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Origin for a S3 bucket configured as a website endpoint.
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3StaticWebsiteOrigin(myBucket),
//   	},
//   })
//
type S3StaticWebsiteOrigin interface {
	HttpOrigin
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
}

// The jsii proxy struct for S3StaticWebsiteOrigin
type jsiiProxy_S3StaticWebsiteOrigin struct {
	jsiiProxy_HttpOrigin
}

func NewS3StaticWebsiteOrigin(bucket awss3.IBucket, props *S3StaticWebsiteOriginProps) S3StaticWebsiteOrigin {
	_init_.Initialize()

	if err := validateNewS3StaticWebsiteOriginParameters(bucket, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3StaticWebsiteOrigin{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.S3StaticWebsiteOrigin",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

func NewS3StaticWebsiteOrigin_Override(s S3StaticWebsiteOrigin, bucket awss3.IBucket, props *S3StaticWebsiteOriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.S3StaticWebsiteOrigin",
		[]interface{}{bucket, props},
		s,
	)
}

func (s *jsiiProxy_S3StaticWebsiteOrigin) Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := s.validateBindParameters(_scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StaticWebsiteOrigin) RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		s,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StaticWebsiteOrigin) RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		s,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}


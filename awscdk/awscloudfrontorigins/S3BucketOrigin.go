package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// A S3 Bucket Origin.
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
type S3BucketOrigin interface {
	awscloudfront.OriginBase
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
}

// The jsii proxy struct for S3BucketOrigin
type jsiiProxy_S3BucketOrigin struct {
	internal.Type__awscloudfrontOriginBase
}

func NewS3BucketOrigin_Override(s S3BucketOrigin, bucket awss3.IBucket, props *S3BucketOriginBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.S3BucketOrigin",
		[]interface{}{bucket, props},
		s,
	)
}

// Create a S3 Origin with default S3 bucket settings (no origin access control).
func S3BucketOrigin_WithBucketDefaults(bucket awss3.IBucket, props *awscloudfront.OriginProps) awscloudfront.IOrigin {
	_init_.Initialize()

	if err := validateS3BucketOrigin_WithBucketDefaultsParameters(bucket, props); err != nil {
		panic(err)
	}
	var returns awscloudfront.IOrigin

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront_origins.S3BucketOrigin",
		"withBucketDefaults",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}

// Create a S3 Origin with Origin Access Control (OAC) configured.
func S3BucketOrigin_WithOriginAccessControl(bucket awss3.IBucket, props *S3BucketOriginWithOACProps) awscloudfront.IOrigin {
	_init_.Initialize()

	if err := validateS3BucketOrigin_WithOriginAccessControlParameters(bucket, props); err != nil {
		panic(err)
	}
	var returns awscloudfront.IOrigin

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront_origins.S3BucketOrigin",
		"withOriginAccessControl",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}

// Create a S3 Origin with Origin Access Identity (OAI) configured OAI is a legacy feature and we **strongly** recommend you to use OAC via `withOriginAccessControl()` unless it is not supported in your required region (e.g. China regions).
func S3BucketOrigin_WithOriginAccessIdentity(bucket awss3.IBucket, props *S3BucketOriginWithOAIProps) awscloudfront.IOrigin {
	_init_.Initialize()

	if err := validateS3BucketOrigin_WithOriginAccessIdentityParameters(bucket, props); err != nil {
		panic(err)
	}
	var returns awscloudfront.IOrigin

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront_origins.S3BucketOrigin",
		"withOriginAccessIdentity",
		[]interface{}{bucket, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketOrigin) Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
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

func (s *jsiiProxy_S3BucketOrigin) RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		s,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketOrigin) RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		s,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}


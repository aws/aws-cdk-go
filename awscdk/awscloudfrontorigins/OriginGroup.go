package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Origin that represents a group.
//
// Consists of a primary Origin,
// and a fallback Origin called when the primary returns one of the provided HTTP status codes.
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
type OriginGroup interface {
	awscloudfront.IOrigin
	// The method called when a given Origin is added (for the first time) to a Distribution.
	Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
}

// The jsii proxy struct for OriginGroup
type jsiiProxy_OriginGroup struct {
	internal.Type__awscloudfrontIOrigin
}

func NewOriginGroup(props *OriginGroupProps) OriginGroup {
	_init_.Initialize()

	if err := validateNewOriginGroupParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OriginGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.OriginGroup",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewOriginGroup_Override(o OriginGroup, props *OriginGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.OriginGroup",
		[]interface{}{props},
		o,
	)
}

func (o *jsiiProxy_OriginGroup) Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	if err := o.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		o,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}


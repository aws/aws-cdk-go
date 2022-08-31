package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfrontorigins/internal"
)

// An Origin that represents a group.
//
// Consists of a primary Origin,
// and a fallback Origin called when the primary returns one of the provided HTTP status codes.
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewOriginGroup(&originGroupProps{
//   			primaryOrigin: origins.NewS3Origin(myBucket),
//   			fallbackOrigin: origins.NewHttpOrigin(jsii.String("www.example.com")),
//   			// optional, defaults to: 500, 502, 503 and 504
//   			fallbackStatusCodes: []*f64{
//   				jsii.Number(404),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type OriginGroup interface {
	awscloudfront.IOrigin
	// The method called when a given Origin is added (for the first time) to a Distribution.
	// Experimental.
	Bind(scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
}

// The jsii proxy struct for OriginGroup
type jsiiProxy_OriginGroup struct {
	internal.Type__awscloudfrontIOrigin
}

// Experimental.
func NewOriginGroup(props *OriginGroupProps) OriginGroup {
	_init_.Initialize()

	j := jsiiProxy_OriginGroup{}

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.OriginGroup",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewOriginGroup_Override(o OriginGroup, props *OriginGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.OriginGroup",
		[]interface{}{props},
		o,
	)
}

func (o *jsiiProxy_OriginGroup) Bind(scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		o,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}


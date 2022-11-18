package awscloudfrontorigins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
)

// Construction properties for {@link OriginGroup}.
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
type OriginGroupProps struct {
	// The fallback origin that should serve requests when the primary fails.
	FallbackOrigin awscloudfront.IOrigin `field:"required" json:"fallbackOrigin" yaml:"fallbackOrigin"`
	// The primary origin that should serve requests for this group.
	PrimaryOrigin awscloudfront.IOrigin `field:"required" json:"primaryOrigin" yaml:"primaryOrigin"`
	// The list of HTTP status codes that, when returned from the primary origin, would cause querying the fallback origin.
	FallbackStatusCodes *[]*float64 `field:"optional" json:"fallbackStatusCodes" yaml:"fallbackStatusCodes"`
}


package awscloudfrontorigins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
)

// Construction properties for `OriginGroup`.
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
type OriginGroupProps struct {
	// The fallback origin that should serve requests when the primary fails.
	FallbackOrigin awscloudfront.IOrigin `field:"required" json:"fallbackOrigin" yaml:"fallbackOrigin"`
	// The primary origin that should serve requests for this group.
	PrimaryOrigin awscloudfront.IOrigin `field:"required" json:"primaryOrigin" yaml:"primaryOrigin"`
	// The list of HTTP status codes that, when returned from the primary origin, would cause querying the fallback origin.
	// Default: - 500, 502, 503 and 504.
	//
	FallbackStatusCodes *[]*float64 `field:"optional" json:"fallbackStatusCodes" yaml:"fallbackStatusCodes"`
	// The selection criteria for the origin group.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/high_availability_origin_failover.html#concept_origin_groups.creating
	//
	// Default: - OriginSelectionCriteria.DEFAULT
	//
	SelectionCriteria awscloudfront.OriginSelectionCriteria `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}


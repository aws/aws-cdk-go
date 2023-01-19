package awscloudfrontorigins

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront"
)

// Properties to use to customize an S3 Origin.
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewS3Origin(myBucket, &s3OriginProps{
//   			customHeaders: map[string]*string{
//   				"Foo": jsii.String("bar"),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type S3OriginProps struct {
	// The number of times that CloudFront attempts to connect to the origin;
	//
	// valid values are 1, 2, or 3 attempts.
	// Experimental.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// Valid values are 1-10 seconds, inclusive.
	// Experimental.
	ConnectionTimeout awscdk.Duration `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// A list of HTTP header names and values that CloudFront adds to requests it sends to the origin.
	// Experimental.
	CustomHeaders *map[string]*string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	// Experimental.
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// Must begin, but not end, with '/' (e.g., '/production/images').
	// Experimental.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// An optional Origin Access Identity of the origin identity cloudfront will use when calling your s3 bucket.
	// Experimental.
	OriginAccessIdentity awscloudfront.IOriginAccessIdentity `field:"optional" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}


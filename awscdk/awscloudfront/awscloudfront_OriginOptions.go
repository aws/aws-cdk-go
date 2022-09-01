package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options to define an Origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originOptions := &originOptions{
//   	connectionAttempts: jsii.Number(123),
//   	connectionTimeout: cdk.duration.minutes(jsii.Number(30)),
//   	customHeaders: map[string]*string{
//   		"customHeadersKey": jsii.String("customHeaders"),
//   	},
//   	originShieldRegion: jsii.String("originShieldRegion"),
//   }
//
type OriginOptions struct {
	// The number of times that CloudFront attempts to connect to the origin;
	//
	// valid values are 1, 2, or 3 attempts.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// Valid values are 1-10 seconds, inclusive.
	ConnectionTimeout awscdk.Duration `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// A list of HTTP header names and values that CloudFront adds to requests it sends to the origin.
	CustomHeaders *map[string]*string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
}


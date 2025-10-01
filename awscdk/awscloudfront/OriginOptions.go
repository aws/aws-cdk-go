package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options to define an Origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originOptions := &OriginOptions{
//   	ConnectionAttempts: jsii.Number(123),
//   	ConnectionTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	CustomHeaders: map[string]*string{
//   		"customHeadersKey": jsii.String("customHeaders"),
//   	},
//   	OriginAccessControlId: jsii.String("originAccessControlId"),
//   	OriginId: jsii.String("originId"),
//   	OriginShieldEnabled: jsii.Boolean(false),
//   	OriginShieldRegion: jsii.String("originShieldRegion"),
//   	ResponseCompletionTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type OriginOptions struct {
	// The number of times that CloudFront attempts to connect to the origin;
	//
	// valid values are 1, 2, or 3 attempts.
	// Default: 3.
	//
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// Valid values are 1-10 seconds, inclusive.
	// Default: Duration.seconds(10)
	//
	ConnectionTimeout awscdk.Duration `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// A list of HTTP header names and values that CloudFront adds to requests it sends to the origin.
	// Default: {}.
	//
	CustomHeaders *map[string]*string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// The unique identifier of an origin access control for this origin.
	// Default: - no origin access control.
	//
	OriginAccessControlId *string `field:"optional" json:"originAccessControlId" yaml:"originAccessControlId"`
	// A unique identifier for the origin.
	//
	// This value must be unique within the distribution.
	// Default: - an originid will be generated for you.
	//
	OriginId *string `field:"optional" json:"originId" yaml:"originId"`
	// Origin Shield is enabled by setting originShieldRegion to a valid region, after this to disable Origin Shield again you must set this flag to false.
	// Default: - true.
	//
	OriginShieldEnabled *bool `field:"optional" json:"originShieldEnabled" yaml:"originShieldEnabled"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	// Default: - origin shield not enabled.
	//
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
	// The time that a request from CloudFront to the origin can stay open and wait for a response.
	//
	// If the complete response isn't received from the origin by this time, CloudFront ends the connection.
	//
	// Valid values are 1-3600 seconds, inclusive.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#response-completion-timeout
	//
	// Default: undefined -  AWS CloudFront default is not enforcing a maximum value.
	//
	ResponseCompletionTimeout awscdk.Duration `field:"optional" json:"responseCompletionTimeout" yaml:"responseCompletionTimeout"`
}


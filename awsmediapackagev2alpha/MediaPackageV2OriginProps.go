package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront"
)

// Properties for a MediaPackage V2 Origin with OAC.
//
// Example:
//   var endpoint OriginEndpoint
//   var group ChannelGroup
//
//
//   cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: awsmediapackagev2alpha.NewMediaPackageV2Origin(endpoint, &MediaPackageV2OriginProps{
//   			ChannelGroup: group,
//   		}),
//   	},
//   })
//
// Experimental.
type MediaPackageV2OriginProps struct {
	// The number of times that CloudFront attempts to connect to the origin;
	//
	// valid values are 1, 2, or 3 attempts.
	// Default: 3.
	//
	// Experimental.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// Valid values are 1-10 seconds, inclusive.
	// Default: Duration.seconds(10)
	//
	// Experimental.
	ConnectionTimeout awscdk.Duration `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// A list of HTTP header names and values that CloudFront adds to requests it sends to the origin.
	// Default: {}.
	//
	// Experimental.
	CustomHeaders *map[string]*string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// The unique identifier of an origin access control for this origin.
	// Default: - no origin access control.
	//
	// Experimental.
	OriginAccessControlId *string `field:"optional" json:"originAccessControlId" yaml:"originAccessControlId"`
	// A unique identifier for the origin.
	//
	// This value must be unique within the distribution.
	// Default: - an originid will be generated for you.
	//
	// Experimental.
	OriginId *string `field:"optional" json:"originId" yaml:"originId"`
	// Origin Shield is enabled by setting originShieldRegion to a valid region, after this to disable Origin Shield again you must set this flag to false.
	// Default: - true.
	//
	// Experimental.
	OriginShieldEnabled *bool `field:"optional" json:"originShieldEnabled" yaml:"originShieldEnabled"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	// Default: - origin shield not enabled.
	//
	// Experimental.
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
	// Experimental.
	ResponseCompletionTimeout awscdk.Duration `field:"optional" json:"responseCompletionTimeout" yaml:"responseCompletionTimeout"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// Must begin, but not end, with '/' (e.g., '/production/images').
	// Default: '/'.
	//
	// Experimental.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// The channel group that the origin endpoint belongs to.
	//
	// Used to derive the egress domain for the CloudFront origin.
	// Experimental.
	ChannelGroup IChannelGroup `field:"required" json:"channelGroup" yaml:"channelGroup"`
	// Optional CDN authorization configuration.
	//
	// If you need CDN auth on this endpoint, provide it here so it is configured
	// on the first `addToResourcePolicy` call. If CDN auth is added separately
	// after this origin is bound, it will be ignored.
	// Default: - no CDN authorization.
	//
	// Experimental.
	CdnAuth *CdnAuthConfiguration `field:"optional" json:"cdnAuth" yaml:"cdnAuth"`
	// An optional Origin Access Control.
	// Default: - an Origin Access Control will be created automatically.
	//
	// Experimental.
	OriginAccessControl interfacesawscloudfront.IOriginAccessControlRef `field:"optional" json:"originAccessControl" yaml:"originAccessControl"`
}


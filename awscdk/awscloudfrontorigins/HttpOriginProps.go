package awscloudfrontorigins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
)

// Properties for an Origin backed by an S3 website-configured bucket, load balancer, or custom HTTP server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpOriginProps := &HttpOriginProps{
//   	ConnectionAttempts: jsii.Number(123),
//   	ConnectionTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	CustomHeaders: map[string]*string{
//   		"customHeadersKey": jsii.String("customHeaders"),
//   	},
//   	HttpPort: jsii.Number(123),
//   	HttpsPort: jsii.Number(123),
//   	KeepaliveTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	OriginAccessControlId: jsii.String("originAccessControlId"),
//   	OriginId: jsii.String("originId"),
//   	OriginPath: jsii.String("originPath"),
//   	OriginShieldEnabled: jsii.Boolean(false),
//   	OriginShieldRegion: jsii.String("originShieldRegion"),
//   	OriginSslProtocols: []originSslPolicy{
//   		awscdk.Aws_cloudfront.*originSslPolicy_SSL_V3,
//   	},
//   	ProtocolPolicy: awscdk.*Aws_cloudfront.OriginProtocolPolicy_HTTP_ONLY,
//   	ReadTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type HttpOriginProps struct {
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
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// Must begin, but not end, with '/' (e.g., '/production/images').
	// Default: '/'.
	//
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// The HTTP port that CloudFront uses to connect to the origin.
	// Default: 80.
	//
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The HTTPS port that CloudFront uses to connect to the origin.
	// Default: 443.
	//
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The valid range is from 1 to 180 seconds, inclusive.
	//
	// Note that values over 60 seconds are possible only after a limit increase request for the origin response timeout quota
	// has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time.
	// Default: Duration.seconds(5)
	//
	KeepaliveTimeout awscdk.Duration `field:"optional" json:"keepaliveTimeout" yaml:"keepaliveTimeout"`
	// The SSL versions to use when interacting with the origin.
	// Default: OriginSslPolicy.TLS_V1_2
	//
	OriginSslProtocols *[]awscloudfront.OriginSslPolicy `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
	// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin.
	// Default: OriginProtocolPolicy.HTTPS_ONLY
	//
	ProtocolPolicy awscloudfront.OriginProtocolPolicy `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin, also known as the origin response timeout.
	//
	// The valid range is from 1 to 180 seconds, inclusive.
	//
	// Note that values over 60 seconds are possible only after a limit increase request for the origin response timeout quota
	// has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time.
	// Default: Duration.seconds(30)
	//
	ReadTimeout awscdk.Duration `field:"optional" json:"readTimeout" yaml:"readTimeout"`
}


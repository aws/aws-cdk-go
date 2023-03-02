package awscloudfrontorigins

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront"
)

// Properties for an Origin backed by an S3 website-configured bucket, load balancer, or custom HTTP server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   httpOriginProps := &httpOriginProps{
//   	connectionAttempts: jsii.Number(123),
//   	connectionTimeout: duration,
//   	customHeaders: map[string]*string{
//   		"customHeadersKey": jsii.String("customHeaders"),
//   	},
//   	httpPort: jsii.Number(123),
//   	httpsPort: jsii.Number(123),
//   	keepaliveTimeout: duration,
//   	originPath: jsii.String("originPath"),
//   	originShieldRegion: jsii.String("originShieldRegion"),
//   	originSslProtocols: []originSslPolicy{
//   		awscdk.Aws_cloudfront.*originSslPolicy_SSL_V3,
//   	},
//   	protocolPolicy: awscdk.*Aws_cloudfront.originProtocolPolicy_HTTP_ONLY,
//   	readTimeout: duration,
//   }
//
// Experimental.
type HttpOriginProps struct {
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
	// The HTTP port that CloudFront uses to connect to the origin.
	// Experimental.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The HTTPS port that CloudFront uses to connect to the origin.
	// Experimental.
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The valid range is from 1 to 180 seconds, inclusive.
	//
	// Note that values over 60 seconds are possible only after a limit increase request for the origin response timeout quota
	// has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time.
	// Experimental.
	KeepaliveTimeout awscdk.Duration `field:"optional" json:"keepaliveTimeout" yaml:"keepaliveTimeout"`
	// The SSL versions to use when interacting with the origin.
	// Experimental.
	OriginSslProtocols *[]awscloudfront.OriginSslPolicy `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
	// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin.
	// Experimental.
	ProtocolPolicy awscloudfront.OriginProtocolPolicy `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin, also known as the origin response timeout.
	//
	// The valid range is from 1 to 180 seconds, inclusive.
	//
	// Note that values over 60 seconds are possible only after a limit increase request for the origin response timeout quota
	// has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time.
	// Experimental.
	ReadTimeout awscdk.Duration `field:"optional" json:"readTimeout" yaml:"readTimeout"`
}


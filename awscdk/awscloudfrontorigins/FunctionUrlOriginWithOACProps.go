package awscloudfrontorigins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
)

// Properties for configuring a Lambda Functions URLs with OAC.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn Function
//
//
//   fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
//   	AuthType: lambda.FunctionUrlAuthType_AWS_IAM,
//   })
//
//   // Define a custom OAC
//   oac := cloudfront.NewFunctionUrlOriginAccessControl(this, jsii.String("MyOAC"), &FunctionUrlOriginAccessControlProps{
//   	OriginAccessControlName: jsii.String("CustomLambdaOAC"),
//   	Signing: cloudfront.Signing_SIGV4_ALWAYS(),
//   })
//
//   // Set up Lambda Function URL with OAC in CloudFront Distribution
//   // Set up Lambda Function URL with OAC in CloudFront Distribution
//   cloudfront.NewDistribution(this, jsii.String("MyDistribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.FunctionUrlOrigin_WithOriginAccessControl(fnUrl, &FunctionUrlOriginWithOACProps{
//   			OriginAccessControl: oac,
//   		}),
//   	},
//   })
//
type FunctionUrlOriginWithOACProps struct {
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
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// Must begin, but not end, with '/' (e.g., '/production/images').
	// Default: '/'.
	//
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// Specifies which IP protocol CloudFront uses when connecting to your origin.
	//
	// If your origin uses both IPv4 and IPv6 protocols, you can choose dualstack to help optimize reliability.
	// Default: OriginIpAddressType.IPV4
	//
	IpAddressType awscloudfront.OriginIpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The valid range is from 1 to 180 seconds, inclusive.
	//
	// Note that values over 60 seconds are possible only after a limit increase request for the origin response timeout quota
	// has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time.
	// Default: Duration.seconds(5)
	//
	KeepaliveTimeout awscdk.Duration `field:"optional" json:"keepaliveTimeout" yaml:"keepaliveTimeout"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin.
	//
	// The valid range is from 1 to 180 seconds, inclusive.
	//
	// Note that values over 60 seconds are possible only after a limit increase request for the origin response timeout quota
	// has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time.
	// Default: Duration.seconds(30)
	//
	ReadTimeout awscdk.Duration `field:"optional" json:"readTimeout" yaml:"readTimeout"`
	// An optional Origin Access Control.
	// Default: - an Origin Access Control will be created.
	//
	OriginAccessControl awscloudfront.IOriginAccessControl `field:"optional" json:"originAccessControl" yaml:"originAccessControl"`
}


package awscloudfront


// An origin.
//
// An origin is the location where content is stored, and from which CloudFront gets content to serve to viewers. To specify an origin:
//
// - Use `S3OriginConfig` to specify an Amazon S3 bucket that is not configured with static website hosting.
// - Use `CustomOriginConfig` to specify all other kinds of origins, including:
//
// - An Amazon S3 bucket that is configured with static website hosting
// - An Elastic Load Balancing load balancer
// - An AWS Elemental MediaPackage endpoint
// - An AWS Elemental MediaStore container
// - Any other HTTP server, running on an Amazon EC2 instance or any other kind of host
//
// For the current maximum number of origins that you can specify per distribution, see [General Quotas on Web Distributions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html#limits-web-distributions) in the *Amazon CloudFront Developer Guide* (quotas were formerly referred to as limits).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originProperty := &originProperty{
//   	domainName: jsii.String("domainName"),
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	connectionAttempts: jsii.Number(123),
//   	connectionTimeout: jsii.Number(123),
//   	customOriginConfig: &customOriginConfigProperty{
//   		originProtocolPolicy: jsii.String("originProtocolPolicy"),
//
//   		// the properties below are optional
//   		httpPort: jsii.Number(123),
//   		httpsPort: jsii.Number(123),
//   		originKeepaliveTimeout: jsii.Number(123),
//   		originReadTimeout: jsii.Number(123),
//   		originSslProtocols: []*string{
//   			jsii.String("originSslProtocols"),
//   		},
//   	},
//   	originAccessControlId: jsii.String("originAccessControlId"),
//   	originCustomHeaders: []interface{}{
//   		&originCustomHeaderProperty{
//   			headerName: jsii.String("headerName"),
//   			headerValue: jsii.String("headerValue"),
//   		},
//   	},
//   	originPath: jsii.String("originPath"),
//   	originShield: &originShieldProperty{
//   		enabled: jsii.Boolean(false),
//   		originShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	s3OriginConfig: &s3OriginConfigProperty{
//   		originAccessIdentity: jsii.String("originAccessIdentity"),
//   	},
//   }
//
type CfnDistribution_OriginProperty struct {
	// The domain name for the origin.
	//
	// For more information, see [Origin Domain Name](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesDomainName) in the *Amazon CloudFront Developer Guide* .
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// A unique identifier for the origin. This value must be unique within the distribution.
	//
	// Use this value to specify the `TargetOriginId` in a `CacheBehavior` or `DefaultCacheBehavior` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The number of times that CloudFront attempts to connect to the origin.
	//
	// The minimum number is 1, the maximum is 3, and the default (if you don’t specify otherwise) is 3.
	//
	// For a custom origin (including an Amazon S3 bucket that’s configured with static website hosting), this value also specifies the number of times that CloudFront attempts to get a response from the origin, in the case of an [Origin Response Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginResponseTimeout) .
	//
	// For more information, see [Origin Connection Attempts](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#origin-connection-attempts) in the *Amazon CloudFront Developer Guide* .
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// The minimum timeout is 1 second, the maximum is 10 seconds, and the default (if you don’t specify otherwise) is 10 seconds.
	//
	// For more information, see [Origin Connection Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#origin-connection-timeout) in the *Amazon CloudFront Developer Guide* .
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// Use this type to specify an origin that is not an Amazon S3 bucket, with one exception.
	//
	// If the Amazon S3 bucket is configured with static website hosting, use this type. If the Amazon S3 bucket is not configured with static website hosting, use the `S3OriginConfig` type instead.
	CustomOriginConfig interface{} `field:"optional" json:"customOriginConfig" yaml:"customOriginConfig"`
	// `CfnDistribution.OriginProperty.OriginAccessControlId`.
	OriginAccessControlId *string `field:"optional" json:"originAccessControlId" yaml:"originAccessControlId"`
	// A list of HTTP header names and values that CloudFront adds to the requests that it sends to the origin.
	//
	// For more information, see [Adding Custom Headers to Origin Requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/add-origin-custom-headers.html) in the *Amazon CloudFront Developer Guide* .
	OriginCustomHeaders interface{} `field:"optional" json:"originCustomHeaders" yaml:"originCustomHeaders"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// For more information, see [Origin Path](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginPath) in the *Amazon CloudFront Developer Guide* .
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// CloudFront Origin Shield. Using Origin Shield can help reduce the load on your origin.
	//
	// For more information, see [Using Origin Shield](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html) in the *Amazon CloudFront Developer Guide* .
	OriginShield interface{} `field:"optional" json:"originShield" yaml:"originShield"`
	// Use this type to specify an origin that is an Amazon S3 bucket that is not configured with static website hosting.
	//
	// To specify any other type of origin, including an Amazon S3 bucket that is configured with static website hosting, use the `CustomOriginConfig` type instead.
	S3OriginConfig interface{} `field:"optional" json:"s3OriginConfig" yaml:"s3OriginConfig"`
}


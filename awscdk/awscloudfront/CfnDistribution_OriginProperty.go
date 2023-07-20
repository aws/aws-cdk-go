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
//   originProperty := &OriginProperty{
//   	DomainName: jsii.String("domainName"),
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	ConnectionAttempts: jsii.Number(123),
//   	ConnectionTimeout: jsii.Number(123),
//   	CustomOriginConfig: &CustomOriginConfigProperty{
//   		OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//
//   		// the properties below are optional
//   		HttpPort: jsii.Number(123),
//   		HttpsPort: jsii.Number(123),
//   		OriginKeepaliveTimeout: jsii.Number(123),
//   		OriginReadTimeout: jsii.Number(123),
//   		OriginSslProtocols: []*string{
//   			jsii.String("originSslProtocols"),
//   		},
//   	},
//   	OriginAccessControlId: jsii.String("originAccessControlId"),
//   	OriginCustomHeaders: []interface{}{
//   		&OriginCustomHeaderProperty{
//   			HeaderName: jsii.String("headerName"),
//   			HeaderValue: jsii.String("headerValue"),
//   		},
//   	},
//   	OriginPath: jsii.String("originPath"),
//   	OriginShield: &OriginShieldProperty{
//   		Enabled: jsii.Boolean(false),
//   		OriginShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	S3OriginConfig: &S3OriginConfigProperty{
//   		OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html
//
type CfnDistribution_OriginProperty struct {
	// The domain name for the origin.
	//
	// For more information, see [Origin Domain Name](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesDomainName) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// A unique identifier for the origin. This value must be unique within the distribution.
	//
	// Use this value to specify the `TargetOriginId` in a `CacheBehavior` or `DefaultCacheBehavior` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The number of times that CloudFront attempts to connect to the origin.
	//
	// The minimum number is 1, the maximum is 3, and the default (if you don't specify otherwise) is 3.
	//
	// For a custom origin (including an Amazon S3 bucket that's configured with static website hosting), this value also specifies the number of times that CloudFront attempts to get a response from the origin, in the case of an [Origin Response Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginResponseTimeout) .
	//
	// For more information, see [Origin Connection Attempts](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#origin-connection-attempts) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-connectionattempts
	//
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// The minimum timeout is 1 second, the maximum is 10 seconds, and the default (if you don't specify otherwise) is 10 seconds.
	//
	// For more information, see [Origin Connection Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#origin-connection-timeout) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-connectiontimeout
	//
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// Use this type to specify an origin that is not an Amazon S3 bucket, with one exception.
	//
	// If the Amazon S3 bucket is configured with static website hosting, use this type. If the Amazon S3 bucket is not configured with static website hosting, use the `S3OriginConfig` type instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-customoriginconfig
	//
	CustomOriginConfig interface{} `field:"optional" json:"customOriginConfig" yaml:"customOriginConfig"`
	// The unique identifier of an origin access control for this origin.
	//
	// For more information, see [Restricting access to an Amazon S3 origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-originaccesscontrolid
	//
	OriginAccessControlId *string `field:"optional" json:"originAccessControlId" yaml:"originAccessControlId"`
	// A list of HTTP header names and values that CloudFront adds to the requests that it sends to the origin.
	//
	// For more information, see [Adding Custom Headers to Origin Requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/add-origin-custom-headers.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-origincustomheaders
	//
	OriginCustomHeaders interface{} `field:"optional" json:"originCustomHeaders" yaml:"originCustomHeaders"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// For more information, see [Origin Path](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginPath) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-originpath
	//
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// CloudFront Origin Shield. Using Origin Shield can help reduce the load on your origin.
	//
	// For more information, see [Using Origin Shield](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-originshield
	//
	OriginShield interface{} `field:"optional" json:"originShield" yaml:"originShield"`
	// Use this type to specify an origin that is an Amazon S3 bucket that is not configured with static website hosting.
	//
	// To specify any other type of origin, including an Amazon S3 bucket that is configured with static website hosting, use the `CustomOriginConfig` type instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html#cfn-cloudfront-distribution-origin-s3originconfig
	//
	S3OriginConfig interface{} `field:"optional" json:"s3OriginConfig" yaml:"s3OriginConfig"`
}


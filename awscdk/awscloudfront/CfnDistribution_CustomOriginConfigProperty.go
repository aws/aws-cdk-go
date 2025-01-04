package awscloudfront


// A custom origin.
//
// A custom origin is any origin that is *not* an Amazon S3 bucket, with one exception. An Amazon S3 bucket that is [configured with static website hosting](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) *is* a custom origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customOriginConfigProperty := &CustomOriginConfigProperty{
//   	OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//
//   	// the properties below are optional
//   	HttpPort: jsii.Number(123),
//   	HttpsPort: jsii.Number(123),
//   	OriginKeepaliveTimeout: jsii.Number(123),
//   	OriginReadTimeout: jsii.Number(123),
//   	OriginSslProtocols: []*string{
//   		jsii.String("originSslProtocols"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html
//
type CfnDistribution_CustomOriginConfigProperty struct {
	// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin. Valid values are:.
	//
	// - `http-only` – CloudFront always uses HTTP to connect to the origin.
	// - `match-viewer` – CloudFront connects to the origin using the same protocol that the viewer used to connect to CloudFront.
	// - `https-only` – CloudFront always uses HTTPS to connect to the origin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-originprotocolpolicy
	//
	OriginProtocolPolicy *string `field:"required" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// The HTTP port that CloudFront uses to connect to the origin.
	//
	// Specify the HTTP port that the origin listens on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-httpport
	//
	// Default: - 80.
	//
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The HTTPS port that CloudFront uses to connect to the origin.
	//
	// Specify the HTTPS port that the origin listens on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-httpsport
	//
	// Default: - 443.
	//
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The minimum timeout is 1 second, the maximum is 60 seconds, and the default (if you don't specify otherwise) is 5 seconds.
	//
	// For more information, see [Keep-alive timeout (custom origins only)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginKeepaliveTimeout) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-originkeepalivetimeout
	//
	// Default: - 5.
	//
	OriginKeepaliveTimeout *float64 `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin.
	//
	// This is also known as the *origin response timeout* . The minimum timeout is 1 second, the maximum is 60 seconds, and the default (if you don't specify otherwise) is 30 seconds.
	//
	// For more information, see [Response timeout (custom origins only)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginResponseTimeout) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-originreadtimeout
	//
	// Default: - 30.
	//
	OriginReadTimeout *float64 `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
	// Specifies the minimum SSL/TLS protocol that CloudFront uses when connecting to your origin over HTTPS.
	//
	// Valid values include `SSLv3` , `TLSv1` , `TLSv1.1` , and `TLSv1.2` .
	//
	// For more information, see [Minimum Origin SSL Protocol](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginSSLProtocols) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-originsslprotocols
	//
	OriginSslProtocols *[]*string `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
}


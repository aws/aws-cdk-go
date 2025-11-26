package previewawscloudfrontmixins


// A custom origin.
//
// A custom origin is any origin that is *not* an Amazon S3 bucket, with one exception. An Amazon S3 bucket that is [configured with static website hosting](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) *is* a custom origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customOriginConfigProperty := &CustomOriginConfigProperty{
//   	HttpPort: jsii.Number(123),
//   	HttpsPort: jsii.Number(123),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	OriginKeepaliveTimeout: jsii.Number(123),
//   	OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   	OriginReadTimeout: jsii.Number(123),
//   	OriginSslProtocols: []*string{
//   		jsii.String("originSslProtocols"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html
//
type CfnDistributionPropsMixin_CustomOriginConfigProperty struct {
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
	// Specifies which IP protocol CloudFront uses when connecting to your origin.
	//
	// If your origin uses both IPv4 and IPv6 protocols, you can choose `dualstack` to help optimize reliability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The minimum timeout is 1 second, the maximum is 120 seconds, and the default (if you don't specify otherwise) is 5 seconds.
	//
	// For more information, see [Keep-alive timeout (custom origins only)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#DownloadDistValuesOriginKeepaliveTimeout) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-originkeepalivetimeout
	//
	// Default: - 5.
	//
	OriginKeepaliveTimeout *float64 `field:"optional" json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin. Valid values are:.
	//
	// - `http-only` – CloudFront always uses HTTP to connect to the origin.
	// - `match-viewer` – CloudFront connects to the origin using the same protocol that the viewer used to connect to CloudFront.
	// - `https-only` – CloudFront always uses HTTPS to connect to the origin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-originprotocolpolicy
	//
	OriginProtocolPolicy *string `field:"optional" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin.
	//
	// This is also known as the *origin response timeout* . The minimum timeout is 1 second, the maximum is 120 seconds, and the default (if you don't specify otherwise) is 30 seconds.
	//
	// For more information, see [Response timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#DownloadDistValuesOriginResponseTimeout) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-originreadtimeout
	//
	// Default: - 30.
	//
	OriginReadTimeout *float64 `field:"optional" json:"originReadTimeout" yaml:"originReadTimeout"`
	// Specifies the minimum SSL/TLS protocol that CloudFront uses when connecting to your origin over HTTPS.
	//
	// Valid values include `SSLv3` , `TLSv1` , `TLSv1.1` , and `TLSv1.2` .
	//
	// For more information, see [Minimum Origin SSL Protocol](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#DownloadDistValuesOriginSSLProtocols) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customoriginconfig.html#cfn-cloudfront-distribution-customoriginconfig-originsslprotocols
	//
	OriginSslProtocols *[]*string `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
}


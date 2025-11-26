package previewawscloudfrontmixins


// A custom origin.
//
// A custom origin is any origin that is *not* an Amazon S3 bucket, with one exception. An Amazon S3 bucket that is [configured with static website hosting](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) *is* a custom origin.
//
// > This property is legacy. We recommend that you use [Origin](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html) instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   legacyCustomOriginProperty := &LegacyCustomOriginProperty{
//   	DnsName: jsii.String("dnsName"),
//   	HttpPort: jsii.Number(123),
//   	HttpsPort: jsii.Number(123),
//   	OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   	OriginSslProtocols: []*string{
//   		jsii.String("originSslProtocols"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-legacycustomorigin.html
//
type CfnDistributionPropsMixin_LegacyCustomOriginProperty struct {
	// The domain name assigned to your CloudFront distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-legacycustomorigin.html#cfn-cloudfront-distribution-legacycustomorigin-dnsname
	//
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The HTTP port that CloudFront uses to connect to the origin.
	//
	// Specify the HTTP port that the origin listens on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-legacycustomorigin.html#cfn-cloudfront-distribution-legacycustomorigin-httpport
	//
	// Default: - 80.
	//
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The HTTPS port that CloudFront uses to connect to the origin.
	//
	// Specify the HTTPS port that the origin listens on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-legacycustomorigin.html#cfn-cloudfront-distribution-legacycustomorigin-httpsport
	//
	// Default: - 443.
	//
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-legacycustomorigin.html#cfn-cloudfront-distribution-legacycustomorigin-originprotocolpolicy
	//
	OriginProtocolPolicy *string `field:"optional" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// The minimum SSL/TLS protocol version that CloudFront uses when communicating with your origin server over HTTPs.
	//
	// For more information, see [Minimum Origin SSL Protocol](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginSSLProtocols) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-legacycustomorigin.html#cfn-cloudfront-distribution-legacycustomorigin-originsslprotocols
	//
	OriginSslProtocols *[]*string `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
}


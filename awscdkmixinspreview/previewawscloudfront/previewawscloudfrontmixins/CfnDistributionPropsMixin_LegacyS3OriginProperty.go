package previewawscloudfrontmixins


// The origin as an Amazon S3 bucket.
//
// > This property is legacy. We recommend that you use [Origin](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origin.html) instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   legacyS3OriginProperty := &LegacyS3OriginProperty{
//   	DnsName: jsii.String("dnsName"),
//   	OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-legacys3origin.html
//
type CfnDistributionPropsMixin_LegacyS3OriginProperty struct {
	// The domain name assigned to your CloudFront distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-legacys3origin.html#cfn-cloudfront-distribution-legacys3origin-dnsname
	//
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The CloudFront origin access identity to associate with the distribution.
	//
	// Use an origin access identity to configure the distribution so that end users can only access objects in an Amazon S3 through CloudFront .
	//
	// > This property is legacy. We recommend that you use [OriginAccessControl](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originaccesscontrol.html) instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-legacys3origin.html#cfn-cloudfront-distribution-legacys3origin-originaccessidentity
	//
	// Default: - "".
	//
	OriginAccessIdentity *string `field:"optional" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}


package awscloudfront


// An Amazon CloudFront VPC origin endpoint configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOriginEndpointConfigProperty := &VpcOriginEndpointConfigProperty{
//   	Arn: jsii.String("arn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	HttpPort: jsii.Number(123),
//   	HttpsPort: jsii.Number(123),
//   	OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   	OriginSslProtocols: []*string{
//   		jsii.String("originSslProtocols"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-vpcorigin-vpcoriginendpointconfig.html
//
type CfnVpcOrigin_VpcOriginEndpointConfigProperty struct {
	// The ARN of the CloudFront VPC origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-vpcorigin-vpcoriginendpointconfig.html#cfn-cloudfront-vpcorigin-vpcoriginendpointconfig-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The name of the CloudFront VPC origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-vpcorigin-vpcoriginendpointconfig.html#cfn-cloudfront-vpcorigin-vpcoriginendpointconfig-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The HTTP port for the CloudFront VPC origin endpoint configuration.
	//
	// The default value is `80` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-vpcorigin-vpcoriginendpointconfig.html#cfn-cloudfront-vpcorigin-vpcoriginendpointconfig-httpport
	//
	// Default: - 80.
	//
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The HTTPS port of the CloudFront VPC origin endpoint configuration.
	//
	// The default value is `443` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-vpcorigin-vpcoriginendpointconfig.html#cfn-cloudfront-vpcorigin-vpcoriginendpointconfig-httpsport
	//
	// Default: - 443.
	//
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// The origin protocol policy for the CloudFront VPC origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-vpcorigin-vpcoriginendpointconfig.html#cfn-cloudfront-vpcorigin-vpcoriginendpointconfig-originprotocolpolicy
	//
	// Default: - "match-viewer".
	//
	OriginProtocolPolicy *string `field:"optional" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// Specifies the minimum SSL/TLS protocol that CloudFront uses when connecting to your origin over HTTPS.
	//
	// Valid values include `SSLv3` , `TLSv1` , `TLSv1.1` , and `TLSv1.2` .
	//
	// For more information, see [Minimum Origin SSL Protocol](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#DownloadDistValuesOriginSSLProtocols) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-vpcorigin-vpcoriginendpointconfig.html#cfn-cloudfront-vpcorigin-vpcoriginendpointconfig-originsslprotocols
	//
	OriginSslProtocols *[]*string `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
}


package previewawscloudfrontmixins


// Configuration for an IPAM CIDR that defines a specific IP address range, IPAM pool, and associated Anycast IP address.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ipamCidrConfigProperty := &IpamCidrConfigProperty{
//   	Cidr: jsii.String("cidr"),
//   	IpamPoolArn: jsii.String("ipamPoolArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfig.html
//
type CfnAnycastIpListPropsMixin_IpamCidrConfigProperty struct {
	// The CIDR that specifies the IP address range for this IPAM configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfig.html#cfn-cloudfront-anycastiplist-ipamcidrconfig-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The Amazon Resource Name (ARN) of the IPAM pool that the CIDR block is assigned to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfig.html#cfn-cloudfront-anycastiplist-ipamcidrconfig-ipampoolarn
	//
	IpamPoolArn *string `field:"optional" json:"ipamPoolArn" yaml:"ipamPoolArn"`
}


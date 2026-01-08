package awscloudfront


// The result for the IPAM CIDR that defines a specific IP address range, IPAM pool, and associated Anycast IP address.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipamCidrConfigResultProperty := &IpamCidrConfigResultProperty{
//   	AnycastIp: jsii.String("anycastIp"),
//   	Cidr: jsii.String("cidr"),
//   	IpamPoolArn: jsii.String("ipamPoolArn"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfigresult.html
//
type CfnAnycastIpList_IpamCidrConfigResultProperty struct {
	// The specified Anycast IP address allocated from the IPAM pool for this CIDR configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfigresult.html#cfn-cloudfront-anycastiplist-ipamcidrconfigresult-anycastip
	//
	AnycastIp *string `field:"optional" json:"anycastIp" yaml:"anycastIp"`
	// The CIDR that specifies the IP address range for this IPAM configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfigresult.html#cfn-cloudfront-anycastiplist-ipamcidrconfigresult-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The Amazon Resource Name (ARN) of the IPAM pool that the CIDR block is assigned to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfigresult.html#cfn-cloudfront-anycastiplist-ipamcidrconfigresult-ipampoolarn
	//
	IpamPoolArn *string `field:"optional" json:"ipamPoolArn" yaml:"ipamPoolArn"`
	// The current status of the IPAM CIDR configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-anycastiplist-ipamcidrconfigresult.html#cfn-cloudfront-anycastiplist-ipamcidrconfigresult-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}


package awsec2


// Properties for defining a `CfnSubnetCidrBlock`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubnetCidrBlockProps := &CfnSubnetCidrBlockProps{
//   	Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html
//
type CfnSubnetCidrBlockProps struct {
	// The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length.
	//
	// This parameter is required for an IPv6 only subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html#cfn-ec2-subnetcidrblock-ipv6cidrblock
	//
	Ipv6CidrBlock *string `field:"required" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// The ID of the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html#cfn-ec2-subnetcidrblock-subnetid
	//
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}


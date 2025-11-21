package awsefs


// Properties for defining a `CfnMountTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMountTargetProps := &CfnMountTargetProps{
//   	FileSystemId: jsii.String("fileSystemId"),
//   	SecurityGroups: []interface{}{
//   		jsii.String("securityGroups"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	IpAddress: jsii.String("ipAddress"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Ipv6Address: jsii.String("ipv6Address"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html
//
type CfnMountTargetProps struct {
	// The ID of the file system for which to create the mount target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-filesystemid
	//
	FileSystemId interface{} `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// VPC security group IDs, of the form `sg-xxxxxxxx` .
	//
	// These must be for the same VPC as the subnet specified. The maximum number of security groups depends on account quota. For more information, see [Amazon VPC Quotas](https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html) in the *Amazon VPC User Guide* (see the *Security Groups* table). If you don't specify a security group, then Amazon EFS uses the default security group for the subnet's VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-securitygroups
	//
	SecurityGroups *[]interface{} `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The ID of the subnet to add the mount target in.
	//
	// For One Zone file systems, use the subnet that is associated with the file system's Availability Zone. The subnet type must be the same type as the `IpAddressType` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-subnetid
	//
	SubnetId interface{} `field:"required" json:"subnetId" yaml:"subnetId"`
	// If the `IpAddressType` for the mount target is IPv4 ( `IPV4_ONLY` or `DUAL_STACK` ), then specify the IPv4 address to use.
	//
	// If you do not specify an `IpAddress` , then Amazon EFS selects an unused IP address from the subnet specified for `SubnetId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-ipaddress
	//
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The IP address type for the mount target.
	//
	// The possible values are `IPV4_ONLY` (only IPv4 addresses), `IPV6_ONLY` (only IPv6 addresses), and `DUAL_STACK` (dual-stack, both IPv4 and IPv6 addresses). If you don’t specify an `IpAddressType` , then `IPV4_ONLY` is used.
	//
	// > The `IPAddressType` must match the IP type of the subnet. Additionally, the `IPAddressType` parameter overrides the value set as the default IP address for the subnet in the VPC. For example, if the `IPAddressType` is `IPV4_ONLY` and `AssignIpv6AddressOnCreation` is `true` , then IPv4 is used for the mount target. For more information, see [Modify the IP addressing attributes of your subnet](https://docs.aws.amazon.com/vpc/latest/userguide/subnet-public-ip.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// If the `IPAddressType` for the mount target is IPv6 ( `IPV6_ONLY` or `DUAL_STACK` ), then specify the IPv6 address to use.
	//
	// If you do not specify an `Ipv6Address` , then Amazon EFS selects an unused IP address from the subnet specified for `SubnetId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-ipv6address
	//
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
}


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
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	IpAddress: jsii.String("ipAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html
//
type CfnMountTargetProps struct {
	// The ID of the file system for which to create the mount target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-filesystemid
	//
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Up to five VPC security group IDs, of the form `sg-xxxxxxxx` .
	//
	// These must be for the same VPC as subnet specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-securitygroups
	//
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The ID of the subnet to add the mount target in.
	//
	// For One Zone file systems, use the subnet that is associated with the file system's Availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-subnetid
	//
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Valid IPv4 address within the address range of the specified subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-ipaddress
	//
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}


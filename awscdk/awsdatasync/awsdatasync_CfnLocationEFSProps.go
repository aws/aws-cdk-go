package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocationEFS`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationEFSProps := &cfnLocationEFSProps{
//   	ec2Config: &ec2ConfigProperty{
//   		securityGroupArns: []*string{
//   			jsii.String("securityGroupArns"),
//   		},
//   		subnetArn: jsii.String("subnetArn"),
//   	},
//   	efsFilesystemArn: jsii.String("efsFilesystemArn"),
//
//   	// the properties below are optional
//   	accessPointArn: jsii.String("accessPointArn"),
//   	fileSystemAccessRoleArn: jsii.String("fileSystemAccessRoleArn"),
//   	inTransitEncryption: jsii.String("inTransitEncryption"),
//   	subdirectory: jsii.String("subdirectory"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationEFSProps struct {
	// The subnet and security group that the Amazon EFS file system uses.
	//
	// The security group that you provide needs to be able to communicate with the security group on the mount target in the subnet specified.
	//
	// The exact relationship between security group M (of the mount target) and security group S (which you provide for DataSync to use at this stage) is as follows:
	//
	// - Security group M (which you associate with the mount target) must allow inbound access for the Transmission Control Protocol (TCP) on the NFS port (2049) from security group S. You can enable inbound connections either by IP address (CIDR range) or security group.
	// - Security group S (provided to DataSync to access EFS) should have a rule that enables outbound connections to the NFS port on one of the file system’s mount targets. You can enable outbound connections either by IP address (CIDR range) or security group.
	//
	// For information about security groups and mount targets, see [Security Groups for Amazon EC2 Instances and Mount Targets](https://docs.aws.amazon.com/efs/latest/ug/security-considerations.html#network-access) in the *Amazon EFS User Guide.*
	Ec2Config interface{} `field:"required" json:"ec2Config" yaml:"ec2Config"`
	// The Amazon Resource Name (ARN) for the Amazon EFS file system.
	EfsFilesystemArn *string `field:"required" json:"efsFilesystemArn" yaml:"efsFilesystemArn"`
	// `AWS::DataSync::LocationEFS.AccessPointArn`.
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// `AWS::DataSync::LocationEFS.FileSystemAccessRoleArn`.
	FileSystemAccessRoleArn *string `field:"optional" json:"fileSystemAccessRoleArn" yaml:"fileSystemAccessRoleArn"`
	// `AWS::DataSync::LocationEFS.InTransitEncryption`.
	InTransitEncryption *string `field:"optional" json:"inTransitEncryption" yaml:"inTransitEncryption"`
	// A subdirectory in the location’s path.
	//
	// This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination. By default, AWS DataSync uses the root directory.
	//
	// > `Subdirectory` must be specified with forward slashes. For example, `/path/to/folder` .
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents a tag that you want to add to the resource.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your resources. We recommend that you create a name tag for your location.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package mixinsawsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLocationEFSPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLocationEFSMixinProps := &CfnLocationEFSMixinProps{
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	Ec2Config: &Ec2ConfigProperty{
//   		SecurityGroupArns: []*string{
//   			jsii.String("securityGroupArns"),
//   		},
//   		SubnetArn: jsii.String("subnetArn"),
//   	},
//   	EfsFilesystemArn: jsii.String("efsFilesystemArn"),
//   	FileSystemAccessRoleArn: jsii.String("fileSystemAccessRoleArn"),
//   	InTransitEncryption: jsii.String("inTransitEncryption"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html
//
type CfnLocationEFSMixinProps struct {
	// Specifies the Amazon Resource Name (ARN) of the access point that DataSync uses to mount your Amazon EFS file system.
	//
	// For more information, see [Accessing restricted file systems](https://docs.aws.amazon.com/datasync/latest/userguide/create-efs-location.html#create-efs-location-iam) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-accesspointarn
	//
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// Specifies the subnet and security groups DataSync uses to connect to one of your Amazon EFS file system's [mount targets](https://docs.aws.amazon.com/efs/latest/ug/accessing-fs.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-ec2config
	//
	Ec2Config interface{} `field:"optional" json:"ec2Config" yaml:"ec2Config"`
	// Specifies the ARN for your Amazon EFS file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-efsfilesystemarn
	//
	EfsFilesystemArn *string `field:"optional" json:"efsFilesystemArn" yaml:"efsFilesystemArn"`
	// Specifies an AWS Identity and Access Management (IAM) role that allows DataSync to access your Amazon EFS file system.
	//
	// For information on creating this role, see [Creating a DataSync IAM role for file system access](https://docs.aws.amazon.com/datasync/latest/userguide/create-efs-location.html#create-efs-location-iam-role) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-filesystemaccessrolearn
	//
	FileSystemAccessRoleArn *string `field:"optional" json:"fileSystemAccessRoleArn" yaml:"fileSystemAccessRoleArn"`
	// Specifies whether you want DataSync to use Transport Layer Security (TLS) 1.2 encryption when it transfers data to or from your Amazon EFS file system.
	//
	// If you specify an access point using `AccessPointArn` or an IAM role using `FileSystemAccessRoleArn` , you must set this parameter to `TLS1_2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-intransitencryption
	//
	InTransitEncryption *string `field:"optional" json:"inTransitEncryption" yaml:"inTransitEncryption"`
	// Specifies a mount path for your Amazon EFS file system.
	//
	// This is where DataSync reads or writes data on your file system (depending on if this is a source or destination location).
	//
	// By default, DataSync uses the root directory (or [access point](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html) if you provide one by using `AccessPointArn` ). You can also include subdirectories using forward slashes (for example, `/path/to/folder` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies the key-value pair that represents a tag that you want to add to the resource.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your resources. We recommend that you create a name tag for your location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationefs.html#cfn-datasync-locationefs-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


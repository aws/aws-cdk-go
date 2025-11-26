package previewawsdatasyncmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLocationFSxOpenZFSPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLocationFSxOpenZFSMixinProps := &CfnLocationFSxOpenZFSMixinProps{
//   	FsxFilesystemArn: jsii.String("fsxFilesystemArn"),
//   	Protocol: &ProtocolProperty{
//   		Nfs: &NFSProperty{
//   			MountOptions: &MountOptionsProperty{
//   				Version: jsii.String("version"),
//   			},
//   		},
//   	},
//   	SecurityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxopenzfs.html
//
type CfnLocationFSxOpenZFSMixinProps struct {
	// The Amazon Resource Name (ARN) of the FSx for OpenZFS file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxopenzfs.html#cfn-datasync-locationfsxopenzfs-fsxfilesystemarn
	//
	FsxFilesystemArn *string `field:"optional" json:"fsxFilesystemArn" yaml:"fsxFilesystemArn"`
	// The type of protocol that AWS DataSync uses to access your file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxopenzfs.html#cfn-datasync-locationfsxopenzfs-protocol
	//
	Protocol interface{} `field:"optional" json:"protocol" yaml:"protocol"`
	// The ARNs of the security groups that are used to configure the FSx for OpenZFS file system.
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	//
	// *Length constraints* : Maximum length of 128.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxopenzfs.html#cfn-datasync-locationfsxopenzfs-securitygrouparns
	//
	SecurityGroupArns *[]*string `field:"optional" json:"securityGroupArns" yaml:"securityGroupArns"`
	// A subdirectory in the location's path that must begin with `/fsx` .
	//
	// DataSync uses this subdirectory to read or write data (depending on whether the file system is a source or destination location).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxopenzfs.html#cfn-datasync-locationfsxopenzfs-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents a tag that you want to add to the resource.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your resources. We recommend that you create a name tag for your location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxopenzfs.html#cfn-datasync-locationfsxopenzfs-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


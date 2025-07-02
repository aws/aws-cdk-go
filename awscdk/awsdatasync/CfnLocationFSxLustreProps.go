package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocationFSxLustre`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationFSxLustreProps := &CfnLocationFSxLustreProps{
//   	SecurityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//
//   	// the properties below are optional
//   	FsxFilesystemArn: jsii.String("fsxFilesystemArn"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxlustre.html
//
type CfnLocationFSxLustreProps struct {
	// The ARNs of the security groups that are used to configure the FSx for Lustre file system.
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	//
	// *Length constraints* : Maximum length of 128.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxlustre.html#cfn-datasync-locationfsxlustre-securitygrouparns
	//
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// Specifies the Amazon Resource Name (ARN) of the FSx for Lustre file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxlustre.html#cfn-datasync-locationfsxlustre-fsxfilesystemarn
	//
	FsxFilesystemArn *string `field:"optional" json:"fsxFilesystemArn" yaml:"fsxFilesystemArn"`
	// Specifies a mount path for your FSx for Lustre file system. The path can include subdirectories.
	//
	// When the location is used as a source, DataSync reads data from the mount path. When the location is used as a destination, DataSync writes data to the mount path. If you don't include this parameter, DataSync uses the file system's root directory ( `/` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxlustre.html#cfn-datasync-locationfsxlustre-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least a name tag for your location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxlustre.html#cfn-datasync-locationfsxlustre-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


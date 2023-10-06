package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocationFSxWindows`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationFSxWindowsProps := &CfnLocationFSxWindowsProps{
//   	SecurityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	User: jsii.String("user"),
//
//   	// the properties below are optional
//   	Domain: jsii.String("domain"),
//   	FsxFilesystemArn: jsii.String("fsxFilesystemArn"),
//   	Password: jsii.String("password"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxwindows.html
//
type CfnLocationFSxWindowsProps struct {
	// The Amazon Resource Names (ARNs) of the security groups that are used to configure the FSx for Windows File Server file system.
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	//
	// *Length constraints* : Maximum length of 128.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxwindows.html#cfn-datasync-locationfsxwindows-securitygrouparns
	//
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// The user who has the permissions to access files and folders in the FSx for Windows File Server file system.
	//
	// For information about choosing a user name that ensures sufficient permissions to files, folders, and metadata, see [user](https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html#FSxWuser) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxwindows.html#cfn-datasync-locationfsxwindows-user
	//
	User *string `field:"required" json:"user" yaml:"user"`
	// Specifies the name of the Windows domain that the FSx for Windows File Server belongs to.
	//
	// If you have multiple domains in your environment, configuring this parameter makes sure that DataSync connects to the right file server.
	//
	// For more information, see [required permissions](https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html#create-fsx-windows-location-permissions) for FSx for Windows File Server locations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxwindows.html#cfn-datasync-locationfsxwindows-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Specifies the Amazon Resource Name (ARN) for the FSx for Windows File Server file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxwindows.html#cfn-datasync-locationfsxwindows-fsxfilesystemarn
	//
	FsxFilesystemArn *string `field:"optional" json:"fsxFilesystemArn" yaml:"fsxFilesystemArn"`
	// Specifies the password of the user who has the permissions to access files and folders in the file system.
	//
	// For more information, see [required permissions](https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html#create-fsx-windows-location-permissions) for FSx for Windows File Server locations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxwindows.html#cfn-datasync-locationfsxwindows-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Specifies a mount path for your file system using forward slashes.
	//
	// This is where DataSync reads or writes data (depending on if this is a source or destination location).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxwindows.html#cfn-datasync-locationfsxwindows-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least a name tag for your location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxwindows.html#cfn-datasync-locationfsxwindows-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


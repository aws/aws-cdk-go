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
//   cfnLocationFSxWindowsProps := &cfnLocationFSxWindowsProps{
//   	securityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	user: jsii.String("user"),
//
//   	// the properties below are optional
//   	domain: jsii.String("domain"),
//   	fsxFilesystemArn: jsii.String("fsxFilesystemArn"),
//   	password: jsii.String("password"),
//   	subdirectory: jsii.String("subdirectory"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationFSxWindowsProps struct {
	// The Amazon Resource Names (ARNs) of the security groups that are used to configure the FSx for Windows File Server file system.
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	//
	// *Length constraints* : Maximum length of 128.
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// The user who has the permissions to access files and folders in the FSx for Windows File Server file system.
	//
	// For information about choosing a user name that ensures sufficient permissions to files, folders, and metadata, see [user](https://docs.aws.amazon.com/datasync/latest/userguide/create-fsx-location.html#FSxWuser) .
	User *string `field:"required" json:"user" yaml:"user"`
	// Specifies the name of the Windows domain that the FSx for Windows File Server belongs to.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Specifies the Amazon Resource Name (ARN) for the FSx for Windows File Server file system.
	FsxFilesystemArn *string `field:"optional" json:"fsxFilesystemArn" yaml:"fsxFilesystemArn"`
	// Specifies the password of the user who has the permissions to access files and folders in the file system.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Specifies a mount path for your file system using forward slashes.
	//
	// This is where DataSync reads or writes data (depending on if this is a source or destination location).
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least a name tag for your location.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


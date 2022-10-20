package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocationFSxWindows`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationFSxWindowsProps := &cfnLocationFSxWindowsProps{
//   	fsxFilesystemArn: jsii.String("fsxFilesystemArn"),
//   	password: jsii.String("password"),
//   	securityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	user: jsii.String("user"),
//
//   	// the properties below are optional
//   	domain: jsii.String("domain"),
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
	// The Amazon Resource Name (ARN) for the FSx for Windows File Server file system.
	FsxFilesystemArn *string `field:"required" json:"fsxFilesystemArn" yaml:"fsxFilesystemArn"`
	// The password of the user who has the permissions to access files and folders in the FSx for Windows File Server file system.
	Password *string `field:"required" json:"password" yaml:"password"`
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
	// The name of the Windows domain that the FSx for Windows File Server belongs to.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// A subdirectory in the location's path.
	//
	// This subdirectory in the Amazon FSx for Windows File Server file system is used to read data from the Amazon FSx for Windows File Server source location or write data to the FSx for Windows File Server destination.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents a tag that you want to add to the resource.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your resources. We recommend that you create a name tag for your location.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


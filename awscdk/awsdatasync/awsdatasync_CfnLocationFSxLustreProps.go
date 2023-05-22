package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
type CfnLocationFSxLustreProps struct {
	// The ARNs of the security groups that are used to configure the FSx for Lustre file system.
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	//
	// *Length constraints* : Maximum length of 128.
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// The Amazon Resource Name (ARN) for the FSx for Lustre file system.
	FsxFilesystemArn *string `field:"optional" json:"fsxFilesystemArn" yaml:"fsxFilesystemArn"`
	// A subdirectory in the location's path.
	//
	// This subdirectory in the FSx for Lustre file system is used to read data from the FSx for Lustre source location or write data to the FSx for Lustre destination.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents a tag that you want to add to the resource.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your resources. We recommend that you create a name tag for your location.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocationFSxOpenZFS`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationFSxOpenZFSProps := &cfnLocationFSxOpenZFSProps{
//   	fsxFilesystemArn: jsii.String("fsxFilesystemArn"),
//   	protocol: &protocolProperty{
//   		nfs: &nFSProperty{
//   			mountOptions: &mountOptionsProperty{
//   				version: jsii.String("version"),
//   			},
//   		},
//   	},
//   	securityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//
//   	// the properties below are optional
//   	subdirectory: jsii.String("subdirectory"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationFSxOpenZFSProps struct {
	// The Amazon Resource Name (ARN) of the FSx for OpenZFS file system.
	FsxFilesystemArn *string `field:"required" json:"fsxFilesystemArn" yaml:"fsxFilesystemArn"`
	// The type of protocol that AWS DataSync uses to access your file system.
	Protocol interface{} `field:"required" json:"protocol" yaml:"protocol"`
	// The ARNs of the security groups that are used to configure the FSx for OpenZFS file system.
	//
	// *Pattern* : `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`
	//
	// *Length constraints* : Maximum length of 128.
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// A subdirectory in the location's path that must begin with `/fsx` .
	//
	// DataSync uses this subdirectory to read or write data (depending on whether the file system is a source or destination location).
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents a tag that you want to add to the resource.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your resources. We recommend that you create a name tag for your location.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


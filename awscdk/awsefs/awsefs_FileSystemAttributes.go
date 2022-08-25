package awsefs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties that describe an existing EFS file system.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//
//   importedFileSystem := efs.fileSystem.fromFileSystemAttributes(this, jsii.String("existingFS"), &fileSystemAttributes{
//   	fileSystemId: jsii.String("fs-12345678"),
//   	 // You can also use fileSystemArn instead of fileSystemId.
//   	securityGroup: ec2.securityGroup.fromSecurityGroupId(this, jsii.String("SG"), jsii.String("sg-123456789"), &securityGroupImportOptions{
//   		allowAllOutbound: jsii.Boolean(false),
//   	}),
//   })
//
type FileSystemAttributes struct {
	// The security group of the file system.
	SecurityGroup awsec2.ISecurityGroup `field:"required" json:"securityGroup" yaml:"securityGroup"`
	// The File System's Arn.
	FileSystemArn *string `field:"optional" json:"fileSystemArn" yaml:"fileSystemArn"`
	// The File System's ID.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
}


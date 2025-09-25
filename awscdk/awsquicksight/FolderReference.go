package awsquicksight


// A reference to a Folder resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   folderReference := &FolderReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	FolderArn: jsii.String("folderArn"),
//   	FolderId: jsii.String("folderId"),
//   }
//
type FolderReference struct {
	// The AwsAccountId of the Folder resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ARN of the Folder resource.
	FolderArn *string `field:"required" json:"folderArn" yaml:"folderArn"`
	// The FolderId of the Folder resource.
	FolderId *string `field:"required" json:"folderId" yaml:"folderId"`
}


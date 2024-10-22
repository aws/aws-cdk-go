package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFolder`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFolderProps := &CfnFolderProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	FolderId: jsii.String("folderId"),
//   	FolderType: jsii.String("folderType"),
//   	Name: jsii.String("name"),
//   	ParentFolderArn: jsii.String("parentFolderArn"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   		},
//   	},
//   	SharingModel: jsii.String("sharingModel"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html
//
type CfnFolderProps struct {
	// The ID for the AWS account where you want to create the folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The ID of the folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-folderid
	//
	FolderId *string `field:"optional" json:"folderId" yaml:"folderId"`
	// The type of folder it is.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-foldertype
	//
	FolderType *string `field:"optional" json:"folderType" yaml:"folderType"`
	// A display name for the folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) for the folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-parentfolderarn
	//
	ParentFolderArn *string `field:"optional" json:"parentFolderArn" yaml:"parentFolderArn"`
	// A structure that describes the principals and the resource-level permissions of a folder.
	//
	// To specify no permissions, omit `Permissions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// The sharing scope of the folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-sharingmodel
	//
	SharingModel *string `field:"optional" json:"sharingModel" yaml:"sharingModel"`
	// A list of tags for the folders that you want to apply overrides to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-folder.html#cfn-quicksight-folder-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


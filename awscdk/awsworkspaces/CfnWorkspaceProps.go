package awsworkspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWorkspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspaceProps := &CfnWorkspaceProps{
//   	BundleId: jsii.String("bundleId"),
//   	DirectoryId: jsii.String("directoryId"),
//   	UserName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	RootVolumeEncryptionEnabled: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserVolumeEncryptionEnabled: jsii.Boolean(false),
//   	VolumeEncryptionKey: jsii.String("volumeEncryptionKey"),
//   	WorkspaceProperties: &WorkspacePropertiesProperty{
//   		ComputeTypeName: jsii.String("computeTypeName"),
//   		RootVolumeSizeGib: jsii.Number(123),
//   		RunningMode: jsii.String("runningMode"),
//   		RunningModeAutoStopTimeoutInMinutes: jsii.Number(123),
//   		UserVolumeSizeGib: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html
//
type CfnWorkspaceProps struct {
	// The identifier of the bundle for the WorkSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-bundleid
	//
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// The identifier of the AWS Directory Service directory for the WorkSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-directoryid
	//
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// The user name of the user for the WorkSpace.
	//
	// This user name must exist in the AWS Directory Service directory for the WorkSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-username
	//
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Indicates whether the data stored on the root volume is encrypted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-rootvolumeencryptionenabled
	//
	RootVolumeEncryptionEnabled interface{} `field:"optional" json:"rootVolumeEncryptionEnabled" yaml:"rootVolumeEncryptionEnabled"`
	// The tags for the WorkSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether the data stored on the user volume is encrypted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-uservolumeencryptionenabled
	//
	UserVolumeEncryptionEnabled interface{} `field:"optional" json:"userVolumeEncryptionEnabled" yaml:"userVolumeEncryptionEnabled"`
	// The symmetric AWS KMS key used to encrypt data stored on your WorkSpace.
	//
	// Amazon WorkSpaces does not support asymmetric KMS keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-volumeencryptionkey
	//
	VolumeEncryptionKey *string `field:"optional" json:"volumeEncryptionKey" yaml:"volumeEncryptionKey"`
	// The WorkSpace properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html#cfn-workspaces-workspace-workspaceproperties
	//
	WorkspaceProperties interface{} `field:"optional" json:"workspaceProperties" yaml:"workspaceProperties"`
}


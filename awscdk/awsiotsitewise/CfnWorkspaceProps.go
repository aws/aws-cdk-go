package awsiotsitewise

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
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   	},
//   	WorkspaceName: jsii.String("workspaceName"),
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkspaceDescription: jsii.String("workspaceDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-workspace.html
//
type CfnWorkspaceProps struct {
	// The encryption configuration for the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-workspace.html#cfn-iotsitewise-workspace-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The name of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-workspace.html#cfn-iotsitewise-workspace-workspacename
	//
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// The ARN of the AWS KMS key used for KMS_BASED_ENCRYPTION.
	//
	// Required when EncryptionConfiguration.EncryptionType is KMS_BASED_ENCRYPTION.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-workspace.html#cfn-iotsitewise-workspace-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-workspace.html#cfn-iotsitewise-workspace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A description of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-workspace.html#cfn-iotsitewise-workspace-workspacedescription
	//
	WorkspaceDescription *string `field:"optional" json:"workspaceDescription" yaml:"workspaceDescription"`
}


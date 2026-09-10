package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWorkspacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWorkspaceMixinProps := &CfnWorkspaceMixinProps{
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionType: jsii.String("encryptionType"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkspaceDescription: jsii.String("workspaceDescription"),
//   	WorkspaceName: jsii.String("workspaceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-workspace.html
//
type CfnWorkspaceMixinProps struct {
	// The encryption configuration for the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-workspace.html#cfn-iotsitewise-workspace-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
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
	// The name of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-workspace.html#cfn-iotsitewise-workspace-workspacename
	//
	WorkspaceName *string `field:"optional" json:"workspaceName" yaml:"workspaceName"`
}


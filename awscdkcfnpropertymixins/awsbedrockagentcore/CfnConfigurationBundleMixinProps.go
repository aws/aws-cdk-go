package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConfigurationBundlePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var configuration interface{}
//
//   cfnConfigurationBundleMixinProps := &CfnConfigurationBundleMixinProps{
//   	BranchName: jsii.String("branchName"),
//   	BundleName: jsii.String("bundleName"),
//   	CommitMessage: jsii.String("commitMessage"),
//   	Components: map[string]interface{}{
//   		"componentsKey": &ComponentConfigurationProperty{
//   			"configuration": configuration,
//   		},
//   	},
//   	CreatedBy: &VersionCreatedBySourceProperty{
//   		Arn: jsii.String("arn"),
//   		Name: jsii.String("name"),
//   	},
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html
//
type CfnConfigurationBundleMixinProps struct {
	// The branch name for version tracking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-branchname
	//
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// The name for the configuration bundle.
	//
	// Names must be unique within your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-bundlename
	//
	BundleName *string `field:"optional" json:"bundleName" yaml:"bundleName"`
	// A commit message describing the version of the configuration bundle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-commitmessage
	//
	CommitMessage *string `field:"optional" json:"commitMessage" yaml:"commitMessage"`
	// A map of component identifiers to their configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-components
	//
	Components interface{} `field:"optional" json:"components" yaml:"components"`
	// The source that created a configuration bundle version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-createdby
	//
	CreatedBy interface{} `field:"optional" json:"createdBy" yaml:"createdBy"`
	// The description for the configuration bundle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to assign to the configuration bundle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


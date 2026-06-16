package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConfigurationBundle`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   cfnConfigurationBundleProps := &CfnConfigurationBundleProps{
//   	BundleName: jsii.String("bundleName"),
//   	Components: map[string]interface{}{
//   		"componentsKey": &ComponentConfigurationProperty{
//   			"configuration": configuration,
//   		},
//   	},
//
//   	// the properties below are optional
//   	BranchName: jsii.String("branchName"),
//   	CommitMessage: jsii.String("commitMessage"),
//   	CreatedBy: &VersionCreatedBySourceProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Arn: jsii.String("arn"),
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
type CfnConfigurationBundleProps struct {
	// The name for the configuration bundle.
	//
	// Names must be unique within your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-bundlename
	//
	BundleName *string `field:"required" json:"bundleName" yaml:"bundleName"`
	// A map of component identifiers to their configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-components
	//
	Components interface{} `field:"required" json:"components" yaml:"components"`
	// The branch name for version tracking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-branchname
	//
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// A commit message describing the version of the configuration bundle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-configurationbundle.html#cfn-bedrockagentcore-configurationbundle-commitmessage
	//
	CommitMessage *string `field:"optional" json:"commitMessage" yaml:"commitMessage"`
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


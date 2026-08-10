package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnActionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnActionMixinProps := &CfnActionMixinProps{
//   	ActionName: jsii.String("actionName"),
//   	ActionType: jsii.String("actionType"),
//   	Description: jsii.String("description"),
//   	MetadataProperties: &MetadataPropertiesProperty{
//   		CommitId: jsii.String("commitId"),
//   		GeneratedBy: jsii.String("generatedBy"),
//   		ProjectId: jsii.String("projectId"),
//   		Repository: jsii.String("repository"),
//   	},
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   	Source: &ActionSourceProperty{
//   		SourceId: jsii.String("sourceId"),
//   		SourceType: jsii.String("sourceType"),
//   		SourceUri: jsii.String("sourceUri"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html
//
type CfnActionMixinProps struct {
	// The name of the action.
	//
	// Must be unique to your account in an AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-actionname
	//
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// The action type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-actiontype
	//
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// The description of the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-metadataproperties
	//
	MetadataProperties interface{} `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// A list of properties to add to the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The status of the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of tags to apply to the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnActionProps := &CfnActionProps{
//   	ActionName: jsii.String("actionName"),
//   	ActionType: jsii.String("actionType"),
//   	Source: &ActionSourceProperty{
//   		SourceUri: jsii.String("sourceUri"),
//
//   		// the properties below are optional
//   		SourceId: jsii.String("sourceId"),
//   		SourceType: jsii.String("sourceType"),
//   	},
//
//   	// the properties below are optional
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
type CfnActionProps struct {
	// The name of the action.
	//
	// Must be unique to your account in an AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-actionname
	//
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The action type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-actiontype
	//
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
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
	// The status of the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of tags to apply to the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-action.html#cfn-sagemaker-action-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


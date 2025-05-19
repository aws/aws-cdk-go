package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIntelligentPromptRouter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntelligentPromptRouterProps := &CfnIntelligentPromptRouterProps{
//   	FallbackModel: &PromptRouterTargetModelProperty{
//   		ModelArn: jsii.String("modelArn"),
//   	},
//   	Models: []interface{}{
//   		&PromptRouterTargetModelProperty{
//   			ModelArn: jsii.String("modelArn"),
//   		},
//   	},
//   	PromptRouterName: jsii.String("promptRouterName"),
//   	RoutingCriteria: &RoutingCriteriaProperty{
//   		ResponseQualityDifference: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html
//
type CfnIntelligentPromptRouterProps struct {
	// Model configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-fallbackmodel
	//
	FallbackModel interface{} `field:"required" json:"fallbackModel" yaml:"fallbackModel"`
	// List of model configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-models
	//
	Models interface{} `field:"required" json:"models" yaml:"models"`
	// Name of the Prompt Router.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-promptroutername
	//
	PromptRouterName *string `field:"required" json:"promptRouterName" yaml:"promptRouterName"`
	// Routing criteria for a prompt router.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-routingcriteria
	//
	RoutingCriteria interface{} `field:"required" json:"routingCriteria" yaml:"routingCriteria"`
	// Description of the Prompt Router.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of Tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


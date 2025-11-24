package mixinsawsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIntelligentPromptRouterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIntelligentPromptRouterMixinProps := &CfnIntelligentPromptRouterMixinProps{
//   	Description: jsii.String("description"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html
//
type CfnIntelligentPromptRouterMixinProps struct {
	// An optional description of the prompt router to help identify its purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The default model to use when the routing criteria is not met.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-fallbackmodel
	//
	FallbackModel interface{} `field:"optional" json:"fallbackModel" yaml:"fallbackModel"`
	// A list of foundation models that the prompt router can route requests to.
	//
	// At least one model must be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-models
	//
	Models interface{} `field:"optional" json:"models" yaml:"models"`
	// The name of the prompt router.
	//
	// The name must be unique within your AWS account in the current region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-promptroutername
	//
	PromptRouterName *string `field:"optional" json:"promptRouterName" yaml:"promptRouterName"`
	// Routing criteria for a prompt router.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-routingcriteria
	//
	RoutingCriteria interface{} `field:"optional" json:"routingCriteria" yaml:"routingCriteria"`
	// An array of key-value pairs to apply to this resource as tags.
	//
	// You can use tags to categorize and manage your AWS resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-intelligentpromptrouter.html#cfn-bedrock-intelligentpromptrouter-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


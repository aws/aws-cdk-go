package awsbedrock


// The target model for a prompt router.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptRouterTargetModelProperty := &PromptRouterTargetModelProperty{
//   	ModelArn: jsii.String("modelArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-intelligentpromptrouter-promptroutertargetmodel.html
//
type CfnIntelligentPromptRouter_PromptRouterTargetModelProperty struct {
	// The target model's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-intelligentpromptrouter-promptroutertargetmodel.html#cfn-bedrock-intelligentpromptrouter-promptroutertargetmodel-modelarn
	//
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
}


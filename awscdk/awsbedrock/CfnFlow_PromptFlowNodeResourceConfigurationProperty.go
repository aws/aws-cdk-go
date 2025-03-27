package awsbedrock


// Contains configurations for a prompt from Prompt management to use in a node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptFlowNodeResourceConfigurationProperty := &PromptFlowNodeResourceConfigurationProperty{
//   	PromptArn: jsii.String("promptArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptflownoderesourceconfiguration.html
//
type CfnFlow_PromptFlowNodeResourceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the prompt from Prompt management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptflownoderesourceconfiguration.html#cfn-bedrock-flow-promptflownoderesourceconfiguration-promptarn
	//
	PromptArn *string `field:"required" json:"promptArn" yaml:"promptArn"`
}


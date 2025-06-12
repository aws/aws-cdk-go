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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownoderesourceconfiguration.html
//
type CfnFlowVersion_PromptFlowNodeResourceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the prompt from Prompt management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownoderesourceconfiguration.html#cfn-bedrock-flowversion-promptflownoderesourceconfiguration-promptarn
	//
	PromptArn *string `field:"required" json:"promptArn" yaml:"promptArn"`
}


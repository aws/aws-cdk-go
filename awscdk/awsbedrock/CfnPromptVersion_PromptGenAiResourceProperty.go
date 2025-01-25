package awsbedrock


// Target resource to invoke with Prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptGenAiResourceProperty := &PromptGenAiResourceProperty{
//   	Agent: &PromptAgentResourceProperty{
//   		AgentIdentifier: jsii.String("agentIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptgenairesource.html
//
type CfnPromptVersion_PromptGenAiResourceProperty struct {
	// Target Agent to invoke with Prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptgenairesource.html#cfn-bedrock-promptversion-promptgenairesource-agent
	//
	Agent interface{} `field:"required" json:"agent" yaml:"agent"`
}


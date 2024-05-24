package awsbedrock


// Defines parameters that the agent needs to invoke from the user to complete the function.
//
// Corresponds to an action in an action group.
//
// This data type is used in the following API operations:
//
// - [CreateAgentActionGroup request](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_CreateAgentActionGroup.html#API_agent_CreateAgentActionGroup_RequestSyntax)
// - [CreateAgentActionGroup response](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_CreateAgentActionGroup.html#API_agent_CreateAgentActionGroup_ResponseSyntax)
// - [UpdateAgentActionGroup request](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_UpdateAgentActionGroup.html#API_agent_UpdateAgentActionGroup_RequestSyntax)
// - [UpdateAgentActionGroup response](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_UpdateAgentActionGroup.html#API_agent_UpdateAgentActionGroup_ResponseSyntax)
// - [GetAgentActionGroup response](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_GetAgentActionGroup.html#API_agent_GetAgentActionGroup_ResponseSyntax)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionProperty := &FunctionProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Parameters: map[string]interface{}{
//   		"parametersKey": &ParameterDetailProperty{
//   			"type": jsii.String("type"),
//
//   			// the properties below are optional
//   			"description": jsii.String("description"),
//   			"required": jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-function.html
//
type CfnAgent_FunctionProperty struct {
	// A name for the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-function.html#cfn-bedrock-agent-function-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the function and its purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-function.html#cfn-bedrock-agent-function-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The parameters that the agent elicits from the user to fulfill the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-function.html#cfn-bedrock-agent-function-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}


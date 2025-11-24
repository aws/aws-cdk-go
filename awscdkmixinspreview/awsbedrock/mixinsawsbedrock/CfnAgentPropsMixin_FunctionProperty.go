package mixinsawsbedrock


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
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   functionProperty := &FunctionProperty{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Parameters: map[string]interface{}{
//   		"parametersKey": &ParameterDetailProperty{
//   			"description": jsii.String("description"),
//   			"required": jsii.Boolean(false),
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	RequireConfirmation: jsii.String("requireConfirmation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-function.html
//
type CfnAgentPropsMixin_FunctionProperty struct {
	// A description of the function and its purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-function.html#cfn-bedrock-agent-function-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-function.html#cfn-bedrock-agent-function-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameters that the agent elicits from the user to fulfill the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-function.html#cfn-bedrock-agent-function-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Contains information if user confirmation is required to invoke the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-function.html#cfn-bedrock-agent-function-requireconfirmation
	//
	RequireConfirmation *string `field:"optional" json:"requireConfirmation" yaml:"requireConfirmation"`
}


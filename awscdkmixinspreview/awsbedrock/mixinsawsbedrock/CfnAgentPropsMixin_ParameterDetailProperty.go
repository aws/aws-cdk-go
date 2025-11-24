package mixinsawsbedrock


// Contains details about a parameter in a function for an action group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parameterDetailProperty := &ParameterDetailProperty{
//   	Description: jsii.String("description"),
//   	Required: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-parameterdetail.html
//
type CfnAgentPropsMixin_ParameterDetailProperty struct {
	// A description of the parameter.
	//
	// Helps the foundation model determine how to elicit the parameters from the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-parameterdetail.html#cfn-bedrock-agent-parameterdetail-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the parameter is required for the agent to complete the function for action group invocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-parameterdetail.html#cfn-bedrock-agent-parameterdetail-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The data type of the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-parameterdetail.html#cfn-bedrock-agent-parameterdetail-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


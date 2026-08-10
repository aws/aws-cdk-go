package awsquicksight


// Parameters for creating a new custom prompt configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPromptInputParametersProperty := &CustomPromptInputParametersProperty{
//   	CustomInstructions: jsii.String("customInstructions"),
//   	Identity: jsii.String("identity"),
//   	OutputStyle: jsii.String("outputStyle"),
//   	ResponseLength: jsii.String("responseLength"),
//   	Tone: jsii.String("tone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinputparameters.html
//
type CfnAgent_CustomPromptInputParametersProperty struct {
	// Custom instructions for the agent behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinputparameters.html#cfn-quicksight-agent-custompromptinputparameters-custominstructions
	//
	CustomInstructions *string `field:"optional" json:"customInstructions" yaml:"customInstructions"`
	// The identity or persona of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinputparameters.html#cfn-quicksight-agent-custompromptinputparameters-identity
	//
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// The output style for the agent responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinputparameters.html#cfn-quicksight-agent-custompromptinputparameters-outputstyle
	//
	OutputStyle *string `field:"optional" json:"outputStyle" yaml:"outputStyle"`
	// The desired response length for the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinputparameters.html#cfn-quicksight-agent-custompromptinputparameters-responselength
	//
	ResponseLength *string `field:"optional" json:"responseLength" yaml:"responseLength"`
	// The tone used in agent responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinputparameters.html#cfn-quicksight-agent-custompromptinputparameters-tone
	//
	Tone *string `field:"optional" json:"tone" yaml:"tone"`
}


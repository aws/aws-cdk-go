package awsquicksight


// Read-only view of the resolved custom prompt interface for the agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   customPromptInterfaceProperty := &CustomPromptInterfaceProperty{
//   	CustomInstructions: jsii.String("customInstructions"),
//   	Identity: jsii.String("identity"),
//   	ModelProfileId: jsii.String("modelProfileId"),
//   	OutputStyle: jsii.String("outputStyle"),
//   	PromptSummary: jsii.String("promptSummary"),
//   	QbsAwsAccountId: jsii.String("qbsAwsAccountId"),
//   	ResponseLength: jsii.String("responseLength"),
//   	SubscriptionId: jsii.String("subscriptionId"),
//   	Tone: jsii.String("tone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinterface.html
//
type CfnAgentPropsMixin_CustomPromptInterfaceProperty struct {
	// Custom instructions for the agent behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinterface.html#cfn-quicksight-agent-custompromptinterface-custominstructions
	//
	CustomInstructions *string `field:"optional" json:"customInstructions" yaml:"customInstructions"`
	// The identity or persona of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinterface.html#cfn-quicksight-agent-custompromptinterface-identity
	//
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// The identifier of the model profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinterface.html#cfn-quicksight-agent-custompromptinterface-modelprofileid
	//
	ModelProfileId *string `field:"optional" json:"modelProfileId" yaml:"modelProfileId"`
	// The output style for the agent responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinterface.html#cfn-quicksight-agent-custompromptinterface-outputstyle
	//
	OutputStyle *string `field:"optional" json:"outputStyle" yaml:"outputStyle"`
	// A summary of the resolved prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinterface.html#cfn-quicksight-agent-custompromptinterface-promptsummary
	//
	PromptSummary *string `field:"optional" json:"promptSummary" yaml:"promptSummary"`
	// The QBS AWS account identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinterface.html#cfn-quicksight-agent-custompromptinterface-qbsawsaccountid
	//
	QbsAwsAccountId *string `field:"optional" json:"qbsAwsAccountId" yaml:"qbsAwsAccountId"`
	// The desired response length for the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinterface.html#cfn-quicksight-agent-custompromptinterface-responselength
	//
	ResponseLength *string `field:"optional" json:"responseLength" yaml:"responseLength"`
	// The subscription identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinterface.html#cfn-quicksight-agent-custompromptinterface-subscriptionid
	//
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// The tone used in agent responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinterface.html#cfn-quicksight-agent-custompromptinterface-tone
	//
	Tone *string `field:"optional" json:"tone" yaml:"tone"`
}


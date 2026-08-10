package awsquicksight


// Custom prompt configuration.
//
// Specify either ExistingPrompt or NewPrompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPromptInputProperty := &CustomPromptInputProperty{
//   	ExistingPrompt: &CustomPromptProfileProperty{
//   		ModelProfileId: jsii.String("modelProfileId"),
//   		QbsAwsAccountId: jsii.String("qbsAwsAccountId"),
//   		SubscriptionId: jsii.String("subscriptionId"),
//   	},
//   	NewPrompt: &CustomPromptInputParametersProperty{
//   		CustomInstructions: jsii.String("customInstructions"),
//   		Identity: jsii.String("identity"),
//   		OutputStyle: jsii.String("outputStyle"),
//   		ResponseLength: jsii.String("responseLength"),
//   		Tone: jsii.String("tone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinput.html
//
type CfnAgent_CustomPromptInputProperty struct {
	// Reference to an existing custom prompt profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinput.html#cfn-quicksight-agent-custompromptinput-existingprompt
	//
	ExistingPrompt interface{} `field:"optional" json:"existingPrompt" yaml:"existingPrompt"`
	// Parameters for creating a new custom prompt configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptinput.html#cfn-quicksight-agent-custompromptinput-newprompt
	//
	NewPrompt interface{} `field:"optional" json:"newPrompt" yaml:"newPrompt"`
}


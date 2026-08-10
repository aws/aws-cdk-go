package awsquicksight


// Properties for CfnAgentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAgentMixinProps := &CfnAgentMixinProps{
//   	ActionConnectors: []*string{
//   		jsii.String("actionConnectors"),
//   	},
//   	AgentId: jsii.String("agentId"),
//   	AgentLifecycle: jsii.String("agentLifecycle"),
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	CustomPromptInput: &CustomPromptInputProperty{
//   		ExistingPrompt: &CustomPromptProfileProperty{
//   			ModelProfileId: jsii.String("modelProfileId"),
//   			QbsAwsAccountId: jsii.String("qbsAwsAccountId"),
//   			SubscriptionId: jsii.String("subscriptionId"),
//   		},
//   		NewPrompt: &CustomPromptInputParametersProperty{
//   			CustomInstructions: jsii.String("customInstructions"),
//   			Identity: jsii.String("identity"),
//   			OutputStyle: jsii.String("outputStyle"),
//   			ResponseLength: jsii.String("responseLength"),
//   			Tone: jsii.String("tone"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	IconId: jsii.String("iconId"),
//   	Name: jsii.String("name"),
//   	Spaces: []*string{
//   		jsii.String("spaces"),
//   	},
//   	StarterPrompts: []*string{
//   		jsii.String("starterPrompts"),
//   	},
//   	Tags: []AgentTagProperty{
//   		&AgentTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WelcomeMessage: jsii.String("welcomeMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html
//
type CfnAgentMixinProps struct {
	// A list of ActionConnector ARNs (max 10) attached to the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-actionconnectors
	//
	ActionConnectors *[]*string `field:"optional" json:"actionConnectors" yaml:"actionConnectors"`
	// The unique identifier for the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-agentid
	//
	AgentId *string `field:"optional" json:"agentId" yaml:"agentId"`
	// The lifecycle stage of the agent.
	//
	// PREVIEW or PUBLISHED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-agentlifecycle
	//
	AgentLifecycle *string `field:"optional" json:"agentLifecycle" yaml:"agentLifecycle"`
	// The ID of the Amazon Web Services account where the agent is being created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Custom prompt configuration.
	//
	// Specify either ExistingPrompt or NewPrompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-custompromptinput
	//
	CustomPromptInput interface{} `field:"optional" json:"customPromptInput" yaml:"customPromptInput"`
	// A description of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The icon identifier for the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-iconid
	//
	IconId *string `field:"optional" json:"iconId" yaml:"iconId"`
	// The display name of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of Space ARNs (max 10) attached to the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-spaces
	//
	Spaces *[]*string `field:"optional" json:"spaces" yaml:"spaces"`
	// A list of up to 3 starter prompts displayed to users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-starterprompts
	//
	StarterPrompts *[]*string `field:"optional" json:"starterPrompts" yaml:"starterPrompts"`
	// A list of key-value pairs to associate with the agent resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-tags
	//
	Tags *[]*CfnAgentPropsMixin_AgentTagProperty `field:"optional" json:"tags" yaml:"tags"`
	// The welcome message displayed when a user opens the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-agent.html#cfn-quicksight-agent-welcomemessage
	//
	WelcomeMessage *string `field:"optional" json:"welcomeMessage" yaml:"welcomeMessage"`
}


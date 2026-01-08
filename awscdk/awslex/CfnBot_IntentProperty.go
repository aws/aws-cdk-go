package awslex


// Represents an action that the user wants to perform.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html
//
type CfnBot_IntentProperty struct {
	// The name of the intent.
	//
	// Intent names must be unique within the locale that contains the intent and can't match the name of any built-in intent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-bedrockagentintentconfiguration
	//
	BedrockAgentIntentConfiguration interface{} `field:"optional" json:"bedrockAgentIntentConfiguration" yaml:"bedrockAgentIntentConfiguration"`
	// A description of the intent.
	//
	// Use the description to help identify the intent in lists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies that Amazon Lex invokes the alias Lambda function for each user input.
	//
	// You can invoke this Lambda function to personalize user interaction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-dialogcodehook
	//
	DialogCodeHook interface{} `field:"optional" json:"dialogCodeHook" yaml:"dialogCodeHook"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Specifies that Amazon Lex invokes the alias Lambda function when the intent is ready for fulfillment.
	//
	// You can invoke this function to complete the bot's transaction with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-fulfillmentcodehook
	//
	FulfillmentCodeHook interface{} `field:"optional" json:"fulfillmentCodeHook" yaml:"fulfillmentCodeHook"`
	// Configuration setting for a response sent to the user before Amazon Lex starts eliciting slots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-initialresponsesetting
	//
	InitialResponseSetting interface{} `field:"optional" json:"initialResponseSetting" yaml:"initialResponseSetting"`
	// A list of contexts that must be active for this intent to be considered by Amazon Lex .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-inputcontexts
	//
	InputContexts interface{} `field:"optional" json:"inputContexts" yaml:"inputContexts"`
	// Sets the response that Amazon Lex sends to the user when the intent is closed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-intentclosingsetting
	//
	IntentClosingSetting interface{} `field:"optional" json:"intentClosingSetting" yaml:"intentClosingSetting"`
	// Provides prompts that Amazon Lex sends to the user to confirm the completion of an intent.
	//
	// If the user answers "no," the settings contain a statement that is sent to the user to end the intent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-intentconfirmationsetting
	//
	IntentConfirmationSetting interface{} `field:"optional" json:"intentConfirmationSetting" yaml:"intentConfirmationSetting"`
	// Provides configuration information for the `AMAZON.KendraSearchIntent` intent. When you use this intent, Amazon Lex searches the specified Amazon Kendra index and returns documents from the index that match the user's utterance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-kendraconfiguration
	//
	KendraConfiguration interface{} `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// A list of contexts that the intent activates when it is fulfilled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-outputcontexts
	//
	OutputContexts interface{} `field:"optional" json:"outputContexts" yaml:"outputContexts"`
	// A unique identifier for the built-in intent to base this intent on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-parentintentsignature
	//
	ParentIntentSignature *string `field:"optional" json:"parentIntentSignature" yaml:"parentIntentSignature"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-qinconnectintentconfiguration
	//
	QInConnectIntentConfiguration interface{} `field:"optional" json:"qInConnectIntentConfiguration" yaml:"qInConnectIntentConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-qnaintentconfiguration
	//
	QnAIntentConfiguration interface{} `field:"optional" json:"qnAIntentConfiguration" yaml:"qnAIntentConfiguration"`
	// A list of utterances that a user might say to signal the intent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-sampleutterances
	//
	SampleUtterances interface{} `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Indicates the priority for slots.
	//
	// Amazon Lex prompts the user for slot values in priority order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-slotpriorities
	//
	SlotPriorities interface{} `field:"optional" json:"slotPriorities" yaml:"slotPriorities"`
	// A list of slots that the intent requires for fulfillment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intent.html#cfn-lex-bot-intent-slots
	//
	Slots interface{} `field:"optional" json:"slots" yaml:"slots"`
}


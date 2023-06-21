package awslex


// Represents an action that the user wants to perform.
//
// Example:
//
//
type CfnBot_IntentProperty struct {
	// The name of the intent.
	//
	// Intent names must be unique within the locale that contains the intent and can't match the name of any built-in intent.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the intent.
	//
	// Use the description to help identify the intent in lists.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies that Amazon Lex invokes the alias Lambda function for each user input.
	//
	// You can invoke this Lambda function to personalize user interaction.
	DialogCodeHook interface{} `field:"optional" json:"dialogCodeHook" yaml:"dialogCodeHook"`
	// Specifies that Amazon Lex invokes the alias Lambda function when the intent is ready for fulfillment.
	//
	// You can invoke this function to complete the bot's transaction with the user.
	FulfillmentCodeHook interface{} `field:"optional" json:"fulfillmentCodeHook" yaml:"fulfillmentCodeHook"`
	// Configuration setting for a response sent to the user before Amazon Lex starts eliciting slots.
	InitialResponseSetting interface{} `field:"optional" json:"initialResponseSetting" yaml:"initialResponseSetting"`
	// A list of contexts that must be active for this intent to be considered by Amazon Lex .
	InputContexts interface{} `field:"optional" json:"inputContexts" yaml:"inputContexts"`
	// Sets the response that Amazon Lex sends to the user when the intent is closed.
	IntentClosingSetting interface{} `field:"optional" json:"intentClosingSetting" yaml:"intentClosingSetting"`
	// Provides prompts that Amazon Lex sends to the user to confirm the completion of an intent.
	//
	// If the user answers "no," the settings contain a statement that is sent to the user to end the intent.
	IntentConfirmationSetting interface{} `field:"optional" json:"intentConfirmationSetting" yaml:"intentConfirmationSetting"`
	// Provides configuration information for the `AMAZON.KendraSearchIntent` intent. When you use this intent, Amazon Lex searches the specified Amazon Kendra index and returns documents from the index that match the user's utterance.
	KendraConfiguration interface{} `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// A list of contexts that the intent activates when it is fulfilled.
	OutputContexts interface{} `field:"optional" json:"outputContexts" yaml:"outputContexts"`
	// A unique identifier for the built-in intent to base this intent on.
	ParentIntentSignature *string `field:"optional" json:"parentIntentSignature" yaml:"parentIntentSignature"`
	// A list of utterances that a user might say to signal the intent.
	SampleUtterances interface{} `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Indicates the priority for slots.
	//
	// Amazon Lex prompts the user for slot values in priority order.
	SlotPriorities interface{} `field:"optional" json:"slotPriorities" yaml:"slotPriorities"`
	// A list of slots that the intent requires for fulfillment.
	Slots interface{} `field:"optional" json:"slots" yaml:"slots"`
}


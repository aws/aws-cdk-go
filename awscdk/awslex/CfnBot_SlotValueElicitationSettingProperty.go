package awslex


// Specifies the elicitation setting details eliciting a slot.
//
// Example:
//
//
type CfnBot_SlotValueElicitationSettingProperty struct {
	// Specifies whether the slot is required or optional.
	SlotConstraint *string `field:"required" json:"slotConstraint" yaml:"slotConstraint"`
	// A list of default values for a slot.
	//
	// Default values are used when Amazon Lex hasn't determined a value for a slot. You can specify default values from context variables, session attributes, and defined values.
	DefaultValueSpecification interface{} `field:"optional" json:"defaultValueSpecification" yaml:"defaultValueSpecification"`
	// The prompt that Amazon Lex uses to elicit the slot value from the user.
	PromptSpecification interface{} `field:"optional" json:"promptSpecification" yaml:"promptSpecification"`
	// If you know a specific pattern that users might respond to an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy.
	//
	// This is optional. In most cases, Amazon Lex is capable of understanding user utterances.
	SampleUtterances interface{} `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Specifies the settings that Amazon Lex uses when a slot value is successfully entered by a user.
	SlotCaptureSetting interface{} `field:"optional" json:"slotCaptureSetting" yaml:"slotCaptureSetting"`
	// Specifies the prompts that Amazon Lex uses while a bot is waiting for customer input.
	WaitAndContinueSpecification interface{} `field:"optional" json:"waitAndContinueSpecification" yaml:"waitAndContinueSpecification"`
}


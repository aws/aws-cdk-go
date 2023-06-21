package awslex


// Settings used when Amazon Lex successfully captures a slot value from a user.
//
// Example:
//
//
type CfnBot_SlotCaptureSettingProperty struct {
	// A list of conditional branches to evaluate after the slot value is captured.
	CaptureConditional interface{} `field:"optional" json:"captureConditional" yaml:"captureConditional"`
	// Specifies the next step that the bot runs when the slot value is captured before the code hook times out.
	CaptureNextStep interface{} `field:"optional" json:"captureNextStep" yaml:"captureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input.
	CaptureResponse interface{} `field:"optional" json:"captureResponse" yaml:"captureResponse"`
	// Code hook called after Amazon Lex successfully captures a slot value.
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// Code hook called when Amazon Lex doesn't capture a slot value.
	ElicitationCodeHook interface{} `field:"optional" json:"elicitationCodeHook" yaml:"elicitationCodeHook"`
	// A list of conditional branches to evaluate when the slot value isn't captured.
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// Specifies the next step that the bot runs when the slot value code is not recognized.
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input when the slot fails to be captured.
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
}


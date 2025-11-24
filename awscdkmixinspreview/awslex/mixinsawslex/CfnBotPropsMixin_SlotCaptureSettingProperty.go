package mixinsawslex


// Settings used when Amazon Lex successfully captures a slot value from a user.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotcapturesetting.html
//
type CfnBotPropsMixin_SlotCaptureSettingProperty struct {
	// A list of conditional branches to evaluate after the slot value is captured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotcapturesetting.html#cfn-lex-bot-slotcapturesetting-captureconditional
	//
	CaptureConditional interface{} `field:"optional" json:"captureConditional" yaml:"captureConditional"`
	// Specifies the next step that the bot runs when the slot value is captured before the code hook times out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotcapturesetting.html#cfn-lex-bot-slotcapturesetting-capturenextstep
	//
	CaptureNextStep interface{} `field:"optional" json:"captureNextStep" yaml:"captureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotcapturesetting.html#cfn-lex-bot-slotcapturesetting-captureresponse
	//
	CaptureResponse interface{} `field:"optional" json:"captureResponse" yaml:"captureResponse"`
	// Code hook called after Amazon Lex successfully captures a slot value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotcapturesetting.html#cfn-lex-bot-slotcapturesetting-codehook
	//
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// Code hook called when Amazon Lex doesn't capture a slot value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotcapturesetting.html#cfn-lex-bot-slotcapturesetting-elicitationcodehook
	//
	ElicitationCodeHook interface{} `field:"optional" json:"elicitationCodeHook" yaml:"elicitationCodeHook"`
	// A list of conditional branches to evaluate when the slot value isn't captured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotcapturesetting.html#cfn-lex-bot-slotcapturesetting-failureconditional
	//
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// Specifies the next step that the bot runs when the slot value code is not recognized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotcapturesetting.html#cfn-lex-bot-slotcapturesetting-failurenextstep
	//
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input when the slot fails to be captured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotcapturesetting.html#cfn-lex-bot-slotcapturesetting-failureresponse
	//
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
}


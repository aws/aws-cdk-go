package previewawslexmixins


// Configuration setting for a response sent to the user before Amazon Lex starts eliciting slots.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-initialresponsesetting.html
//
type CfnBotPropsMixin_InitialResponseSettingProperty struct {
	// Settings that specify the dialog code hook that is called by Amazon Lex at a step of the conversation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-initialresponsesetting.html#cfn-lex-bot-initialresponsesetting-codehook
	//
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// Provides a list of conditional branches.
	//
	// Branches are evaluated in the order that they are entered in the list. The first branch with a condition that evaluates to true is executed. The last branch in the list is the default branch. The default branch should not have any condition expression. The default branch is executed if no other branch has a matching condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-initialresponsesetting.html#cfn-lex-bot-initialresponsesetting-conditional
	//
	Conditional interface{} `field:"optional" json:"conditional" yaml:"conditional"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-initialresponsesetting.html#cfn-lex-bot-initialresponsesetting-initialresponse
	//
	InitialResponse interface{} `field:"optional" json:"initialResponse" yaml:"initialResponse"`
	// The next step in the conversation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-initialresponsesetting.html#cfn-lex-bot-initialresponsesetting-nextstep
	//
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
}


package awslex


// Settings that specify the dialog code hook that is called by Amazon Lex at a step of the conversation.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehookinvocationsetting.html
//
type CfnBot_DialogCodeHookInvocationSettingProperty struct {
	// Indicates whether a Lambda function should be invoked for the dialog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehookinvocationsetting.html#cfn-lex-bot-dialogcodehookinvocationsetting-enablecodehookinvocation
	//
	EnableCodeHookInvocation interface{} `field:"required" json:"enableCodeHookInvocation" yaml:"enableCodeHookInvocation"`
	// Determines whether a dialog code hook is used when the intent is activated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehookinvocationsetting.html#cfn-lex-bot-dialogcodehookinvocationsetting-isactive
	//
	IsActive interface{} `field:"required" json:"isActive" yaml:"isActive"`
	// Contains the responses and actions that Amazon Lex takes after the Lambda function is complete.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehookinvocationsetting.html#cfn-lex-bot-dialogcodehookinvocationsetting-postcodehookspecification
	//
	PostCodeHookSpecification interface{} `field:"required" json:"postCodeHookSpecification" yaml:"postCodeHookSpecification"`
	// A label that indicates the dialog step from which the dialog code hook is happening.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehookinvocationsetting.html#cfn-lex-bot-dialogcodehookinvocationsetting-invocationlabel
	//
	InvocationLabel *string `field:"optional" json:"invocationLabel" yaml:"invocationLabel"`
}


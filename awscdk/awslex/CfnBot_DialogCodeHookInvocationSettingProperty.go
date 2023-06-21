package awslex


// Settings that specify the dialog code hook that is called by Amazon Lex at a step of the conversation.
//
// Example:
//
//
type CfnBot_DialogCodeHookInvocationSettingProperty struct {
	// Indicates whether a Lambda function should be invoked for the dialog.
	EnableCodeHookInvocation interface{} `field:"required" json:"enableCodeHookInvocation" yaml:"enableCodeHookInvocation"`
	// Determines whether a dialog code hook is used when the intent is activated.
	IsActive interface{} `field:"required" json:"isActive" yaml:"isActive"`
	// Contains the responses and actions that Amazon Lex takes after the Lambda function is complete.
	PostCodeHookSpecification interface{} `field:"required" json:"postCodeHookSpecification" yaml:"postCodeHookSpecification"`
	// A label that indicates the dialog step from which the dialog code hook is happening.
	InvocationLabel *string `field:"optional" json:"invocationLabel" yaml:"invocationLabel"`
}


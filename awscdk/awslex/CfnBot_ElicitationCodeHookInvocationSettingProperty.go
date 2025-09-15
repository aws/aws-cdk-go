package awslex


// Settings that specify the dialog code hook that is called by Amazon Lex between eliciting slot values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elicitationCodeHookInvocationSettingProperty := &ElicitationCodeHookInvocationSettingProperty{
//   	EnableCodeHookInvocation: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	InvocationLabel: jsii.String("invocationLabel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-elicitationcodehookinvocationsetting.html
//
type CfnBot_ElicitationCodeHookInvocationSettingProperty struct {
	// Indicates whether a Lambda function should be invoked for the dialog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-elicitationcodehookinvocationsetting.html#cfn-lex-bot-elicitationcodehookinvocationsetting-enablecodehookinvocation
	//
	EnableCodeHookInvocation interface{} `field:"required" json:"enableCodeHookInvocation" yaml:"enableCodeHookInvocation"`
	// A label that indicates the dialog step from which the dialog code hook is happening.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-elicitationcodehookinvocationsetting.html#cfn-lex-bot-elicitationcodehookinvocationsetting-invocationlabel
	//
	InvocationLabel *string `field:"optional" json:"invocationLabel" yaml:"invocationLabel"`
}


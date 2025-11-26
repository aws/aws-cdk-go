package previewawslexmixins


// Provides an expression that evaluates to true or false.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionProperty := &ConditionProperty{
//   	ExpressionString: jsii.String("expressionString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-condition.html
//
type CfnBotPropsMixin_ConditionProperty struct {
	// The expression string that is evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-condition.html#cfn-lex-bot-condition-expressionstring
	//
	ExpressionString *string `field:"optional" json:"expressionString" yaml:"expressionString"`
}


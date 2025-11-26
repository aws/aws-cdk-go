package previewawslexmixins


// Specifies the text input specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   textInputSpecificationProperty := &TextInputSpecificationProperty{
//   	StartTimeoutMs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-textinputspecification.html
//
type CfnBotPropsMixin_TextInputSpecificationProperty struct {
	// Time for which a bot waits before re-prompting a customer for text input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-textinputspecification.html#cfn-lex-bot-textinputspecification-starttimeoutms
	//
	StartTimeoutMs *float64 `field:"optional" json:"startTimeoutMs" yaml:"startTimeoutMs"`
}


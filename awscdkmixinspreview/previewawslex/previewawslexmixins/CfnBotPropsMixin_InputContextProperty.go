package previewawslexmixins


// A context that must be active for an intent to be selected by Amazon Lex.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputContextProperty := &InputContextProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-inputcontext.html
//
type CfnBotPropsMixin_InputContextProperty struct {
	// The name of the context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-inputcontext.html#cfn-lex-bot-inputcontext-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}


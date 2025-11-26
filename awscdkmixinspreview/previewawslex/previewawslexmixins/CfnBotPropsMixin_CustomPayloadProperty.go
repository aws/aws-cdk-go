package previewawslexmixins


// A custom response string that Amazon Lex sends to your application.
//
// You define the content and structure the string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customPayloadProperty := &CustomPayloadProperty{
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-custompayload.html
//
type CfnBotPropsMixin_CustomPayloadProperty struct {
	// The string that is sent to your application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-custompayload.html#cfn-lex-bot-custompayload-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


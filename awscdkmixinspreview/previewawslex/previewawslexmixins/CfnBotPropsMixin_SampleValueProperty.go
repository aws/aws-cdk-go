package previewawslexmixins


// Defines one of the values for a slot type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sampleValueProperty := &SampleValueProperty{
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-samplevalue.html
//
type CfnBotPropsMixin_SampleValueProperty struct {
	// The value that can be used for a slot type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-samplevalue.html#cfn-lex-bot-samplevalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


package awslex


// A sample utterance that invokes an intent or respond to a slot elicitation prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sampleUtteranceProperty := &SampleUtteranceProperty{
//   	Utterance: jsii.String("utterance"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-sampleutterance.html
//
type CfnBot_SampleUtteranceProperty struct {
	// A sample utterance that invokes an intent or respond to a slot elicitation prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-sampleutterance.html#cfn-lex-bot-sampleutterance-utterance
	//
	Utterance *string `field:"required" json:"utterance" yaml:"utterance"`
}


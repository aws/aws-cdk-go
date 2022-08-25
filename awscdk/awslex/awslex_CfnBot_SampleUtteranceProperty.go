package awslex


// A sample utterance that invokes and intent or responds to a slot elicitation prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sampleUtteranceProperty := &sampleUtteranceProperty{
//   	utterance: jsii.String("utterance"),
//   }
//
type CfnBot_SampleUtteranceProperty struct {
	// The sample utterance that Amazon Lex uses to build its machine-learning model to recognize intents.
	Utterance *string `field:"required" json:"utterance" yaml:"utterance"`
}


package awslex


// A sample utterance that invokes an intent or respond to a slot elicitation prompt.
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
	// A sample utterance that invokes an intent or respond to a slot elicitation prompt.
	Utterance *string `field:"required" json:"utterance" yaml:"utterance"`
}


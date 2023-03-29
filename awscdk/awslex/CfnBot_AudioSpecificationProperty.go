package awslex


// Specifies the audio input specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioSpecificationProperty := &AudioSpecificationProperty{
//   	EndTimeoutMs: jsii.Number(123),
//   	MaxLengthMs: jsii.Number(123),
//   }
//
type CfnBot_AudioSpecificationProperty struct {
	// Time for which a bot waits after the customer stops speaking to assume the utterance is finished.
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// Time for how long Amazon Lex waits before speech input is truncated and the speech is returned to application.
	MaxLengthMs *float64 `field:"required" json:"maxLengthMs" yaml:"maxLengthMs"`
}


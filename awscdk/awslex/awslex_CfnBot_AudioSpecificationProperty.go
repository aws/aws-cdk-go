package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioSpecificationProperty := &audioSpecificationProperty{
//   	endTimeoutMs: jsii.Number(123),
//   	maxLengthMs: jsii.Number(123),
//   }
//
type CfnBot_AudioSpecificationProperty struct {
	// `CfnBot.AudioSpecificationProperty.EndTimeoutMs`.
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// `CfnBot.AudioSpecificationProperty.MaxLengthMs`.
	MaxLengthMs *float64 `field:"required" json:"maxLengthMs" yaml:"maxLengthMs"`
}


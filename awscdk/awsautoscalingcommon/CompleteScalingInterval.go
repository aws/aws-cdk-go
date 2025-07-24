package awsautoscalingcommon


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   completeScalingInterval := &CompleteScalingInterval{
//   	Lower: jsii.Number(123),
//   	Upper: jsii.Number(123),
//
//   	// the properties below are optional
//   	Change: jsii.Number(123),
//   }
//
type CompleteScalingInterval struct {
	Lower *float64 `field:"required" json:"lower" yaml:"lower"`
	Upper *float64 `field:"required" json:"upper" yaml:"upper"`
	Change *float64 `field:"optional" json:"change" yaml:"change"`
}


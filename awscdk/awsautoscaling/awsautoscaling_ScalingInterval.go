package awsautoscaling


// A range of metric values in which to apply a certain scaling operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingInterval := &scalingInterval{
//   	change: jsii.Number(123),
//
//   	// the properties below are optional
//   	lower: jsii.Number(123),
//   	upper: jsii.Number(123),
//   }
//
type ScalingInterval struct {
	// The capacity adjustment to apply in this interval.
	//
	// The number is interpreted differently based on AdjustmentType:
	//
	// - ChangeInCapacity: add the adjustment to the current capacity.
	//   The number can be positive or negative.
	// - PercentChangeInCapacity: add or remove the given percentage of the current
	//    capacity to itself. The number can be in the range [-100..100].
	// - ExactCapacity: set the capacity to this number. The number must
	//    be positive.
	Change *float64 `field:"required" json:"change" yaml:"change"`
	// The lower bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is higher than this value.
	Lower *float64 `field:"optional" json:"lower" yaml:"lower"`
	// The upper bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is lower than this value.
	Upper *float64 `field:"optional" json:"upper" yaml:"upper"`
}


package awsautoscaling


// An adjustment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adjustmentTier := &AdjustmentTier{
//   	Adjustment: jsii.Number(123),
//
//   	// the properties below are optional
//   	LowerBound: jsii.Number(123),
//   	UpperBound: jsii.Number(123),
//   }
//
type AdjustmentTier struct {
	// What number to adjust the capacity with.
	//
	// The number is interpeted as an added capacity, a new fixed capacity or an
	// added percentage depending on the AdjustmentType value of the
	// StepScalingPolicy.
	//
	// Can be positive or negative.
	Adjustment *float64 `field:"required" json:"adjustment" yaml:"adjustment"`
	// Lower bound where this scaling tier applies.
	//
	// The scaling tier applies if the difference between the metric
	// value and its alarm threshold is higher than this value.
	// Default: -Infinity if this is the first tier, otherwise the upperBound of the previous tier.
	//
	LowerBound *float64 `field:"optional" json:"lowerBound" yaml:"lowerBound"`
	// Upper bound where this scaling tier applies.
	//
	// The scaling tier applies if the difference between the metric
	// value and its alarm threshold is lower than this value.
	// Default: +Infinity.
	//
	UpperBound *float64 `field:"optional" json:"upperBound" yaml:"upperBound"`
}


package awsautoscalingcommon

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarms := &alarms{
//   	lowerAlarmIntervalIndex: jsii.Number(123),
//   	upperAlarmIntervalIndex: jsii.Number(123),
//   }
//
// Experimental.
type Alarms struct {
	// Experimental.
	LowerAlarmIntervalIndex *float64 `field:"optional" json:"lowerAlarmIntervalIndex" yaml:"lowerAlarmIntervalIndex"`
	// Experimental.
	UpperAlarmIntervalIndex *float64 `field:"optional" json:"upperAlarmIntervalIndex" yaml:"upperAlarmIntervalIndex"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arbitraryIntervals := &arbitraryIntervals{
//   	absolute: jsii.Boolean(false),
//   	intervals: []scalingInterval{
//   		&scalingInterval{
//   			change: jsii.Number(123),
//
//   			// the properties below are optional
//   			lower: jsii.Number(123),
//   			upper: jsii.Number(123),
//   		},
//   	},
//   }
//
// Experimental.
type ArbitraryIntervals struct {
	// Experimental.
	Absolute *bool `field:"required" json:"absolute" yaml:"absolute"`
	// Experimental.
	Intervals *[]*ScalingInterval `field:"required" json:"intervals" yaml:"intervals"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   completeScalingInterval := &completeScalingInterval{
//   	lower: jsii.Number(123),
//   	upper: jsii.Number(123),
//
//   	// the properties below are optional
//   	change: jsii.Number(123),
//   }
//
// Experimental.
type CompleteScalingInterval struct {
	// Experimental.
	Lower *float64 `field:"required" json:"lower" yaml:"lower"`
	// Experimental.
	Upper *float64 `field:"required" json:"upper" yaml:"upper"`
	// Experimental.
	Change *float64 `field:"optional" json:"change" yaml:"change"`
}

// Experimental.
type IRandomGenerator interface {
	// Experimental.
	NextBoolean() *bool
	// Experimental.
	NextInt(min *float64, max *float64) *float64
}

// The jsii proxy for IRandomGenerator
type jsiiProxy_IRandomGenerator struct {
	_ byte // padding
}

func (i *jsiiProxy_IRandomGenerator) NextBoolean() *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"nextBoolean",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRandomGenerator) NextInt(min *float64, max *float64) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"nextInt",
		[]interface{}{min, max},
		&returns,
	)

	return returns
}

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
// Experimental.
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
	// Experimental.
	Change *float64 `field:"required" json:"change" yaml:"change"`
	// The lower bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is higher than this value.
	// Experimental.
	Lower *float64 `field:"optional" json:"lower" yaml:"lower"`
	// The upper bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is lower than this value.
	// Experimental.
	Upper *float64 `field:"optional" json:"upper" yaml:"upper"`
}


package awsautoscalingcommon

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// TODO: EXAMPLE
//
// Experimental.
type Alarms struct {
	// Experimental.
	LowerAlarmIntervalIndex *float64 `json:"lowerAlarmIntervalIndex" yaml:"lowerAlarmIntervalIndex"`
	// Experimental.
	UpperAlarmIntervalIndex *float64 `json:"upperAlarmIntervalIndex" yaml:"upperAlarmIntervalIndex"`
}

// TODO: EXAMPLE
//
// Experimental.
type ArbitraryIntervals struct {
	// Experimental.
	Absolute *bool `json:"absolute" yaml:"absolute"`
	// Experimental.
	Intervals *[]*ScalingInterval `json:"intervals" yaml:"intervals"`
}

// TODO: EXAMPLE
//
// Experimental.
type CompleteScalingInterval struct {
	// Experimental.
	Lower *float64 `json:"lower" yaml:"lower"`
	// Experimental.
	Upper *float64 `json:"upper" yaml:"upper"`
	// Experimental.
	Change *float64 `json:"change" yaml:"change"`
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
// TODO: EXAMPLE
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
	Change *float64 `json:"change" yaml:"change"`
	// The lower bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is higher than this value.
	// Experimental.
	Lower *float64 `json:"lower" yaml:"lower"`
	// The upper bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is lower than this value.
	// Experimental.
	Upper *float64 `json:"upper" yaml:"upper"`
}


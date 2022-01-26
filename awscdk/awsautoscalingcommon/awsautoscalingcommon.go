package awsautoscalingcommon

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// TODO: EXAMPLE
//
type Alarms struct {
	LowerAlarmIntervalIndex *float64 `json:"lowerAlarmIntervalIndex" yaml:"lowerAlarmIntervalIndex"`
	UpperAlarmIntervalIndex *float64 `json:"upperAlarmIntervalIndex" yaml:"upperAlarmIntervalIndex"`
}

// TODO: EXAMPLE
//
type ArbitraryIntervals struct {
	Absolute *bool `json:"absolute" yaml:"absolute"`
	Intervals *[]*ScalingInterval `json:"intervals" yaml:"intervals"`
}

// TODO: EXAMPLE
//
type CompleteScalingInterval struct {
	Lower *float64 `json:"lower" yaml:"lower"`
	Upper *float64 `json:"upper" yaml:"upper"`
	Change *float64 `json:"change" yaml:"change"`
}

type IRandomGenerator interface {
	NextBoolean() *bool
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
	Change *float64 `json:"change" yaml:"change"`
	// The lower bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is higher than this value.
	Lower *float64 `json:"lower" yaml:"lower"`
	// The upper bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is lower than this value.
	Upper *float64 `json:"upper" yaml:"upper"`
}


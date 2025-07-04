package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a traffic routing config of type 'TimeBasedLinear'.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeBasedLinearTrafficRouting := awscdk.Aws_codedeploy.TimeBasedLinearTrafficRouting_AllAtOnce()
//
type TimeBasedLinearTrafficRouting interface {
	TrafficRouting
	// The amount of time between additional traffic shifts.
	Interval() awscdk.Duration
	// The percentage to increase traffic on each traffic shift.
	Percentage() *float64
	// Return a TrafficRoutingConfig of type `TimeBasedLinear`.
	Bind(_scope constructs.Construct) *TrafficRoutingConfig
}

// The jsii proxy struct for TimeBasedLinearTrafficRouting
type jsiiProxy_TimeBasedLinearTrafficRouting struct {
	jsiiProxy_TrafficRouting
}

func (j *jsiiProxy_TimeBasedLinearTrafficRouting) Interval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeBasedLinearTrafficRouting) Percentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentage",
		&returns,
	)
	return returns
}


func NewTimeBasedLinearTrafficRouting(props *TimeBasedLinearTrafficRoutingProps) TimeBasedLinearTrafficRouting {
	_init_.Initialize()

	if err := validateNewTimeBasedLinearTrafficRoutingParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimeBasedLinearTrafficRouting{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.TimeBasedLinearTrafficRouting",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewTimeBasedLinearTrafficRouting_Override(t TimeBasedLinearTrafficRouting, props *TimeBasedLinearTrafficRoutingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.TimeBasedLinearTrafficRouting",
		[]interface{}{props},
		t,
	)
}

// Shifts 100% of traffic in a single shift.
func TimeBasedLinearTrafficRouting_AllAtOnce() TrafficRouting {
	_init_.Initialize()

	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.TimeBasedLinearTrafficRouting",
		"allAtOnce",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Shifts a specified percentage of traffic, waits for a specified amount of time, then shifts the rest of traffic.
func TimeBasedLinearTrafficRouting_TimeBasedCanary(props *TimeBasedCanaryTrafficRoutingProps) TrafficRouting {
	_init_.Initialize()

	if err := validateTimeBasedLinearTrafficRouting_TimeBasedCanaryParameters(props); err != nil {
		panic(err)
	}
	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.TimeBasedLinearTrafficRouting",
		"timeBasedCanary",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Keeps shifting a specified percentage of traffic until reaching 100%, waiting for a specified amount of time in between each traffic shift.
func TimeBasedLinearTrafficRouting_TimeBasedLinear(props *TimeBasedLinearTrafficRoutingProps) TrafficRouting {
	_init_.Initialize()

	if err := validateTimeBasedLinearTrafficRouting_TimeBasedLinearParameters(props); err != nil {
		panic(err)
	}
	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.TimeBasedLinearTrafficRouting",
		"timeBasedLinear",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeBasedLinearTrafficRouting) Bind(_scope constructs.Construct) *TrafficRoutingConfig {
	if err := t.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *TrafficRoutingConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}


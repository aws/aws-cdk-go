package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Define a traffic routing config of type 'AllAtOnce'.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allAtOnceTrafficRouting := awscdk.Aws_codedeploy.AllAtOnceTrafficRouting_AllAtOnce()
//
type AllAtOnceTrafficRouting interface {
	TrafficRouting
	// Return a TrafficRoutingConfig of type `AllAtOnce`.
	Bind(_scope constructs.Construct) *TrafficRoutingConfig
}

// The jsii proxy struct for AllAtOnceTrafficRouting
type jsiiProxy_AllAtOnceTrafficRouting struct {
	jsiiProxy_TrafficRouting
}

func NewAllAtOnceTrafficRouting() AllAtOnceTrafficRouting {
	_init_.Initialize()

	j := jsiiProxy_AllAtOnceTrafficRouting{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.AllAtOnceTrafficRouting",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAllAtOnceTrafficRouting_Override(a AllAtOnceTrafficRouting) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.AllAtOnceTrafficRouting",
		nil, // no parameters
		a,
	)
}

// Shifts 100% of traffic in a single shift.
func AllAtOnceTrafficRouting_AllAtOnce() TrafficRouting {
	_init_.Initialize()

	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.AllAtOnceTrafficRouting",
		"allAtOnce",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Shifts a specified percentage of traffic, waits for a specified amount of time, then shifts the rest of traffic.
func AllAtOnceTrafficRouting_TimeBasedCanary(props *TimeBasedCanaryTrafficRoutingProps) TrafficRouting {
	_init_.Initialize()

	if err := validateAllAtOnceTrafficRouting_TimeBasedCanaryParameters(props); err != nil {
		panic(err)
	}
	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.AllAtOnceTrafficRouting",
		"timeBasedCanary",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Keeps shifting a specified percentage of traffic until reaching 100%, waiting for a specified amount of time in between each traffic shift.
func AllAtOnceTrafficRouting_TimeBasedLinear(props *TimeBasedLinearTrafficRoutingProps) TrafficRouting {
	_init_.Initialize()

	if err := validateAllAtOnceTrafficRouting_TimeBasedLinearParameters(props); err != nil {
		panic(err)
	}
	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.AllAtOnceTrafficRouting",
		"timeBasedLinear",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AllAtOnceTrafficRouting) Bind(_scope constructs.Construct) *TrafficRoutingConfig {
	if err := a.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *TrafficRoutingConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}


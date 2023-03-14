package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents how traffic is shifted during a CodeDeploy deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficRouting := awscdk.Aws_codedeploy.TrafficRouting_AllAtOnce()
//
type TrafficRouting interface {
	// Returns the traffic routing configuration.
	Bind(scope constructs.Construct) *TrafficRoutingConfig
}

// The jsii proxy struct for TrafficRouting
type jsiiProxy_TrafficRouting struct {
	_ byte // padding
}

func NewTrafficRouting_Override(t TrafficRouting) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.TrafficRouting",
		nil, // no parameters
		t,
	)
}

// Shifts 100% of traffic in a single shift.
func TrafficRouting_AllAtOnce() TrafficRouting {
	_init_.Initialize()

	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.TrafficRouting",
		"allAtOnce",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Shifts a specified percentage of traffic, waits for a specified amount of time, then shifts the rest of traffic.
func TrafficRouting_TimeBasedCanary(props *TimeBasedCanaryTrafficRoutingProps) TrafficRouting {
	_init_.Initialize()

	if err := validateTrafficRouting_TimeBasedCanaryParameters(props); err != nil {
		panic(err)
	}
	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.TrafficRouting",
		"timeBasedCanary",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Keeps shifting a specified percentage of traffic until reaching 100%, waiting for a specified amount of time in between each traffic shift.
func TrafficRouting_TimeBasedLinear(props *TimeBasedLinearTrafficRoutingProps) TrafficRouting {
	_init_.Initialize()

	if err := validateTrafficRouting_TimeBasedLinearParameters(props); err != nil {
		panic(err)
	}
	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.TrafficRouting",
		"timeBasedLinear",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrafficRouting) Bind(scope constructs.Construct) *TrafficRoutingConfig {
	if err := t.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *TrafficRoutingConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


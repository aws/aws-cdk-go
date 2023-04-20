package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a traffic routing config of type 'TimeBasedCanary'.
//
// Example:
//   config := codedeploy.NewLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &LambdaDeploymentConfigProps{
//   	TrafficRouting: codedeploy.NewTimeBasedCanaryTrafficRouting(&TimeBasedCanaryTrafficRoutingProps{
//   		Interval: awscdk.Duration_Minutes(jsii.Number(15)),
//   		Percentage: jsii.Number(5),
//   	}),
//   	DeploymentConfigName: jsii.String("MyDeploymentConfig"),
//   })
//
type TimeBasedCanaryTrafficRouting interface {
	TrafficRouting
	// The amount of time between additional traffic shifts.
	Interval() awscdk.Duration
	// The percentage to increase traffic on each traffic shift.
	Percentage() *float64
	// Return a TrafficRoutingConfig of type `TimeBasedCanary`.
	Bind(_scope constructs.Construct) *TrafficRoutingConfig
}

// The jsii proxy struct for TimeBasedCanaryTrafficRouting
type jsiiProxy_TimeBasedCanaryTrafficRouting struct {
	jsiiProxy_TrafficRouting
}

func (j *jsiiProxy_TimeBasedCanaryTrafficRouting) Interval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeBasedCanaryTrafficRouting) Percentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentage",
		&returns,
	)
	return returns
}


func NewTimeBasedCanaryTrafficRouting(props *TimeBasedCanaryTrafficRoutingProps) TimeBasedCanaryTrafficRouting {
	_init_.Initialize()

	if err := validateNewTimeBasedCanaryTrafficRoutingParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimeBasedCanaryTrafficRouting{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.TimeBasedCanaryTrafficRouting",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewTimeBasedCanaryTrafficRouting_Override(t TimeBasedCanaryTrafficRouting, props *TimeBasedCanaryTrafficRoutingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.TimeBasedCanaryTrafficRouting",
		[]interface{}{props},
		t,
	)
}

// Shifts 100% of traffic in a single shift.
func TimeBasedCanaryTrafficRouting_AllAtOnce() TrafficRouting {
	_init_.Initialize()

	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.TimeBasedCanaryTrafficRouting",
		"allAtOnce",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Shifts a specified percentage of traffic, waits for a specified amount of time, then shifts the rest of traffic.
func TimeBasedCanaryTrafficRouting_TimeBasedCanary(props *TimeBasedCanaryTrafficRoutingProps) TrafficRouting {
	_init_.Initialize()

	if err := validateTimeBasedCanaryTrafficRouting_TimeBasedCanaryParameters(props); err != nil {
		panic(err)
	}
	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.TimeBasedCanaryTrafficRouting",
		"timeBasedCanary",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Keeps shifting a specified percentage of traffic until reaching 100%, waiting for a specified amount of time in between each traffic shift.
func TimeBasedCanaryTrafficRouting_TimeBasedLinear(props *TimeBasedLinearTrafficRoutingProps) TrafficRouting {
	_init_.Initialize()

	if err := validateTimeBasedCanaryTrafficRouting_TimeBasedLinearParameters(props); err != nil {
		panic(err)
	}
	var returns TrafficRouting

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.TimeBasedCanaryTrafficRouting",
		"timeBasedLinear",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimeBasedCanaryTrafficRouting) Bind(_scope constructs.Construct) *TrafficRoutingConfig {
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


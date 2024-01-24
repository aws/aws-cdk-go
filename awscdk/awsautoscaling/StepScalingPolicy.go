package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a acaling strategy which scales depending on absolute values of some metric.
//
// You can specify the scaling behavior for various values of the metric.
//
// Implemented using one or more CloudWatch alarms and Step Scaling Policies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var metric metric
//
//   stepScalingPolicy := awscdk.Aws_autoscaling.NewStepScalingPolicy(this, jsii.String("MyStepScalingPolicy"), &StepScalingPolicyProps{
//   	AutoScalingGroup: autoScalingGroup,
//   	Metric: metric,
//   	ScalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			Change: jsii.Number(123),
//
//   			// the properties below are optional
//   			Lower: jsii.Number(123),
//   			Upper: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AdjustmentType: awscdk.*Aws_autoscaling.AdjustmentType_CHANGE_IN_CAPACITY,
//   	Cooldown: cdk.Duration_Minutes(jsii.Number(30)),
//   	DatapointsToAlarm: jsii.Number(123),
//   	EstimatedInstanceWarmup: cdk.Duration_*Minutes(jsii.Number(30)),
//   	EvaluationPeriods: jsii.Number(123),
//   	MetricAggregationType: awscdk.*Aws_autoscaling.MetricAggregationType_AVERAGE,
//   	MinAdjustmentMagnitude: jsii.Number(123),
//   })
//
type StepScalingPolicy interface {
	constructs.Construct
	LowerAction() StepScalingAction
	LowerAlarm() awscloudwatch.Alarm
	// The tree node.
	Node() constructs.Node
	UpperAction() StepScalingAction
	UpperAlarm() awscloudwatch.Alarm
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for StepScalingPolicy
type jsiiProxy_StepScalingPolicy struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_StepScalingPolicy) LowerAction() StepScalingAction {
	var returns StepScalingAction
	_jsii_.Get(
		j,
		"lowerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) LowerAlarm() awscloudwatch.Alarm {
	var returns awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"lowerAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) UpperAction() StepScalingAction {
	var returns StepScalingAction
	_jsii_.Get(
		j,
		"upperAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) UpperAlarm() awscloudwatch.Alarm {
	var returns awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"upperAlarm",
		&returns,
	)
	return returns
}


func NewStepScalingPolicy(scope constructs.Construct, id *string, props *StepScalingPolicyProps) StepScalingPolicy {
	_init_.Initialize()

	if err := validateNewStepScalingPolicyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepScalingPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.StepScalingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStepScalingPolicy_Override(s StepScalingPolicy, scope constructs.Construct, id *string, props *StepScalingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.StepScalingPolicy",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func StepScalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStepScalingPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.StepScalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepScalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


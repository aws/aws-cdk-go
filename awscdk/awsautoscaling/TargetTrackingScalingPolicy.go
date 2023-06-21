package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

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
//   targetTrackingScalingPolicy := awscdk.Aws_autoscaling.NewTargetTrackingScalingPolicy(this, jsii.String("MyTargetTrackingScalingPolicy"), &TargetTrackingScalingPolicyProps{
//   	AutoScalingGroup: autoScalingGroup,
//   	TargetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	Cooldown: cdk.Duration_Minutes(jsii.Number(30)),
//   	CustomMetric: metric,
//   	DisableScaleIn: jsii.Boolean(false),
//   	EstimatedInstanceWarmup: cdk.Duration_*Minutes(jsii.Number(30)),
//   	PredefinedMetric: awscdk.*Aws_autoscaling.PredefinedMetric_ASG_AVERAGE_CPU_UTILIZATION,
//   	ResourceLabel: jsii.String("resourceLabel"),
//   })
//
type TargetTrackingScalingPolicy interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// ARN of the scaling policy.
	ScalingPolicyArn() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TargetTrackingScalingPolicy
type jsiiProxy_TargetTrackingScalingPolicy struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_TargetTrackingScalingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetTrackingScalingPolicy) ScalingPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyArn",
		&returns,
	)
	return returns
}


func NewTargetTrackingScalingPolicy(scope constructs.Construct, id *string, props *TargetTrackingScalingPolicyProps) TargetTrackingScalingPolicy {
	_init_.Initialize()

	if err := validateNewTargetTrackingScalingPolicyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TargetTrackingScalingPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.TargetTrackingScalingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTargetTrackingScalingPolicy_Override(t TargetTrackingScalingPolicy, scope constructs.Construct, id *string, props *TargetTrackingScalingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.TargetTrackingScalingPolicy",
		[]interface{}{scope, id, props},
		t,
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
func TargetTrackingScalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTargetTrackingScalingPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.TargetTrackingScalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TargetTrackingScalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


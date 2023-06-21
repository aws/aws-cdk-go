package awsapplicationautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represent an attribute for which autoscaling can be configured.
//
// This class is basically a light wrapper around ScalableTarget, but with
// all methods protected instead of public so they can be selectively
// exposed and/or more specific versions of them can be exposed by derived
// classes for individual services support autoscaling.
//
// Typical use cases:
//
// - Hide away the PredefinedMetric enum for target tracking policies.
// - Don't expose all scaling methods (for example Dynamo tables don't support
//   Step Scaling, so the Dynamo subclass won't expose this method).
type BaseScalableAttribute interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	Props() *BaseScalableAttributeProps
	// Scale out or in based on a metric value.
	DoScaleOnMetric(id *string, props *BasicStepScalingPolicyProps)
	// Scale out or in based on time.
	DoScaleOnSchedule(id *string, props *ScalingSchedule)
	// Scale out or in in order to keep a metric around a target value.
	DoScaleToTrackMetric(id *string, props *BasicTargetTrackingScalingPolicyProps)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BaseScalableAttribute
type jsiiProxy_BaseScalableAttribute struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BaseScalableAttribute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseScalableAttribute) Props() *BaseScalableAttributeProps {
	var returns *BaseScalableAttributeProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewBaseScalableAttribute_Override(b BaseScalableAttribute, scope constructs.Construct, id *string, props *BaseScalableAttributeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationautoscaling.BaseScalableAttribute",
		[]interface{}{scope, id, props},
		b,
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
func BaseScalableAttribute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBaseScalableAttribute_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.BaseScalableAttribute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseScalableAttribute) DoScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) {
	if err := b.validateDoScaleOnMetricParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"doScaleOnMetric",
		[]interface{}{id, props},
	)
}

func (b *jsiiProxy_BaseScalableAttribute) DoScaleOnSchedule(id *string, props *ScalingSchedule) {
	if err := b.validateDoScaleOnScheduleParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"doScaleOnSchedule",
		[]interface{}{id, props},
	)
}

func (b *jsiiProxy_BaseScalableAttribute) DoScaleToTrackMetric(id *string, props *BasicTargetTrackingScalingPolicyProps) {
	if err := b.validateDoScaleToTrackMetricParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"doScaleToTrackMetric",
		[]interface{}{id, props},
	)
}

func (b *jsiiProxy_BaseScalableAttribute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


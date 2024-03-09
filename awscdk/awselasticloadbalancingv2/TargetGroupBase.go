package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define the target of a load balancer.
type TargetGroupBase interface {
	constructs.Construct
	ITargetGroup
	// Default port configured for members of this target group.
	DefaultPort() *float64
	// Full name of first load balancer.
	//
	// This identifier is emitted as a dimensions of the metrics of this target
	// group.
	//
	// Example value: `app/my-load-balancer/123456789`.
	FirstLoadBalancerFullName() *string
	// Health check for the members of this target group.
	HealthCheck() *HealthCheck
	SetHealthCheck(val *HealthCheck)
	// A token representing a list of ARNs of the load balancers that route traffic to this target group.
	LoadBalancerArns() *string
	// List of constructs that need to be depended on to ensure the TargetGroup is associated to a load balancer.
	LoadBalancerAttached() constructs.IDependable
	// Configurable dependable with all resources that lead to load balancer attachment.
	LoadBalancerAttachedDependencies() constructs.DependencyGroup
	// The tree node.
	Node() constructs.Node
	// The ARN of the target group.
	TargetGroupArn() *string
	// The full name of the target group.
	TargetGroupFullName() *string
	// ARNs of load balancers load balancing to this TargetGroup.
	TargetGroupLoadBalancerArns() *[]*string
	// The name of the target group.
	TargetGroupName() *string
	// The types of the directly registered members of this target group.
	TargetType() TargetType
	SetTargetType(val TargetType)
	// Register the given load balancing target as part of this group.
	AddLoadBalancerTarget(props *LoadBalancerTargetProps)
	// Set/replace the target group's health check.
	ConfigureHealthCheck(healthCheck *HealthCheck)
	// Set a non-standard attribute on the target group.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
	ValidateHealthCheck() *[]*string
	ValidateTargetGroup() *[]*string
}

// The jsii proxy struct for TargetGroupBase
type jsiiProxy_TargetGroupBase struct {
	internal.Type__constructsConstruct
	jsiiProxy_ITargetGroup
}

func (j *jsiiProxy_TargetGroupBase) DefaultPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) FirstLoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstLoadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) HealthCheck() *HealthCheck {
	var returns *HealthCheck
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) LoadBalancerArns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) LoadBalancerAttached() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"loadBalancerAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) LoadBalancerAttachedDependencies() constructs.DependencyGroup {
	var returns constructs.DependencyGroup
	_jsii_.Get(
		j,
		"loadBalancerAttachedDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) TargetGroupFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) TargetGroupLoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupLoadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) TargetType() TargetType {
	var returns TargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


func NewTargetGroupBase_Override(t TargetGroupBase, scope constructs.Construct, id *string, baseProps *BaseTargetGroupProps, additionalProps interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.TargetGroupBase",
		[]interface{}{scope, id, baseProps, additionalProps},
		t,
	)
}

func (j *jsiiProxy_TargetGroupBase)SetHealthCheck(val *HealthCheck) {
	if err := j.validateSetHealthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_TargetGroupBase)SetTargetType(val TargetType) {
	_jsii_.Set(
		j,
		"targetType",
		val,
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
func TargetGroupBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTargetGroupBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.TargetGroupBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TargetGroupBase) AddLoadBalancerTarget(props *LoadBalancerTargetProps) {
	if err := t.validateAddLoadBalancerTargetParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addLoadBalancerTarget",
		[]interface{}{props},
	)
}

func (t *jsiiProxy_TargetGroupBase) ConfigureHealthCheck(healthCheck *HealthCheck) {
	if err := t.validateConfigureHealthCheckParameters(healthCheck); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"configureHealthCheck",
		[]interface{}{healthCheck},
	)
}

func (t *jsiiProxy_TargetGroupBase) SetAttribute(key *string, value *string) {
	if err := t.validateSetAttributeParameters(key); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"setAttribute",
		[]interface{}{key, value},
	)
}

func (t *jsiiProxy_TargetGroupBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TargetGroupBase) ValidateHealthCheck() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validateHealthCheck",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TargetGroupBase) ValidateTargetGroup() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validateTargetGroup",
		nil, // no parameters
		&returns,
	)

	return returns
}


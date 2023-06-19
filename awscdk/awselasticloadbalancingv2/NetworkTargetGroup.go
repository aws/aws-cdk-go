package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v3"
)

// Define a Network Target Group.
//
// Example:
//   var listener networkListener
//   var asg1 autoScalingGroup
//   var asg2 autoScalingGroup
//
//
//   group := listener.AddTargets(jsii.String("AppFleet"), &AddNetworkTargetsProps{
//   	Port: jsii.Number(443),
//   	Targets: []iNetworkLoadBalancerTarget{
//   		asg1,
//   	},
//   })
//
//   group.AddTarget(asg2)
//
// Experimental.
type NetworkTargetGroup interface {
	TargetGroupBase
	INetworkTargetGroup
	// Default port configured for members of this target group.
	// Experimental.
	DefaultPort() *float64
	// Full name of first load balancer.
	// Experimental.
	FirstLoadBalancerFullName() *string
	// Experimental.
	HealthCheck() *HealthCheck
	// Experimental.
	SetHealthCheck(val *HealthCheck)
	// A token representing a list of ARNs of the load balancers that route traffic to this target group.
	// Experimental.
	LoadBalancerArns() *string
	// List of constructs that need to be depended on to ensure the TargetGroup is associated to a load balancer.
	// Experimental.
	LoadBalancerAttached() awscdk.IDependable
	// Configurable dependable with all resources that lead to load balancer attachment.
	// Experimental.
	LoadBalancerAttachedDependencies() awscdk.ConcreteDependable
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The ARN of the target group.
	// Experimental.
	TargetGroupArn() *string
	// The full name of the target group.
	// Experimental.
	TargetGroupFullName() *string
	// ARNs of load balancers load balancing to this TargetGroup.
	// Experimental.
	TargetGroupLoadBalancerArns() *[]*string
	// The name of the target group.
	// Experimental.
	TargetGroupName() *string
	// The types of the directly registered members of this target group.
	// Experimental.
	TargetType() TargetType
	// Experimental.
	SetTargetType(val TargetType)
	// Register the given load balancing target as part of this group.
	// Experimental.
	AddLoadBalancerTarget(props *LoadBalancerTargetProps)
	// Add a load balancing target to this target group.
	// Experimental.
	AddTarget(targets ...INetworkLoadBalancerTarget)
	// Set/replace the target group's health check.
	// Experimental.
	ConfigureHealthCheck(healthCheck *HealthCheck)
	// The number of targets that are considered healthy.
	// Experimental.
	MetricHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of targets that are considered unhealthy.
	// Experimental.
	MetricUnHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Register a listener that is load balancing to this target group.
	//
	// Don't call this directly. It will be called by listeners.
	// Experimental.
	RegisterListener(listener INetworkListener)
	// Set a non-standard attribute on the target group.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
	//
	// Experimental.
	SetAttribute(key *string, value *string)
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for NetworkTargetGroup
type jsiiProxy_NetworkTargetGroup struct {
	jsiiProxy_TargetGroupBase
	jsiiProxy_INetworkTargetGroup
}

func (j *jsiiProxy_NetworkTargetGroup) DefaultPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) FirstLoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstLoadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) HealthCheck() *HealthCheck {
	var returns *HealthCheck
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) LoadBalancerArns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) LoadBalancerAttached() awscdk.IDependable {
	var returns awscdk.IDependable
	_jsii_.Get(
		j,
		"loadBalancerAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) LoadBalancerAttachedDependencies() awscdk.ConcreteDependable {
	var returns awscdk.ConcreteDependable
	_jsii_.Get(
		j,
		"loadBalancerAttachedDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) TargetGroupFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) TargetGroupLoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupLoadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) TargetType() TargetType {
	var returns TargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


// Experimental.
func NewNetworkTargetGroup(scope constructs.Construct, id *string, props *NetworkTargetGroupProps) NetworkTargetGroup {
	_init_.Initialize()

	if err := validateNewNetworkTargetGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkTargetGroup{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2.NetworkTargetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewNetworkTargetGroup_Override(n NetworkTargetGroup, scope constructs.Construct, id *string, props *NetworkTargetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2.NetworkTargetGroup",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NetworkTargetGroup)SetHealthCheck(val *HealthCheck) {
	if err := j.validateSetHealthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_NetworkTargetGroup)SetTargetType(val TargetType) {
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

// Import an existing target group.
// Experimental.
func NetworkTargetGroup_FromTargetGroupAttributes(scope constructs.Construct, id *string, attrs *TargetGroupAttributes) INetworkTargetGroup {
	_init_.Initialize()

	if err := validateNetworkTargetGroup_FromTargetGroupAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns INetworkTargetGroup

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.NetworkTargetGroup",
		"fromTargetGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing listener.
// Deprecated: Use `fromTargetGroupAttributes` instead.
func NetworkTargetGroup_Import(scope constructs.Construct, id *string, props *TargetGroupImportProps) INetworkTargetGroup {
	_init_.Initialize()

	if err := validateNetworkTargetGroup_ImportParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns INetworkTargetGroup

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.NetworkTargetGroup",
		"import",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func NetworkTargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkTargetGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.NetworkTargetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkTargetGroup) AddLoadBalancerTarget(props *LoadBalancerTargetProps) {
	if err := n.validateAddLoadBalancerTargetParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addLoadBalancerTarget",
		[]interface{}{props},
	)
}

func (n *jsiiProxy_NetworkTargetGroup) AddTarget(targets ...INetworkLoadBalancerTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addTarget",
		args,
	)
}

func (n *jsiiProxy_NetworkTargetGroup) ConfigureHealthCheck(healthCheck *HealthCheck) {
	if err := n.validateConfigureHealthCheckParameters(healthCheck); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"configureHealthCheck",
		[]interface{}{healthCheck},
	)
}

func (n *jsiiProxy_NetworkTargetGroup) MetricHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := n.validateMetricHealthyHostCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricHealthyHostCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkTargetGroup) MetricUnHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := n.validateMetricUnHealthyHostCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricUnHealthyHostCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkTargetGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		n,
		"onPrepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkTargetGroup) OnSynthesize(session constructs.ISynthesisSession) {
	if err := n.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (n *jsiiProxy_NetworkTargetGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkTargetGroup) Prepare() {
	_jsii_.InvokeVoid(
		n,
		"prepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkTargetGroup) RegisterListener(listener INetworkListener) {
	if err := n.validateRegisterListenerParameters(listener); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"registerListener",
		[]interface{}{listener},
	)
}

func (n *jsiiProxy_NetworkTargetGroup) SetAttribute(key *string, value *string) {
	if err := n.validateSetAttributeParameters(key); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"setAttribute",
		[]interface{}{key, value},
	)
}

func (n *jsiiProxy_NetworkTargetGroup) Synthesize(session awscdk.ISynthesisSession) {
	if err := n.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{session},
	)
}

func (n *jsiiProxy_NetworkTargetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkTargetGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}


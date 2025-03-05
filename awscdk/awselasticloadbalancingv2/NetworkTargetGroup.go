package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Network Target Group.
//
// Example:
//   import elb "github.com/aws/aws-cdk-go/awscdk"
//   import elb2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var clb loadBalancer
//   var alb applicationLoadBalancer
//   var nlb networkLoadBalancer
//
//
//   albListener := alb.AddListener(jsii.String("ALBListener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//   })
//   albTargetGroup := albListener.AddTargets(jsii.String("ALBFleet"), &AddApplicationTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   nlbListener := nlb.AddListener(jsii.String("NLBListener"), &BaseNetworkListenerProps{
//   	Port: jsii.Number(80),
//   })
//   nlbTargetGroup := nlbListener.AddTargets(jsii.String("NLBFleet"), &AddNetworkTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &ServerDeploymentGroupProps{
//   	LoadBalancers: []loadBalancer{
//   		codedeploy.*loadBalancer_Classic(clb),
//   		codedeploy.*loadBalancer_Application(albTargetGroup),
//   		codedeploy.*loadBalancer_Network(nlbTargetGroup),
//   	},
//   })
//
type NetworkTargetGroup interface {
	TargetGroupBase
	INetworkTargetGroup
	// Default port configured for members of this target group.
	DefaultPort() *float64
	// Full name of first load balancer.
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
	// All metrics available for this target group.
	Metrics() INetworkTargetGroupMetrics
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
	// Add a load balancing target to this target group.
	AddTarget(targets ...INetworkLoadBalancerTarget)
	// Set/replace the target group's health check.
	ConfigureHealthCheck(healthCheck *HealthCheck)
	// The number of targets that are considered healthy.
	// Default: Average over 5 minutes.
	//
	// Deprecated: Use ``NetworkTargetGroup.metrics.healthyHostCount`` instead
	MetricHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of targets that are considered unhealthy.
	// Default: Average over 5 minutes.
	//
	// Deprecated: Use ``NetworkTargetGroup.metrics.healthyHostCount`` instead
	MetricUnHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Register a listener that is load balancing to this target group.
	//
	// Don't call this directly. It will be called by listeners.
	RegisterListener(listener INetworkListener)
	// Set a non-standard attribute on the target group.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
	ValidateHealthCheck() *[]*string
	ValidateTargetGroup() *[]*string
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

func (j *jsiiProxy_NetworkTargetGroup) LoadBalancerAttached() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"loadBalancerAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) LoadBalancerAttachedDependencies() constructs.DependencyGroup {
	var returns constructs.DependencyGroup
	_jsii_.Get(
		j,
		"loadBalancerAttachedDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) Metrics() INetworkTargetGroupMetrics {
	var returns INetworkTargetGroupMetrics
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) Node() constructs.Node {
	var returns constructs.Node
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


func NewNetworkTargetGroup(scope constructs.Construct, id *string, props *NetworkTargetGroupProps) NetworkTargetGroup {
	_init_.Initialize()

	if err := validateNewNetworkTargetGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkTargetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkTargetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNetworkTargetGroup_Override(n NetworkTargetGroup, scope constructs.Construct, id *string, props *NetworkTargetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkTargetGroup",
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
func NetworkTargetGroup_FromTargetGroupAttributes(scope constructs.Construct, id *string, attrs *TargetGroupAttributes) INetworkTargetGroup {
	_init_.Initialize()

	if err := validateNetworkTargetGroup_FromTargetGroupAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns INetworkTargetGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkTargetGroup",
		"fromTargetGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
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
func NetworkTargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkTargetGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkTargetGroup",
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

func (n *jsiiProxy_NetworkTargetGroup) ValidateHealthCheck() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validateHealthCheck",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkTargetGroup) ValidateTargetGroup() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validateTargetGroup",
		nil, // no parameters
		&returns,
	)

	return returns
}

